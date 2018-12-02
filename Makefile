GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u

OUTPUT_DIR=build
BINARY_NAME=digidoc
TARGET=$(OUTPUT_DIR)/$(BINARY_NAME)

all: build test
.PHONY: build
build:
	protoc -I grpc/ grpc/digidoc.proto --go_out=plugins=grpc:grpc
	$(GOBUILD) -o $(TARGET) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(TARGET)
run:
	$(GOBUILD) -o $(TARGET) -v
	./$(TARGET) server start
deps:
	$(GOGET) google.golang.org/grpc
	$(GOGET) github.com/golang/protobuf/protoc-gen-go
	protoc -I grpc/ grpc/digidoc.proto --go_out=plugins=grpc:grpc
	$(GOGET)
