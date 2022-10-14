# round-robin
round-robin is balancing algorithm written in golang

## Forked
Forked from https://github.com/hlts2/round-robin to work with strings instead of URLs

## Installation

```shell
go get github.com/hlts2/round-robin
```

## Example

```go
rr, _ := roundrobin.New(
    "192.168.33.10",
    "192.168.33.11",
    "192.168.33.12",
    "192.168.33.13",
)

rr.Next() // "192.168.33.10"
rr.Next() // "192.168.33.11"
rr.Next() // "192.168.33.12"
rr.Next() // "192.168.33.13"
rr.Next() // "192.168.33.10"
rr.Next() // "192.168.33.11"
```
## Author
[hlts2](https://github.com/hlts2)

## LICENSE
round-robin released under MIT license, refer [LICENSE](https://github.com/hlts2/round-robin/blob/master/LICENSE) file.
