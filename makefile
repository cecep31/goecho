bin:
	@mkdir bin
run:
	@go build -o bin
	@./bin/goecho
build:
	@go build -o bin