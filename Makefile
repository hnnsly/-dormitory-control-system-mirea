build:
	go build -o bin/dorm

run: build
	./bin/dorm

test:
	@go test -v ./..