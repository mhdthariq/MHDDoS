# Build script for MHDDoS Go version

# Build for current platform
echo "Building for current platform..."
go build -o mhddos main.go

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o bin/mhddos-linux-amd64 main.go
GOOS=linux GOARCH=arm64 go build -o bin/mhddos-linux-arm64 main.go

# Build for Windows
echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o bin/mhddos-windows-amd64.exe main.go

# Build for macOS
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o bin/mhddos-darwin-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o bin/mhddos-darwin-arm64 main.go

echo "Build complete! Binaries are in the bin/ directory"
