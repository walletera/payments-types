# payments-types

What's on this repo
- the Payments API OpenApi specification
- a generated Golang client and server
- the Payments Events API

## How to generate the Golang client

To generate the golang client code from the openapi specification you need to install [ogen](https://ogen.dev/) by executing in a terminal
```bash
go install -v github.com/ogen-go/ogen/cmd/ogen@latest
```

Once you have the generator installed open a terminal an execute the following command
```bash
go generate ./...
```
