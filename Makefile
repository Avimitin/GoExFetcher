build:
	@if [ ! -e "bin" ]; then mkdir bin; fi
	go build -o ./bin/exfetcher ./cmd/cli/main.go
	@echo "build successfully. File store in ./bin"

help:
	@echo "Build binary file: make build"
	@echo "Help message: make help"

clean:
	@if [ -e "bin" ]; then rm -rf ./bin; fi
	@echo "clean done"
