all:
	CGO_ENABLED=0 go build -o alexa-parrot .

# Used for CI publishingz
docker-push:
	docker build -t euank/alexa-parrot:$(shell git rev-parse --short HEAD)
