

fmt:
	gofmt -w . && goimports -w . 

test:
	go test -v ./... 

cover:
	go test -coverprofile=test-coverage.out && go tool cover -html=test-coverage.out