generate:
	protoc --proto_path=proto proto/*.proto  --go_out=. --go-grpc_out=.

build-server:
	go build -o calculator-server cmd/server/main.go

build-client:
	go build -o calculator-client cmd/client/main.go

build-docker:
	docker build -t calculator-server .

run-server:
	docker run -p 2706:2706 calculator-server
