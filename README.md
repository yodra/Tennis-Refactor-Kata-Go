# Tennis in [Go](http://golang.org/)

## Installation for newer versions of Go

Newer versions support go modules. In that case the subdirectory 'tenniskata' is a go module and you can just use it as is.

## Installation for older versions of Go

Assuming you have a proper [workspace](http://golang.org/doc/code.html#Workspaces) set up, run
```
go get github.com/yodra/Tennis-Refactoring-Kata-Go/tenniskata
```

## Running Tests

On the command line, enter the ```.../Tennis-Refactoring-Kata-Go``` directory and run
```
go test ./...
```

## Instructions

1. Choose a game (1, 2 or 3)
2. Refactor keeping the tests green
3. Refactor tests to use testify library
4. Create a PR with the team name