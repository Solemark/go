# Go Meta Repo

run code with:

    go run name.go

run tests with:

    go test
    OR
    go test -v

build with:
	
	go build name.go
	OR
	go build -ldflags -w name.go

-ldflags -w removes debug symbols resulting in a smaller binary

target another system with:

	GOOS=platform GOARCH=system go build name.go
	e.g
	GOOS=windows GOARCH=amd64 go build name.go
