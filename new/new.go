package new

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"

	tmpl "github.com/go-tars/tars/template"
	"github.com/urfave/cli/v2"
	"github.com/xlab/treeprint"
)

type file struct {
	Path string
	Tmpl string
}

type config struct {
	App     string
	Server  string
	Servant string
	Module  string
	// tars new example -type
	Command string
	// github.com/go-tars/foo
	Dir string
	// $GOPATH/src/github.com/go-tars/foo
	GoDir string
	// $GOPATH
	GoPath string
	// UseGoPath
	UseGoPath bool
	// Files
	Files []file
}

func write(c config, file, tmpl string) error {
	fn := template.FuncMap{
		"title": func(s string) string {
			return strings.ReplaceAll(strings.Title(s), "-", "")
		},
		"dehyphen": func(s string) string {
			return strings.ReplaceAll(s, "-", "")
		},
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("f").Funcs(fn).Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}

func create(c config) error {
	// check if dir exists
	if _, err := os.Stat(c.GoDir); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", c.GoDir)
	}

	fmt.Printf("Creating service in %s\n\n", c.GoDir)

	t := treeprint.New()

	// write the files
	for _, file := range c.Files {
		f := filepath.Join(c.GoDir, file.Path)
		dir := filepath.Dir(f)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}

		addFileToTree(t, file.Path)
		if err := write(c, f, file.Tmpl); err != nil {
			return err
		}
	}

	// print tree
	fmt.Println(t.String())

	// just wait
	<-time.After(time.Millisecond * 250)

	return nil
}

func addFileToTree(root treeprint.Tree, file string) {

	split := strings.Split(file, "/")
	curr := root
	for i := 0; i < len(split)-1; i++ {
		n := curr.FindByValue(split[i])
		if n != nil {
			curr = n
		} else {
			curr = curr.AddBranch(split[i])
		}
	}
	if curr.FindByValue(split[len(split)-1]) == nil {
		curr.AddNode(split[len(split)-1])
	}

}

func Run(ctx *cli.Context) {
	app := ctx.String("app")
	server := ctx.String("server")
	servant := ctx.String("servant")
	dir := ctx.Args().First()
	module := dir
	useGoPath := ctx.Bool("gopath")
	useGoModule := os.Getenv("GO111MODULE")

	if len(app) == 0 {
		fmt.Println("app not defined")
		return
	}

	if len(server) == 0 {
		fmt.Println("server not defined")
		return
	}

	if len(servant) == 0 {
		fmt.Println("servant not defined")
		return
	}

	command := "tars new "

	var goPath string
	var goDir string

	// only set gopath if told to use it
	if useGoPath {
		goPath = build.Default.GOPATH
		fmt.Println("gopath", goPath)

		// don't know GOPATH, runaway....
		if len(goPath) == 0 {
			fmt.Println("unknown GOPATH")
			return
		}

		// attempt to split path if not windows
		if runtime.GOOS == "windows" {
			goPath = strings.Split(goPath, ";")[0]
		} else {
			goPath = strings.Split(goPath, ":")[0]
		}
		goDir = filepath.Join(goPath, "src", path.Clean(dir))
	} else {
		goDir = path.Clean(dir)
	}

	c := config{
		App:       app,
		Server:    server,
		Servant:   servant,
		Module:    module,
		Command:   command,
		Dir:       dir,
		GoDir:     goDir,
		GoPath:    goPath,
		UseGoPath: useGoPath,
	}

	// create service config
	c.Files = []file{
		{"main.go", tmpl.Main},
		{"config.conf", tmpl.Config},
		{"client/client.go", tmpl.Client},
		{c.Servant + "_imp.go", tmpl.ServantImp},
		{c.Servant + ".tars", tmpl.ServantTars},
		{"start.sh", tmpl.StartSh},
		{"test/test.go", tmpl.Test},
		{"Makefile", tmpl.Makefile},
		{".gitignore", tmpl.GitIgnore},
	}

	// set gomodule
	if useGoModule != "off" {
		c.Files = append(c.Files, file{"go.mod", tmpl.Module})
	}

	if err := create(c); err != nil {
		fmt.Println(err)
		return
	}
}

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "new",
			Usage:       "Create a service template",
			Description: `'tars new' scaffolds a new service skeleton. Example: 'tars new my-app && cd my-app'`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "app",
					Aliases: []string{"a"},
					Usage:   "AppName.",
					Value:   "tars",
				},
				&cli.StringFlag{
					Name:    "server",
					Aliases: []string{"s"},
					Usage:   "ServerName.",
					Value:   "DemoServer",
				},
				&cli.StringFlag{
					Name:    "servant",
					Aliases: []string{"st"},
					Usage:   "ServantName.",
					Value:   "DemoServant",
				},
				&cli.BoolFlag{
					Name:    "gopath",
					Aliases: []string{"g"},
					Usage:   "Create the service in the gopath.",
				},
			},
			Action: func(c *cli.Context) error {
				Run(c)
				return nil
			},
		},
	}
}
