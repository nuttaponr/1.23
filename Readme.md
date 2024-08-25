## TL;DR

ใน Go 1.23, `for-range` loop  ใช้ได้กับ

- `func(func() bool)`
- `func(func(K) bool)`
- `func(func(K, V) bool)`

```
go test -v .
```