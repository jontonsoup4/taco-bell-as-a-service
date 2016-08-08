NAME=$(shell basename $(CURDIR))

DEPS=github.com/oleiade/reflections github.com/bradfitz/slice github.com/gorilla/mux

default: run

build:
	go build -o build/${NAME}


clean:
	rm -rf build

install:
	go get -v $(DEPS)

run:
	if [ -f build/${NAME} ]; then echo "Running from build..."; build/${NAME}; else echo "Serving..."; go run *.go; fi;
