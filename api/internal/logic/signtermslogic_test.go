package logic

import (
	"fmt"
	"testing"

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
