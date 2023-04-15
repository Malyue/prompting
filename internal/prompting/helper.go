package prompting

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"prompting/internal/pkg/log"
	"prompting/internal/prompting/store"
	_mysql "prompting/pkg/db/Mysql"
	"strings"
)

const (
	// recommendedHomeDir 定义放置服务配置的默认目录
	recommendedHomeDir = ".prompting"

	// defaultConfigName 指定服务的默认配置文件名
	defaultConfigName = "prompting.yaml"
)

func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		// On Unix,including Linux and macOS, it returns $HOME environment variable
		// On Windows, it returns %USERPROFILE% environment variable
		// On plan 9,it returns the $home environment variable
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)

		// 这里直接先写死home
		home = "../../../configs/"

		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)

	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为 Prompting，如果是 Prompting，将自动转变为大写。
	viper.SetEnvPrefix("Prompting")
	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read config file", "error", err)
	}

	// 打印 viper 中的配置信息
	log.Debugw("Using config file", "file", viper.ConfigFileUsed())
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 返回
// 注：`viper.Get<Type>()` 中 key 的名字需要使用 '.' 分割，以跟 YAML 中保持相同的缩进
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

// initStore 读取db配置，创建gorm.DB 实例，并初始化 store 层
func initStore() error {
	dbOptions := &_mysql.MySQLOptions{
		Host:                  viper.GetString("db.host"),
		Port:                  viper.GetInt("db.port"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	ins, err := _mysql.NewMysql(dbOptions)
	if err != nil {
		return err
	}
	_ = store.NewStore(ins)

	return nil
}
