fmt:
	gofmt -l -s -w .

build:
	cd go-src/cmd && go build -o ../../app .
