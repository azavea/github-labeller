GOPATH = $(shell pwd)
GOENV = GOPATH=$(GOPATH)
BINDIR=$(GOPATH)/bin
BUILD_FLAGS = -v

all: build
release: linux darwin windows

deps:
	GOPATH=$(GOPATH) go get golang.org/x/oauth2
	GOPATH=$(GOPATH) go get github.com/BurntSushi/toml
	GOPATH=$(GOPATH) go get github.com/google/go-github/github

build: deps
	$(GOENV) go build $(BUILD_FLAGS) github-labeller.go

linux: deps
	$(GOENV) GOOS=linux go build $(BUILD_FLAGS) github-labeller.go
	mkdir -p $(BINDIR)
	mv github-labeller $(BINDIR)/github-labeller.linux

darwin: deps
	$(GOENV) GOOS=darwin go build $(BUILD_FLAGS) github-labeller.go
	mkdir -p $(BINDIR)
	mv github-labeller $(BINDIR)/github-labeller.darwin

windows: deps
	$(GOENV) GOOS=windows go build $(BUILD_FLAGS) github-labeller.go
	mv github-labeller.exe $(BINDIR)/github-labeller.exe

clean:
	rm -rf bin/
	rm -rf pkg/
	rm -rf src/
	rm -f github-labeller
	rm -f github-labeller.exe
