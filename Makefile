.PHONY: build run

NAME = heimdall

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

install:
	glide install

docker-build:
	docker build -t $(NAME) .

docker-run:
	docker run -it --rm --name $(NAME) -p 80:80 $(NAME)

docker-rm:
	docker rm $(NAME)

docker-rmi:
	docker rmi $(NAME)

default: docker-build
