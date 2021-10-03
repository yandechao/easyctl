package install

import (
	// embed
	_ "embed"
	"github.com/spf13/cobra"
	"github.com/weiliang-ms/easyctl/pkg/install"
	"github.com/weiliang-ms/easyctl/pkg/util/command"
)

//go:embed asset/redis_cluster.yaml
var redisClusterConfig []byte

// redisClusterCmd 安装redis集群指令
var redisClusterCmd = &cobra.Command{
	Use:     "redis-cluster [flags]",
	Short:   "配置主机互信",
	Example: "\neasyctl set password-less --server-list=server.yaml",
	Args:    cobra.ExactValidArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := command.SetExecutorDefault(command.ExecutorEntity{
			Cmd:           cmd,
			Fnc:           install.RedisCluster,
			DefaultConfig: redisClusterConfig,
		}, configFile); err != nil {
			panic(err)
		}
	},
}