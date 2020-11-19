# tars


## Install

```
go get github.com/go-tars/tars@master
```

Also required:

- [protoc](https://github.com/google/protobuf)
- [protoc-gen-go](https://github.com/golang/protobuf)

## Usage

```bash
tars new --app tars --server DemoServer --servant DemoApi --gopath github.com/go-tars/demo
cd $GOPATH/github.com/go-tars/demo
./start.sh
```

- [app](https://tarscloud.github.io/TarsDocs/base/tars-concept.html#main-chapter-1)
- [server](https://tarscloud.github.io/TarsDocs/base/tars-concept.html#main-chapter-2)
- [servant](https://tarscloud.github.io/TarsDocs/base/tars-concept.html#main-chapter-3)