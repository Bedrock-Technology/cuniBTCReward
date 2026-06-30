// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"cuniBTCReward/api/internal/svc"
	"cuniBTCReward/api/internal/types"
	"cuniBTCReward/model"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spruceid/siwe-go"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SignTermsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignTermsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignTermsLogic {
	return &SignTermsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignTermsLogic) SignTerms(req *types.SignTermsReq) (resp *types.SignTermsResp, err error) {
	message, err := siwe.ParseMessage(req.Message)
	if err != nil {
		return
	}
	if *message.GetStatement() == "" || message.GetAddress().String() == "" {
		return resp, fmt.Errorf("not contain, statement or address")
	}
	statementMd5 := md5.Sum([]byte(*message.GetStatement()))
	statementMd5Str := hex.EncodeToString(statementMd5[:])

	var symbolStatementMd5 string
	for _, v := range l.svcCtx.Config.Terms {
		if v.Symbol == req.Symbol {
			symbolStatementMd5 = v.TermMd5
		}
	}

	if statementMd5Str != symbolStatementMd5 {
		l.Infof("address: %s, symbolStatementMd5: %s, statementMd5Str: %s, message: %s, statement: %s",
			message.GetAddress().String(), symbolStatementMd5, statementMd5Str, req.Message, *message.GetStatement())
		l.Errorf("address: %s, signTerms error", message.GetAddress().String())
		return nil, fmt.Errorf("statement not equal")
	}

	contract, err := IsContract(l.svcCtx.Config.EvmHost, message.GetAddress().String())
	if err != nil {
		l.Errorf("isContract error:", err)
	}

	if contract {
		messageHash := accounts.TextHash([]byte(req.Message))
		safeHash := fmt.Sprintf("0x%x", GetSafeMessageHash(common.HexToAddress(message.GetAddress().String()), big.NewInt(1), messageHash))
		safeResp, err1 := httpc.Do(context.Background(),
			http.MethodGet, fmt.Sprintf("https://api.safe.global/tx-service/eth/api/v1/messages/%s", safeHash), nil)
		if err1 != nil {
			logx.Errorf("get safe error")
			return
		}
		defer safeResp.Body.Close()
		if safeResp.StatusCode != http.StatusOK {
			logx.Errorf("not found in safe, status: %d", safeResp.StatusCode)
			return nil, fmt.Errorf("not found in safe")
		}
		//safe wallet is 1/1
		valid, _ := VerifySafeSignature(l.svcCtx.Config.EvmHost, message.GetAddress().String(), fmt.Sprintf("0x%x", messageHash), req.Signature)
		//save into db
		term := model.SignTerms{
			Address:     message.GetAddress().String(),
			Symbol:      req.Symbol,
			TermHash:    statementMd5Str,
			Message:     req.Message,
			Signature:   req.Signature,
			MessageHash: safeHash,
		}
		if valid {
			term.Valid = true
		}
		if err := l.svcCtx.Database.WithContext(l.ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "address"}, {Name: "symbol"}, {Name: "term_hash"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"message_hash": term.MessageHash,
				"valid":        term.Valid,
				"updated_at":   gorm.Expr("NOW()"),
			}),
		}).Create(&term).Error; err != nil {
			return nil, err
		}
	} else {
		_, err = message.VerifyEIP191(req.Signature)
		if err != nil {
			return
		}

		//save into db
		term := model.SignTerms{
			Address:   message.GetAddress().String(),
			Symbol:    req.Symbol,
			TermHash:  statementMd5Str,
			Message:   req.Message,
			Signature: req.Signature,
			Valid:     true,
		}
		if err := l.svcCtx.Database.WithContext(l.ctx).Save(&term).Error; err != nil {
			return nil, err
		}
	}
	resp = &types.SignTermsResp{
		Address:  message.GetAddress().String(),
		TermsMd5: statementMd5Str,
		Symbol:   req.Symbol,
	}
	return
}

var EIP1271MagicValue = [4]byte{0x16, 0x26, 0xba, 0x7e}

const EIP1271ABI = `[
	{
		"constant": true,
		"inputs": [
			{"name": "_hash", "type": "bytes32"},
			{"name": "_signature", "type": "bytes"}
		],
		"name": "isValidSignature",
		"outputs": [
			{"name": "magicValue", "type": "bytes4"}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`

func VerifySafeSignature(rpcURL, safeAddrHex, messageHashHex, signatureHex string) (bool, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return false, fmt.Errorf("failed to connect to RPC: %v", err)
	}
	defer client.Close()

	parsedABI, err := abi.JSON(strings.NewReader(EIP1271ABI))
	if err != nil {
		return false, fmt.Errorf("failed to parse ABI: %v", err)
	}

	safeAddress := common.HexToAddress(safeAddrHex)

	msgHashBytes, err := hex.DecodeString(strings.TrimPrefix(messageHashHex, "0x"))
	if err != nil || len(msgHashBytes) != 32 {
		return false, fmt.Errorf("invalid message hash format")
	}
	var messageHash [32]byte
	copy(messageHash[:], msgHashBytes)

	signature, err := hex.DecodeString(strings.TrimPrefix(signatureHex, "0x"))
	if err != nil {
		return false, fmt.Errorf("invalid signature format")
	}

	inputData, err := parsedABI.Pack("isValidSignature", messageHash, signature)
	if err != nil {
		return false, fmt.Errorf("failed to pack arguments: %v", err)
	}

	msg := ethereum.CallMsg{
		To:   &safeAddress,
		Data: inputData,
	}

	outputData, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return false, fmt.Errorf("contract call failed: %v", err)
	}

	var magicValue [4]byte
	err = parsedABI.UnpackIntoInterface(&magicValue, "isValidSignature", outputData)
	if err != nil {
		return false, fmt.Errorf("failed to unpack output: %v", err)
	}

	if magicValue == EIP1271MagicValue {
		return true, nil
	}

	return false, nil
}

func IsContract(rpcHost string, addressHex string) (bool, error) {
	address := common.HexToAddress(addressHex)
	client, err := ethclient.Dial(rpcHost)
	if err != nil {
		return false, err
	}
	defer client.Close()
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false, err
	}

	return len(bytecode) > 0, nil
}
func GetSafeMessageHash(safeAddress common.Address, chainID *big.Int, messageHash []byte) []byte {
	domainTypeHash := crypto.Keccak256([]byte("EIP712Domain(uint256 chainId,address verifyingContract)"))

	domainData := make([]byte, 0, 96) // 32*3 bytes
	domainData = append(domainData, domainTypeHash...)
	domainData = append(domainData, math.U256Bytes(chainID)...)
	domainData = append(domainData, common.LeftPadBytes(safeAddress.Bytes(), 32)...)

	domainSeparator := crypto.Keccak256(domainData)

	safeMsgTypeHash := crypto.Keccak256([]byte("SafeMessage(bytes message)"))

	msgValueHash := crypto.Keccak256(messageHash)

	structData := make([]byte, 0, 64) // 32*2 bytes
	structData = append(structData, safeMsgTypeHash...)
	structData = append(structData, msgValueHash...)

	structHash := crypto.Keccak256(structData)

	finalData := make([]byte, 0, 66) // 2 + 32 + 32 bytes
	finalData = append(finalData, []byte{0x19, 0x01}...)
	finalData = append(finalData, domainSeparator...)
	finalData = append(finalData, structHash...)

	return crypto.Keccak256(finalData)
}
