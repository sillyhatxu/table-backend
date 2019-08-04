go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
docker build -t xushikuan/temp-backend .
docker push xushikuan/temp-backend:1.6
go clean