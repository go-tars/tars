package template

var (
	StartSh = `#!/bin/bash
make
./{{.Server}} --config=config.conf
`
)
