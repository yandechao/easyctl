package set

import (
	"github.com/spf13/cobra"
	"github.com/weiliang-ms/easyctl/pkg/set"
)

// 文件描述符
var ulimitCmd = &cobra.Command{
	Use:   "ulimit [flags]",
	Short: "配置ulimit",
	Args:  cobra.ExactValidArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if runErr := Set(Entity{Cmd: cmd, Fnc: set.Ulimit}); runErr != nil {
			panic(runErr)
		}
	},
}