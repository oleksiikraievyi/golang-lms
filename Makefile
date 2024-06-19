APP_NAME=lms
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@swag init && go build -o $(APP_NAME)

.PHONY: docker-build
## docker-build: Build docker image
docker-build:
	@docker build -t $(APP_NAME):$(VERSION) .

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(APP_NAME)

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(APP_NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go mod download

.PHONY: test
## test: Run tests with verbose mode
test:
	@go test -v ./...

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo