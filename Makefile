.PHONY: build run

NAME = heimdall

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
