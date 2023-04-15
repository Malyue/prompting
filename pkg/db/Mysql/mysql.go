package Mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// MySQLOptions 定义 Mysql 数据库选项
type MySQLOptions struct {
	Host                  string        `json:"host" yaml:"host"`
	Port                  int           `json:"port" yaml:"port"`
	Username              string        `json:"username" yaml:"username"`
	Password              string        `json:"password" yaml:"password"`
	Database              string        `json:"database" yaml:"database"`
	MaxIdleConnections    int           `json:"max_idle_connections" yaml:"max_idle_connections"`
	MaxOpenConnections    int           `json:"max_open_connections" yaml:"max_open_connections"`
	MaxConnectionLifeTime time.Duration `json:"max_connection_life_time" yaml:"max_connection_life_time"`
	LogLevel              int           `json:"log_level" yaml:"log_level"`
}

// DSN 从 MysqlOptions 返回 DSN
func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Host,
		o.Port,
		o.Database,
		true,
		"Local")
}

// NewMySQL 使用给定的选项创建一个新的 gorm 数据库实例
func NewMysql(opts *MySQLOptions) (*gorm.DB, error) {
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}
	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns 设置到数据库的最大打开连接数
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime 设置连接可重用的最长时间
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns 设置空闲连接池的最大连接数
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
