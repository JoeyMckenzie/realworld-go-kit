## conduit-bin

The `conduit-bin` module is our primary application driver, instantiating all dependencies required
for our higher level policy modules (logging, DB connections, HTTP router, etc). We aim to keep
`main.go` as thin as possible, offloading "dependency builders" to their respective modules for `main.go`
to consume.