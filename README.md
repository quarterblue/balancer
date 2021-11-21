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
$ go mod download
$ go build -o /balancer
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

## References

- <a href="https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf">Chord: A Scalable Peer-to-peer Lookup Protocol for Internet Applications</a>
- <a href="https://ai.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html">Consistent Hashing with Bounded Loads </a>
- <a href="https://arxiv.org/pdf/1908.08762.pdf">Revisiting Consistent Hashing With Bounded Loads</a>
- <a href="https://www.gsd.inesc-id.pt/~jgpaiva/pubs/nca13.pdf">Rollerchain: a DHT for Efficient Replication</a>

## License

Licensed under the MIT License.
