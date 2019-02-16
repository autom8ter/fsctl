package clone

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)
func TestClone(t *testing.T) {
	dir, err := ioutil.TempDir("", "config")
	if err != nil {
		panic(err)
	}
	c := NewAuthCloner("https://github.com/autom8ter/configs.git", "autom8ter", os.Getenv("CFGTOKEN"))
	r := c.Clone(dir)
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == ".git" {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Println(path)
		return nil
	}); err != nil {
		t.Fatal(err.Error())
	}
	if r == nil {
		t.Fatal("failed to clone mem repo")
	}
}

