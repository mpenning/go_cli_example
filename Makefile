all:
	# Start go module configuration from scratch...
	-rm go.mod
	-rm go.sum
	-rm pinger
	-rm pinger-static
	go mod init pinger
	# Download the required dependencies at specified versions...
	go get github.com/gookit/slog@v0.5.8
	go get github.com/prometheus-community/pro-bing@v0.7.0
	go get github.com/urfave/cli/v3
	go mod tidy
	go fmt
	go vet
	# Buld a statically-linked binary named ./pinger-static
	CGO_ENABLED=0 go build -ldflags="-extldflags -static" -v -o pinger-static
	# Buld a dynamically-linked binary named ./pinger
	go build -v -o pinger main.go
