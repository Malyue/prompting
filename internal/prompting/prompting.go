package prompting

import (
	"github.com/spf13/cobra"
	"prompting/internal/pkg/log"
	"prompting/pkg/version/verflag"
)

var cfgFile string

func NewPromptingCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该命令会出现在帮助信息中
		Use: "prompting",
		// 命令的简短描述
		Short: "A project Of Prompting-Community",
		// 命令的详细描述
		Long: "A project Of Prompting-Community",

		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()

			// 初始化日志
			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中

			return run()
		},
	}
}

// 实际的业务代码入口函数
func run() error {

}


