package airdrop

import (
	"encoding/hex"
	"testing"
)

func TestGetMerkleRootAndProof(t *testing.T) {
	leaves := []TreeLeaf{
		{Address: "0x1234567890abcdef1234567890abcdef12345678", Amount: "100000000"},
		{Address: "0x1bcdef1234567890abcdef1234567890abcdef12", Amount: "200000000"},
		{Address: "0x1890abcdef1234567890abcdef1234567890abcd", Amount: "300000000"},
		{Address: "0xabcdef1234567890abcdef1234567890abcdef12", Amount: "400000000"},
		{Address: "0x5Ea021b0F5B814CAD162F350f01324c121142223", Amount: "500000000"},
	}
	tree, err := GetMerkleRootAndProof(leaves)
	if err != nil {
		t.Fatalf("GetMerkleRootAndProof returned error: %v", err)
	}
	t.Log("root:", hex.EncodeToString(tree.GetRoot()))
	proof, err := tree.TreeProofMarshal()
	if err != nil {
		t.Fatalf("TreeProofMarshal returned error: %v", err)
	}
	t.Log("proof:", string(proof))
}
