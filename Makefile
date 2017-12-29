clean:
ifneq ($(shell docker network ls | grep bayamo_network | wc -l | xargs), 0)
	docker network rm bayamo_network
endif

build:
	docker build -t itstommy/bayamo .

create_network:
ifeq ($(shell docker network ls | grep bayamo_network | wc -l | xargs), 0)
	docker network create bayamo_network
endif

create_server: build
	docker run --rm \
		--name bayamo_server_tmp \
		--env-file ${PWD}/.env \
		-e GOOS=darwin \
		-td itstommy/bayamo
	docker exec -t bayamo_server_tmp go build -o bin/run cmd/main.go
	docker cp bayamo_server_tmp:/go/src/github.com/shavit/bayamo/bin bin
	docker stop bayamo_server_tmp

start_dev: create_network
	docker run --rm \
		--net bayamo_network \
		--env-file ${PWD}/.env \
		-v ${PWD}:/go/src/github.com/shavit/bayamo \
		-ti itstommy/bayamo

start_server: create_network
	docker run --rm \
		--net bayamo_network \
		--name bayamo_server \
		--env-file ${PWD}/.env \
		-v ${PWD}:/go/src/github.com/shavit/bayamo \
		-ti itstommy/bayamo go run cmd/main.go server

protoc:
	export GOPATH=~/Go
	export PATH=${PATH}:${GOPATH}/bin
	protoc -I . proto/bayamo.proto --go_out=plugins=grpc:.
