set_env:
	@export GOPATH=$HOME
	@export GO111MODULE=on
	@go mod tidy

run:
	@go run cmd/main.go

compile_java_lexers:
	@javac -classpath ./parser/antlr-4.10.1-complete.jar Styx*.java