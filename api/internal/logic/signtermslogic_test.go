package logic

import (
	"cuniBTCReward/api/internal/types"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	fmt.Println(verifySig(
		"0xac07f2721EcD955c4370e7388922fA547E922A4f",
		"0x5da3c3d911791822a05b819834ef508b48cf69ecc4d3a9eea69e0938a51f98e876dbae6514288f2b4b46302749470d5727c9f679616c2ae0340b71691c9fb2d31c",
		[]byte("Example `personal_sign` message"),
	))
}

func TestSign2(t *testing.T) {
	fmt.Println(verifySig(
		"0xac07f2721EcD955c4370e7388922fA547E922A4f",
		"0xaa0bdb820c6b41d86064a079bd698b2280b896e91c7a92e29651a0c15cc91a333cadb8f89da4d16e371267dd87788b5f9e25fca809b8b7c6747b3815bdaf102e1c",
		[]byte(`{\"address\":\"0xac07f2721EcD955c4370e7388922fA547E922A4f\",\"nonce\":0,\"content\":\"agree to Bedrock’s Terms of Service\",\"expireTime\":1779169253}`),
	))
}

func TestSign3(t *testing.T) {
	message := "{\"address\":\"0xac07f2721EcD955c4370e7388922fA547E922A4f\",\"nonce\":0,\"content\":\"agree to Bedrock’s Terms of Service\",\"expireTime\":1779169253}"
	fmt.Println(verifySig(
		"0xac07f2721EcD955c4370e7388922fA547E922A4f",
		"0xaa0bdb820c6b41d86064a079bd698b2280b896e91c7a92e29651a0c15cc91a333cadb8f89da4d16e371267dd87788b5f9e25fca809b8b7c6747b3815bdaf102e1c",
		[]byte(message),
	))
}
func TestSign4(t *testing.T) {
	message := types.Message{
		Address:    "0xac07f2721EcD955c4370e7388922fA547E922A4f",
		Nonce:      0,
		Content:    "agree to Bedrock’s Terms of Service",
		ExpireTime: 1779169253,
	}
	messageJson, _ := json.Marshal(&message)
	req := types.SignTermsReq{
		Message:   string(messageJson),
		Signature: "0x4c46904b62a8889db17b3d84b129eea1283e6313f6ca635802e95d4da72bf06f74b270b1fd28de9f0931cf0ca079229d5088743ae117e743a17af2205b0dda6a1c",
	}
	fmt.Printf("message:%s\n", req.Message)
	reqJson, _ := json.Marshal(req)
	fmt.Printf("req:%s\n", string(reqJson))

	fmt.Println(verifySig(
		message.Address,
		req.Signature,
		[]byte(req.Message),
	))
}
