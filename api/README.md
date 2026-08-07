## OpenAPI generation

```bash
go install github.com/zeromicro/go-zero/tools/goctl@v1.10.1
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
go generate ./api/openapi
```

Run these commands from the repository root. Generation runs goctl, converts its
Swagger 2.0 output to OpenAPI 3.0.3, restores Go numeric formats, and uses
`oapi-codegen` to validate the document and regenerate
`openapi/openapi.gen.go`.

To guarantee generation happens before compilation, use `make build`. The Go
toolchain does not run `go generate` implicitly when invoking `go build`.
