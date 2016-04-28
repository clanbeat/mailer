default: clean build

run: build
	PORT=2200 \
	./mailer

build: fmt
	@go build

clean:
	@rm -rf mailer dist out

fmt:
	@gofmt -w -s *.go
	@gofmt -w -s */*/*.go
