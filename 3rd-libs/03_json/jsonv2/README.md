# JSON v2 Benchmark

JSON serialization/deserialization benchmark comparing three implementations:

- `encoding/json` (classic standard library)
- `encoding/json/v2` (new standard library, Go 1.24+)
- `github.com/bytedance/sonic` (third-party, JIT-accelerated)

The benchmark reuses the `Employee` data model (`BasicInfo` + `JobInfo`) across
all implementations, and is driven by `json_v2_test.go`.

## Run

```bash
go test -bench=. -benchmem -run=^$ .
```

## Environment

- Go: go1.27.0 (darwin/arm64, Apple M5 Pro)
- sonic: v1.15.2

> **WARNING**: sonic/ast only supports (go1.17~1.26 and amd64 CPU) or
> (go1.20~1.26 and arm64 CPU). Under Go 1.27 the sonic benchmarks fall back to
> `encoding/json`, so its results below reflect the fallback path rather than
> sonic's JIT-accelerated fast path.

## Benchmark Results

### Unmarshal (decode)

| Benchmark | Time/op | Bytes/op | Allocs/op |
|---|---:|---:|---:|
| BenchmarkStdUnmarshal | 419.0 ns | 256 B | 5 |
| BenchmarkV2Unmarshal | 335.1 ns | 256 B | 5 |
| BenchmarkSonicUnmarshal | 772.9 ns | 1081 B | 14 |
| BenchmarkSonicFastestUnmarshal | 762.8 ns | 1081 B | 14 |
| BenchmarkSonicNoValidateUnmarshal | 751.5 ns | 985 B | 13 |

### Marshal (encode)

| Benchmark | Time/op | Bytes/op | Allocs/op |
|---|---:|---:|---:|
| BenchmarkStdMarshal | 188.5 ns | 80 B | 1 |
| BenchmarkV2Marshal | 182.5 ns | 80 B | 1 |
| BenchmarkSonicMarshal | 227.3 ns | 240 B | 3 |
| BenchmarkSonicFastestMarshal | 229.2 ns | 240 B | 3 |

## Observations

- `encoding/json/v2` is faster than the classic `encoding/json` on both
  Unmarshal (~20%) and Marshal (~3%) with identical allocations.
- Under Go 1.27, sonic falls back to the standard library and is therefore
  slower and allocates more than the native libraries in this run. To observe
  sonic's real JIT advantage, run on a supported toolchain (Go <= 1.26).
