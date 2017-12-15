NAME=kubevalonline
IMAGE_NAME=donachys/$(NAME)
PACKAGE_NAME=github.com/donachys/$(NAME)
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
TAG=0.0.0# $$(git describe --abbrev=0 --tags)

clean:
	rm -rf bin
vendor:
	echo "vendor..."
	dep ensure

darwin: vendor
	env GOOS=darwin GOAARCH=amd64 go build -v -o $(CURDIR)/bin/darwin/amd64/$(NAME) ./cmd/$(NAME)

linux: vendor
	echo "linux..."
	env GOOS=linux GOAARCH=amd64 go build -v -o $(CURDIR)/bin/linux/amd64/$(NAME) ./cmd/$(NAME)

docker:
	docker build -t $(IMAGE_NAME):$(TAG) .
	docker tag $(IMAGE_NAME):$(TAG) $(IMAGE_NAME):latest

