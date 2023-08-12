package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Options is used to define version cmd options
type Options struct {
}

// NewOptions is used to create a new Options instance
func NewOptions() *Options {
	return &Options{}
}

// NewVersionCmd is used to create a new version cmd instance
func NewVersionCmd() *cobra.Command {
	o := NewOptions()

	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ekko",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Run()
		},
	}
}

func (o *Options) Run() error {
	fmt.Println("ekko version 0.0.1")

	return nil
}
