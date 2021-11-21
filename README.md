# balancer

<a href="https://github.com/quarterblue/balancer/actions/workflows/go.yml" target="_blank">
  <img src="https://github.com/quarterblue/balancer/actions/workflows/go.yml/badge.svg" alt="GitHub Passing">
</a>
<a href="https://github.com/quarterblue/pbalancer/blob/main/LICENSE" target="_blank">
  <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License">
</a>

Balancer is a fault-tolerant implementation of a modified <a href="https://pdos.csail.mit.edu/papers/ton:chord/paper-ton.pdf">Chord distributed hash tables</a>. It can be used as a standalone load balancer for distributed systems or a fully functional kv storage (with interactive cli). As a load balancer library, the keys are distributed using <a href="https://ai.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html">consistent hashing with bounded loads</a> (<a href="https://arxiv.org/pdf/1908.08762.pdf">w/ free random jumps</a>) to avoid cascading failure. As a kv storage, an interactive cli provides easy access. It can be used as a library to build decentralized peer-to-peer application on top of the chord ring. It exposes customizable replication factor and implements <a href="https://www.gsd.inesc-id.pt/~jgpaiva/pubs/nca13.pdf">multiple chords rings</a> to provide fault-tolerance.

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
