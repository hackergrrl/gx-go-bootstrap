# gx-go-bootstrap

> Generate bootstrapping scripts for gx projects written in Go


[gx](https://github.com/whyrusleeping/gx) is a package manager build on
[IPFS][]. Go projects can use [gx-go](https://github.com/whyrusleeping/gx-go) to
manage their dependencies.

However, what about bootstrapping? Using gx in Go makes your project
incompatible with `go get`-style installation. It'd be nice if there was tooling
to make this simple: enter `gx-go-bootstrap`!

`gx-go-bootstrap` is a command that will generate a Makefile and some support
scripts. It's entirely self contained, and has everything end-users need to
build or install your gx project.


## Example

Let's start a new Go project with gx, and add bootstrapping scripts:

```
$ cat > hello.go
package main

import "fmt"

func main () {
  fmt.Println("hello gx world")
}
^D

$ gx init --lang=go

$ gx publish
package hello published with hash: QmckcuLdy2KVWNJgVXjVYS3KNboWoV81AkJ92fyQVGZ6Wr

$ gx-go-bootstrap
writing Makefile
writing bin/check_go_path
writing bin/check_go_version
writing bin/check_gx_program
writing bin/dist_get
```

Now users can install your Go project by running

```
$ go get -d github.com/noffle/hello

$ cd ${GOPATH}/src/github.com/noffle/hello

$ make install
```


## Install

With Go installed, run

```
$ go get github.com/noffle/gx-go-bootstrap
```

The bootstrapping scripts will download `gx` and `gx-go` as needed on the fly.


## See Also

- [`gx`](https://github.com/whyrusleeping/gx)
- [`gx-go`](https://github.com/whyrusleeping/gx-go)
- [IPFS][]


## License

MIT


[IPFS]: https://ipfs.io
