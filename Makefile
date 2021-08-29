.PHONY:all
all: build

build:
	go build -o bin/papers-map main.go
