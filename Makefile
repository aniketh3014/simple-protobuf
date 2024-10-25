gen:
	protoc -I=protos --go_out=pb --go_opt=paths=source_relative protos/*.proto

run:
	go run main.go

clean:
	rm -rf pb/*.go

test:
	go test -race ./...