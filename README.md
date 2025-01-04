# aws-go-gin-poc
Proof of concept with AWS Lambda, Golang, Gin

## Prepare build
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap cmd/main.go

## zip file to upload
zip mcclAPI3.zip bootstrap

