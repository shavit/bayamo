build:
	docker build -t itstommy/bayamo .

start_dev:
	docker run --rm \
		--env-file ${PWD}/.env \
		-v ${PWD}:/go/src/github.com/shavit/bayamo \
		-ti itstommy/bayamo

start_server:
	docker run --rm \
		--name bayamo_server \
		--env-file ${PWD}/.env \
		-ti itstommy/bayamo go run cmd/main.go server

protoc:
	export GOPATH=~/Go
	export PATH=$PATH:$GOPATH/bin
	protoc -I . proto/bayamo.proto --go_out=plugins=grpc:.
