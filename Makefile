LINTER_CFG = ./config-dev/.golangci.yml
PKG = ./...

# DEV
create-env:
	go mod download

run-linter:
	$(GOPATH)/bin/golangci-lint run $(PKG) --config=$(LINTER_CFG)
	go fmt $(PKG)

build:
	rm -f bin/main.bin
	go build -o bin/main.bin cmd/app/main.go

# easyjson
generate:
	go generate ${PKG}

clear:
	sudo rm -rf fullchain.pem privkey.pem bin/*