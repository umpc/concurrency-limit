# ConcurrencyLimit

[![GoDoc](https://godoc.org/github.com/umpc/concurrency-limit?status.svg)](https://godoc.org/github.com/umpc/concurrency-limit)

```sh
go get -u github.com/umpc/concurrency-limit
```

Easily limit the number of running goroutines per given group of functions:

```go
const limit = 2
cl := New(limit)

for i := 0; i < limit + 1; i++ {
  go cl.Exec(func() {
    fmt.Println("Hello world!")
  })
}
```

## Benchmarks

```sh
BenchmarkNew2-8         30000000                43.4 ns/op            96 B/op          1 allocs/op
BenchmarkExec2-8         5000000               339 ns/op               0 B/op          0 allocs/op
```

The above benchmark tests were ran using a 2.6GHz Intel Core i7-6700HQ (Skylake) CPU.

## License

The source code is available under the [MIT License](https://opensource.org/licenses/MIT).
