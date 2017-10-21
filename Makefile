## env
CGO_ENABLED=0
GOOS=linux
GO=go

##
all: build

build:
	CGO_ENABLED=$(CGO_ENABLED) GOPATH=$(GOPATH) $(GO) install game24; \
	mkdir -p bin; \
	cp -f $(GOPATH)/bin/* bin/ ;

clean:
	rm -rf $(GOPATH)/pkg
	rm -rf libs/pb/*.pb.go
