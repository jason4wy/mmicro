# Hello Service

This is the Hello service

Generated with

```
micro new hello/srv --namespace=io.github.jason4wy --alias=hello --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: io.github.jason4wy.srv.hello
- Type: srv
- Alias: hello

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./hello-srv
```

Build a docker image
```
make docker
```