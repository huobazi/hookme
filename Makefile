
build:
	@go build  -o dist/hookme -ldflags "`govvv -flags`" cmd/hookme/main.go