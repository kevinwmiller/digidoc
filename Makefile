GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u

OUTPUT_DIR=build
BINARY_NAME=digidoc
TARGET=$(OUTPUT_DIR)/$(BINARY_NAME)

all: test build
.PHONY: build
build:
	protoc -I server/ server/digidoc.proto --go_out=plugins=grpc:server
	$(GOBUILD) -o $(TARGET) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(TARGET)
run:
	$(GOBUILD) -o $(TARGET) -v
	./$(TARGET)
deps:
	$(GOGET) google.golang.org/grpc
	$(GOGET) github.com/golang/protobuf/protoc-gen-go