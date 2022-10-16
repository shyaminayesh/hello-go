default: run

run:
	go run hello.go

dist:
	rm -rf hello
	go build -o hello -ldflags "-s -w" -trimpath hello.go