# balancer

Balancer is a fault-tolerant implementation of the Chord distributed hash table. 

## Features

- Easy to use
- Fault-tolerent with customizable replication factor
- Flexible (use as library or cli)
- Standalone load balancer
- Fully functional kv storage
- Performant gRPC networking

## Installation

To use as a `library`:

```shell
$ cd projectdir/
$ go get github.com/quarterblue/balancer
```

Import into your Go project:

```go
import (
  	"github.com/quarterblue/balancer"
)
```

To use as a `CLI key-value storage`:

```shell
$ git clone github.com/quarterblue/balancer
$ cd balancer/
$ go run main.go
```

## Usage

You can use it has an interactive Key-Value storage:

```shell
$ go run main.go -port 9001 -ring -cli
```
To store a kv pair or to retrieve: 

```shell
$ [balancer] CMD-> get compare
$ [balancer] CMD-> value: merkletree
```

## License

Licensed under the MIT License.
