package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/blackhorseya/ekko/pkg/contextx"
	ib "github.com/blackhorseya/ekko/pkg/entity/domain/issue/biz"
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewCLI)

var (
	path string
)

type impl struct {
	logger *zap.Logger
	v      *viper.Viper

	cmd *cobra.Command
	biz ib.IBiz
}

// NewCLI serve caller to create adapter cli
func NewCLI(v *viper.Viper, biz ib.IBiz) adapters.CLI {
	i := &impl{
		v:   v,
		biz: biz,
		cmd: &cobra.Command{
			Use:   "ekko",
			Short: "A command line tool for ekko",
			Long: `
███████╗██╗  ██╗██╗  ██╗ ██████╗ 
██╔════╝██║ ██╔╝██║ ██╔╝██╔═══██╗
█████╗  █████╔╝ █████╔╝ ██║   ██║
██╔══╝  ██╔═██╗ ██╔═██╗ ██║   ██║
███████╗██║  ██╗██║  ██╗╚██████╔╝
╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ 
`,
			Run: func(cmd *cobra.Command, args []string) {
				ret, err := biz.GetByID(contextx.Background(), 1642767828740214784)
				if err != nil {
					log.Println(err)
					return
				}

				out, err := json.Marshal(ret)
				if err != nil {
					log.Println(err)
					return
				}

				fmt.Println(string(out))
			},
		},
	}

	cobra.OnInitialize(i.initConfig)

	i.cmd.PersistentFlags().StringVarP(&path, "config", "c", "", "config file (default is $HOME/.ekko.yaml)")

	return i
}

func (i *impl) initConfig() {
	if path != "" {
		i.v.SetConfigFile(path)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		i.v.AddConfigPath(home)
		i.v.SetConfigType("yaml")
		i.v.SetConfigName(".ekko")
	}

	i.v.AutomaticEnv()

	err := i.v.ReadInConfig()
	cobra.CheckErr(err)

	log.Println("Using config file:", i.v.ConfigFileUsed())
}

func (i *impl) Execute() error {
	return i.cmd.Execute()
}
