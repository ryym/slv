package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ryym/slv/cli"
)

const SRC = `package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		name := sc.Text()
		fmt.Printf("Hi, %s\n", name)
	}
}
`

const TEST = `
[[test]]
in = "Bob"
out = "Hi, Bob"
`

func TestBasicFlow(t *testing.T) {
	dir, err := ioutil.TempDir("", "slv")
	fatalIf(err, t)
	defer os.RemoveAll(dir)

	err = os.Chdir(dir)
	fatalIf(err, t)

	app := cli.CreateApp()

	err = app.Run([]string{"slv", "new", "hello"})
	fatalIf(err, t)

	err = os.Chdir(filepath.Join(dir, "hello"))
	fatalIf(err, t)

	err = ioutil.WriteFile(filepath.Join("src", "hello.go"), []byte(SRC), 0644)
	fatalIf(err, t)

	err = ioutil.WriteFile(filepath.Join("test", "t.toml"), []byte(TEST), 0644)
	fatalIf(err, t)

	err = app.Run([]string{"slv", "test", "go"})
	fatalIf(err, t)
}

func fatalIf(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
