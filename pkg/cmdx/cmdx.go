package cmdx

import (
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServiceCmd represents the service command.
type ServiceCmd struct {
	Use        string
	Short      string
	GetService func(v *viper.Viper) (adapterx.Servicer, error)
}

// NewServiceCmd creates a new service command.
func NewServiceCmd(use string, short string, svc func(v *viper.Viper) (adapterx.Servicer, error)) *cobra.Command {
	return (&ServiceCmd{Use: use, Short: short, GetService: svc}).NewCmd()
}

// NewCmd creates a new service command.
func (c *ServiceCmd) NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   c.Use,
		Short: c.Short,
		Run: func(cmd *cobra.Command, args []string) {
			v := viper.GetViper()

			service, err := c.GetService(v)
			cobra.CheckErr(err)

			err = service.Start()
			cobra.CheckErr(err)

			err = service.AwaitSignal()
			cobra.CheckErr(err)
		},
	}
}
