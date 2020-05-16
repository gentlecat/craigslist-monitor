fmt :
	go fmt -x ./...

test : fmt
	go test ./... -bench .

run : fmt
	go run main.go
