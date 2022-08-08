# Tutorial

In this repo, each section has its own `main` package.
But most won't run anything because the real learning is in the tests for each section.

Run a single test

```
go test ./03_integers
```

Run all tests

```
go test ./...
```

## Learn Go with Tests

[https://github.com/quii/learn-go-with-tests/blob/master/README.md](https://github.com/quii/learn-go-with-tests/blob/master/README.md)

## General getting started

[Getting started on go.dev](https://go.dev/doc/tutorial/getting-started)

```
$ go mod init example/hello
go: creating new go.mod: module example/hello
```

**External package**

After importing a new external package, add as requirement to go.sum

```
# Assume rsc.io/quote was added to project like this...
# import "rsc.io/quote"

# Then ...
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```
