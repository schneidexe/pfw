# pfw
Lightweight port forwarder in go (inspired by http://blog.evilissimo.net/simple-port-fowarder-in-golang)

## setup

Just download the binary for your OS and arch from the [releases](https://github.com/schneidexe/pfw/releases) page. 

## build 
```
docker build -t pfw . && \
docker run --name pfw pfw && \
docker cp pfw:/go/src/github.com/schneidexe/pfw/bin . && \
docker rm pfw && \
docker rmi pfw
```

## test

tdb
