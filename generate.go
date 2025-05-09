package payments

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --package publicapi --target publicapi --clean openapi/public-api.yaml
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --package privateapi --target privateapi --clean openapi/private-api.yaml
