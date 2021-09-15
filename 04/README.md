# 04

## How to:

1. Calendar

```
$ vi main.go
func main() {
	Calendar()
}
$ go run *.go
```

2. Docker

```
$ vi main.go
func main() {
	Docker()
}
$ go run *.go ps
```

3. Tree

```
$ vi main.go
func main() {
	CmdTree()
}
$ go build -o tree *.go
$ ./tree --help
Usage of ./tree:
  -L int
        Descend only level directories deep.
  -d string
        Directory (default ".")
```
