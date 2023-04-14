package verflag

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"prompting/pkg/version"
	"strconv"
)

// 作为pflag.Value
type versionValue int

// 定义一个枚举类型，表示程序版本号是否需要展示，以及展示方式
const (
	VersionFalse versionValue = 0
	VersionTrue  versionValue = 1
	VersionRaw   versionValue = 2
)

const (
	strRawVersion   = "raw"
	versionFlagName = "version"
)

var versionFlag = Version(versionFlagName, VersionFalse, "Print version information and quit.")

// IsBoolFlag 表示这个类型的值可以被解析为一个布尔值
func (v *versionValue) IsBoolFlag() bool {
	return true
}

// Get 返回该类型的值的指针
func (v *versionValue) Get() interface{} {
	return v
}

// String 转换成对应字符串
func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}

	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

// Set
func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

func (v *versionValue) Type() string {
	return "version"
}

// 定义了一个具有指定名称和用法的标志
// VersionVar 可以在任意FkagSet上注册这个包的标志
func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)
	pflag.Lookup(name).NoOptDefVal = "true"
}

// Version 封装VersionVar函数
func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)

	return p
}

// AddFlages 在任意FlagSet上注册这个包的标志，这样它们指向与全局标志相同的值.
func AddFlages(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested 将检查是否传递了 `--version` 标志，如果是，则打印版本并退出.
func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
