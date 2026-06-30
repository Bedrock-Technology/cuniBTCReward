package logic

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spruceid/siwe-go"
)

// https://1.x.wagmi.sh/examples/sign-in-with-ethereum
func TestVerify1(t *testing.T) {
	option := map[string]any{"statement": "Sign in with Ethereum to the app.", "issuedAt": "2026-06-04T15:08:13.730Z"}
	message, err := siwe.InitMessage("1.x.wagmi.sh", "0xac07f2721EcD955c4370e7388922fA547E922A4f", "https://1.x.wagmi.sh", "OzmhV6oMfX4gIMLue", option)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Printf("\n%q\n", message.String())
	fmt.Printf("\n%s\n", message.String())
	signature := "0xb478e926a02385d3307c6ca105f486097fc03b50f368f7a34c651a176018b10834081ee24e35aac0d933fced95d21967bc1de6e6da39c462541d13599bbf99cf1c"
	pkey, err := message.VerifyEIP191(signature)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(crypto.PubkeyToAddress(*pkey))
}

// https://1.x.wagmi.sh/examples/sign-in-with-ethereum
func TestVerify2(t *testing.T) {
	messageStr := "1.x.wagmi.sh wants you to sign in with your Ethereum account:\n0xac07f2721EcD955c4370e7388922fA547E922A4f\n\nSign in with Ethereum to the app.\n\nURI: https://1.x.wagmi.sh\nVersion: 1\nChain ID: 1\nNonce: OzmhV6oMfX4gIMLue\nIssued At: 2026-06-04T15:08:13.730Z"
	message, err := siwe.ParseMessage(messageStr)
	if err != nil {
		t.Log(err)
		return
	}
	signature := "0xb478e926a02385d3307c6ca105f486097fc03b50f368f7a34c651a176018b10834081ee24e35aac0d933fced95d21967bc1de6e6da39c462541d13599bbf99cf1c"
	pkey, err := message.VerifyEIP191(signature)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(crypto.PubkeyToAddress(*pkey))
}
func TestVerify3(t *testing.T) {
	messageStr := "1.x.wagmi.sh wants you to sign in with your Ethereum account:\n0xe8239B17034c372CDF8A5F8d3cCb7Cf1795c4572\n\nSign in with Ethereum to the app.\n\nURI: https://1.x.wagmi.sh\nVersion: 1\nChain ID: 1\nNonce: OzmhV6oMfX4gIMLue\nIssued At: 2026-06-04T15:08:13.730Z"
	message, err := siwe.ParseMessage(messageStr)
	if err != nil {
		t.Log(err)
		return
	}
	signature := "0xb478e926a02385d3307c6ca105f486097fc03b50f368f7a34c651a176018b10834081ee24e35aac0d933fced95d21967bc1de6e6da39c462541d13599bbf99cf1c"
	pkey, err := message.VerifyEIP191(signature)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(crypto.PubkeyToAddress(*pkey))
}

var RPC = ""

func TestVerify4(t *testing.T) {
	messageStr := "vaults.localhost wants you to sign in with your Ethereum account:\n0x916E4FE9C7b277828815f94f79d9E15dBFD37083\n\nI agree to Bedrock Vault ToS v1.0, https://legal.bedrock.technology/tos-selini-v1\n\nURI: https://vaults.localhost\nVersion: 1\nChain ID: 1\nNonce: 0xa8198f3c\nIssued At: 2026-06-29T10:12:21.185Z"
	signature := "0x4f34dfef91a97acc7fbfbad0b32e289925ebf0ec9f506b3bbb3e53527154c2162cccda5ad350d8aa9778c2dd9ccb3423408c05aeaf1296442f9d07d8351beee81c"
	originalHash := accounts.TextHash([]byte(messageStr))
	t.Logf("safeHash:0x%x", originalHash)
	b, err := VerifySafeSignature(RPC, "0x916E4FE9C7b277828815f94f79d9E15dBFD37083", fmt.Sprintf("0x%x", originalHash), signature)
	t.Log(b)
	t.Log(err)
}
func TestVerify4Fail(t *testing.T) {
	messageStr := "vaults.localhost wants you to sign in with your Ethereum account:\n0x916E4FE9C7b277828815f94f79d9E15dBFD37083\n\nI agree to Bedrock Vault ToS v1.0, https://legal.bedrock.technology/tos-selini-v1\n\nURI: https://vaults.localhost\nVersion: 1\nChain ID: 1\nNonce: 0xa8198f3c\nIssued At: 2026-06-29T10:12:21.185Z"
	signature := "0x4f34dfef91a97acc7fbfbad0b32e289925ebf0ec9f506b3bbb3e53527154c2162cccda5ad350d8aa9778c2dd9ccb3423408c05aeaf1296442f9d07d8351beee81c"
	originalHash := accounts.TextHash([]byte(messageStr))
	t.Logf("safeHash:0x%x", originalHash)
	b, err := VerifySafeSignature(RPC, "0xf7874e8076Bf31519f54B72d2dD826ed3136CC12", fmt.Sprintf("0x%x", originalHash), signature)
	t.Log(b)
	t.Log(err)
}

func TestVerify5(t *testing.T) {
	SafeMessage := "0x189b396b1cf8609e616defc9edd14d5b2c2d58e7aaffef8a3b681bcc7941d3fc"
	signature := "0x4f34dfef91a97acc7fbfbad0b32e289925ebf0ec9f506b3bbb3e53527154c2162cccda5ad350d8aa9778c2dd9ccb3423408c05aeaf1296442f9d07d8351beee81c"
	b, err := VerifySafeSignature(RPC, "0x916E4FE9C7b277828815f94f79d9E15dBFD37083", SafeMessage, signature)
	t.Log(b)
	t.Log(err)
}

func TestVerify6(t *testing.T) {
	rawMessage := "vaults.localhost wants you to sign in with your Ethereum account:\n0x916E4FE9C7b277828815f94f79d9E15dBFD37083\n\nI agree to Bedrock Vault ToS v1.0, https://legal.bedrock.technology/tos-selini-v1\n\nURI: https://vaults.localhost\nVersion: 1\nChain ID: 1\nNonce: 0xa8198f3c\nIssued At: 2026-06-29T10:12:21.185Z"
	safeAddressHex := "0x916E4FE9C7b277828815f94f79d9E15dBFD37083"
	chainID := big.NewInt(1)

	safeAddress := common.HexToAddress(safeAddressHex)

	messageHash := HashSafeMessageString(rawMessage)
	fmt.Printf("1. Message Hash (EIP-191): 0x%x\n", messageHash)

	safeMessageHash := GetSafeMessageHash(safeAddress, chainID, messageHash)
	fmt.Printf("2. Safe Message Hash (EIP-712): 0x%x\n", safeMessageHash)
}

func HashSafeMessageString(message string) []byte {
	return accounts.TextHash([]byte(message))
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
