gofmt -s -w . && go test $(go list ./... | grep -v mocks | grep -v main) -cover