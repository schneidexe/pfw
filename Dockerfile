FROM golang

ENV USER root

WORKDIR /go/src/github.com/schneidexe/pfw

ADD . /go/src/github.com/schneidexe/pfw

RUN mkdir bin

CMD for OS in darwin linux windows; do \
        export GOOS=$OS; \
        for ARCH in 386 amd64; do \
            export GOARCH=$ARCH; \
            echo building $GOOS-$GOARCH; \
            if [ $GOOS = "windows" ]; then \
                go build -o bin/pfw-$GOOS-$GOARCH.exe; \
            else \
                go build -o bin/pfw-$GOOS-$GOARCH; \
            fi \
        done; \
    done
