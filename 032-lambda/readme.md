GOARCH=amd64 GOOS=linux go build main.go
zip function.zip main
aws lambda update-function-code --function-name Cron_GYB1714 --zip-file fileb://function.zip