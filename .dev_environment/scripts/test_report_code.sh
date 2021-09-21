go test $(go list ./... | grep -v mocks | grep -v main) -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
rm coverage.out