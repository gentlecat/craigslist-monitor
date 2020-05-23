fmt :
	go fmt -x ./...

test : fmt
	go test ./... -bench .

run : fmt
	go run main.go

build-monitor : fmt
	go build -o ./build/monitor ./cmd/monitor/monitor.go
