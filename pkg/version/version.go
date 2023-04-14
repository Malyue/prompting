package version

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"runtime"
)

var (
	// GitVersion 是语义化的版本号
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是ISO8601 格式的构建时间
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是Git 的SHA1值
	GitCommit = "$Format:%H$"
	// GitTreeState 代表构建时Git仓库的状态，可能的值有：clean, dirty
	GitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String 返回人性化的版本信息字符串
func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return string(s)
	}

	return info.GitVersion
}

// Text 将版本信息编码为 UTF-8 格式文本，并返回
func (info Info) Text() ([]byte, error) {
	// 图形化
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = ""
	table.AddRow("gitVersion:", info.GitVersion)
	table.AddRow("gitCommit:", info.GitCommit)
	table.AddRow("gitTreeState:", info.GitTreeState)
	table.AddRow("buildDate", info.BuildDate)
	table.AddRow("goVerion", info.GoVersion)
	table.AddRow("compiler", info.Compiler)
	table.AddRow("platform", info.Platform)

	return table.Bytes(), nil
}

func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

// Get 返回详细的代码库版本信息，用来标明二进制文件由哪个版本的代码构建
func Get() Info {
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}