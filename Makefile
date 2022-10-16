default: run

run:
	go run hello.go

dist:
	rm -rf hello
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello -ldflags "-s -w" -trimpath hello.go

docker:
	docker image build -t shyaminayesh/hello .