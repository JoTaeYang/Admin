set GOOS=linux
set CGO_ENABLED=0
set GOARCH=amd64

go build -ldflags="-s -w" -trimpath -o admin-back.exe

docker buildx build --platform linux/amd64 -t admin-back -f Dockerfile.amd64 .