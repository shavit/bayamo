build:
	docker build -t itstommy/bayamo .

start_dev:
	docker run --rm \
		--env-file ${PWD}/.env \
		-v ${PWD}:/go/src/github.com/shavit/bayamo \
		-ti itstommy/bayamo
