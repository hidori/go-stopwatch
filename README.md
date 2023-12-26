# Go Stopwatch

Stopwatch package for Go

## INSTALL

```bash
go get -u github.com/hidori/go-stopwatch
```

## USAGE

```go
sw := stopwatch.NewStopwatch()

sw.Start()
time.Sleep(1 * time.Second)
fmt.Println(sw.Duration())
// about 1.0s

time.Sleep(1 * time.Second)
fmt.Println(sw.Duration())
// about 2.0s

sw.Stop()
time.Sleep(1 * time.Second)
fmt.Println(sw.Duration())
// about 2.0s

time.Sleep(1 * time.Second)
fmt.Println(sw.Duration())
// about 2.0s
```
