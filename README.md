# tars


## Install

```
go get github.com/go-tars/tars@master
```

Also required:

- [tars2go](https://github.com/TarsCloud/TarsGo)
```
go install $GOPATH/src/github.com/TarsCloud/TarsGo/tars/tools/tars2go
```

## Usage

1. new a demo
```bash
tars new --app itars --server DemoServer --servant DemoApi --gopath github.com/go-tars/demo
```

- [app](https://tarscloud.github.io/TarsDocs/base/tars-concept.html#main-chapter-1)
- [server](https://tarscloud.github.io/TarsDocs/base/tars-concept.html#main-chapter-2)
- [servant](https://tarscloud.github.io/TarsDocs/base/tars-concept.html#main-chapter-3)


2. start demo server 

```bash
cd $GOPATH/src/github.com/go-tars/demo
make
```

3. test

```
make test
```
