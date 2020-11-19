package template

var (
	Makefile = `all: build run

build:
	tars2go -outdir=tars-protocol -module=github.com/go-tars/demo *.tars
	go build -o {{.Server}}

run:
	./{{.Server}} --config=config.conf

run_test:
	go run test/test.go --config config.conf
`
)
