project_name = atstash
package = github.com/lenfree/$(project_name)

all: release

.PHONY: install
install:
	go get -v

.PHONY: script
script:
	go vet ./...
	go test -v -race ./...

.PHONY: release
release: install script
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(project_name)-linux-amd64 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/$(project_name)-darwin-amd64 $(package)
