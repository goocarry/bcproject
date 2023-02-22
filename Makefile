build:
	go build -o ./bin/bcproject

run: build
	./bin/bcproject

test: 
	go test -v ./...