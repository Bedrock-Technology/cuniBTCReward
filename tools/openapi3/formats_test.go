package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFormatsFromGoFileIncludesEmbeddedFields(t *testing.T) {
	path := filepath.Join(t.TempDir(), "types.go")
	source := `package types
type Base struct {
	Epoch uint64 ` + "`json:\"epoch\"`" + `
}
type Response struct {
	Base
	Total int64 ` + "`json:\"total\"`" + `
	APY float64 ` + "`json:\"apy\"`" + `
}
`
	if err := os.WriteFile(path, []byte(source), 0o600); err != nil {
		t.Fatal(err)
	}

	formats, err := formatsFromGoFile(path)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"epoch": "uint64", "total": "int64", "apy": "double"}
	for field, expected := range want {
		if actual := formats["Response"][field]; actual != expected {
			t.Errorf("Response.%s format = %q, want %q", field, actual, expected)
		}
	}
}
