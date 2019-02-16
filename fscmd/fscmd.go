package fscmd

import (
	"github.com/autom8ter/fsctl"
	"github.com/autom8ter/fsctl/clone"
	"github.com/spf13/cobra"
)

type FsCmd struct {
	*fsctl.Fs
	*cobra.Command
}

func NewFsCmd(name, usg string, c clone.CloneFunc) *FsCmd {
	fs, err := fsctl.NewFs(c)
	if err != nil {
		panic(err)
	}
	return &FsCmd{
		Fs: fs,
		Command: &cobra.Command{
			Use:   name,
			Short: usg,
		},
	}
}

func(f *FsCmd) Init() {
	for _, c := range f.Commands() {
		_ = f.BindPFlags(c.Flags())
		_ = f.BindPFlags(c.PersistentFlags())
	}
}