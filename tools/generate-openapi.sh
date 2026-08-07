#!/bin/sh

set -eu

tools_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
repository_dir=$(dirname -- "$tools_dir")
go_bin=$(go env GOPATH)/bin

if [ ! -x "$go_bin/goctl" ]; then
  echo "missing $go_bin/goctl; install it with:" >&2
  echo "  go install github.com/zeromicro/go-zero/tools/goctl@v1.10.1" >&2
  exit 1
fi
if [ ! -x "$go_bin/oapi-codegen" ]; then
  echo "missing $go_bin/oapi-codegen; install it with:" >&2
  echo "  go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest" >&2
  exit 1
fi

cd "$repository_dir/api"
"$go_bin/goctl" api swagger -api cuniBTCReward.api -dir docs
go run ../tools/openapi3 -in docs/cuniBTCReward.json -types internal/types/types.go
cd openapi
"$go_bin/oapi-codegen" -config oapi-codegen.yaml ../docs/cuniBTCReward.json
