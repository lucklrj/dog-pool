package config

import (
	"flag"
	"lib/system"

	"errors"
	"github.com/go-ini/ini"
)

type System struct {
	Port int
}
type Database struct {
	Host            string
	Port            int
	User            string
	PassWord        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	AllowDatabase   []string
}
type Memcached struct {
	Host            string
	Port            int
	MaxOpenConns    int
	ConnMaxLifeTime int
	ConnTimeOut     int
}
type Redis struct {
	Host            string
	Port            int
	MaxOpenConns    int
	ConnMaxLifeTime int
}

var SystemConfig = new(System)
var DatabaseConfig = new(Database)
var MemcachedConfig = new(Memcached)
var RedisConfig = new(Redis)

func init() {
	InputConfigFilePath := flag.String("f", "./config.ini", "系统配置文件")
	flag.Parse()

	cfg, err := ini.Load(*InputConfigFilePath)
	if err != nil {
		system.Exit(errors.New("无法加载配置文件"))
	}
	err = cfg.Section("system").MapTo(SystemConfig)
	if err != nil {
		system.Exit(errors.New("无法加载[system]配置"))
	}
	err = cfg.Section("database").MapTo(DatabaseConfig)
	if err != nil {
		system.Exit(errors.New("无法加载[database]配置"))
	}
	err = cfg.Section("memcached").MapTo(MemcachedConfig)
	if err != nil {
		system.Exit(errors.New("无法加载[memcached]配置"))
	}
	err = cfg.Section("redis").MapTo(RedisConfig)
	if err != nil {
		system.Exit(errors.New("无法加载[redis]配置"))
	}
}
