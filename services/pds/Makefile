NAME=pds
VERSION=v0

docker-server:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	docker build -t $(NAME):$(VERSION) .
	rm main