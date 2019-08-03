go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
docker build -t xushikuan/temp-backend .
docker tag xushikuan/temp-backend:latest xushikuan/temp-backend:1.5
docker push xushikuan/temp-backend:1.5
go clean