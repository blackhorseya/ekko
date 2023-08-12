package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Options is used to define config cmd options
type Options struct {
}

// NewOptions is used to create a new Options instance
func NewOptions() *Options {
	return &Options{}
}

// NewConfigCmd is used to create config command
func NewConfigCmd() *cobra.Command {
	o := NewOptions()

	cmd := &cobra.Command{
		Use:   "config",
		Short: "config is used to manage ekko config",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "show is used to show ekko config",
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.Show()
		},
	})

	return cmd
}

func (o *Options) Show() error {
	fmt.Println("not implement yet")

	return nil
}
