# Go Meta Repo

get submodule code:

	git submodule update --init --recursive

run code with:

    go run name.go

run tests with:

    go test
    OR
    go test -v

build with:
	
	go build name.go

OR

target another system with:

	GOOS=platform GOARCH=system go build name.go
	e.g
	GOOS=windows GOARCH=amd64 go build name.go
