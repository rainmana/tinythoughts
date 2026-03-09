.PHONY: build install clean test test-verbose test-coverage run setup-config

BINARY_NAME=tinythoughts
INSTALL_PATH=/usr/local/bin

build:
	go build -o $(BINARY_NAME) ./cmd/tinythoughts

install: build
	sudo mv $(BINARY_NAME) $(INSTALL_PATH)/

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out

test:
	go test ./...

test-verbose:
	go test ./... -v

test-coverage:
	go test ./... -cover
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

run: build
	./$(BINARY_NAME)

setup-config:
	mkdir -p ~/.config/tinythoughts/frameworks
	cp frameworks/*.yaml ~/.config/tinythoughts/frameworks/ 2>/dev/null || true
	cp tinythoughts.yaml.example ~/.config/tinythoughts/tinythoughts.yaml 2>/dev/null || true
	@echo "Configuration setup complete!"
