
.PHONY: all build clean

all: build

build:
	go build -v -o ./bin/nagbot ./cmd/nagbot 

clean:
	rm -r bin
