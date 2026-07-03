package logic

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
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
	messageStr := "localhost:3000 wants you to sign in with your Ethereum account:\n0x916E4FE9C7b277828815f94f79d9E15dBFD37083\n\nI agree to Bedrock Vault ToS v1.0, https://legal.bedrock.technology/tos-selini-v1\n\nURI: http://localhost:3000\nVersion: 1\nChain ID: 1\nNonce: 0x3fd02540\nIssued At: 2026-07-03T06:36:07.238Z"
	signature := "0xb9c609a96fdce262684b87ac0d50384c42fdeef6ef3b267aad24bca0cb9e33982fe4c2eda8b68d44b45585fbce7b2eea1f8942bd11cbbb078617e23f18bbb0e21b8d1b86d379826671c4228f3ba9556426a65d4f0e6454e05e6733eb492d2c838d1aba22d1e37465854363021947cf1077b086bc1730560a5bb3f2d73c7e4cb5081b"
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
