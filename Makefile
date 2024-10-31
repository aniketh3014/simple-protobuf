gen:
	protoc -I=protos \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		protos/*.proto

run:
	go run main.go

clean:
	rm -rf pb/*.go

test:
	go test -race ./...

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -port 0.0.0.0:8080