package slv

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-test/deep"
	"github.com/ryym/slv/slv/tp"
)

func TestNewProblem(t *testing.T) {
	dir, err := ioutil.TempDir("", "slv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	err = NewProblem(&tp.CmdNewOpts{Name: "hello"})
	if err != nil {
		t.Fatal(err)
	}

	fs, err := ioutil.ReadDir(filepath.Join(dir, "hello"))
	if err != nil {
		t.Fatal(err)
	}

	names := make([]string, len(fs))
	for i, f := range fs {
		names[i] = f.Name()
	}

	expected := []string{".slv", "src", "test"}
	if diff := deep.Equal(names, expected); diff != nil {
		t.Error(diff)
	}
}
