DARWINARCH=arm64

darwin: ## Build for Darwin (macOS)
	GOOS=darwin GOARCH=$(DARWINARCH) go build -o handler cmd/nimbus/main.go

func: ## Build for Azure Function (Linux).
	GOOS=linux GOARCH=amd64 go build -o handler cmd/nimbus/main.go

clean: ## Remove build file.
	rm handler