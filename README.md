# my-grpc

## Benchmark

```
$ go test -bench . --benchmem                                                                                                                            master 
BenchmarkGRPC-4            10000            133770 ns/op            2868 B/op         63 allocs/op
BenchmarkREST-4            10000            105410 ns/op            3891 B/op         55 allocs/op
PASS
ok      github.com/Jun-Chang/my-grpc/client     2.452s
```
