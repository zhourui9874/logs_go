package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strconv"
)

type config_info struct {
	debug_dir     string
	info_dir      string
	dangerous_dir string
	errors_dir    string
	crash_dir     string
	level_nume    int
}

var config *config_info

var base_dir string

func init() {
	config = new(config_info)
	cfg, err := ini.InsensitiveLoad("./conf/conf.ini")
	if err != nil {
		fmt.Println("read log conf file error : ", err.Error())
		return
	}
	sfg := cfg.Section("log_conf")
	config.debug_dir = sfg.Key("debug").MustString("")
	config.info_dir = sfg.Key("info").MustString("")
	config.dangerous_dir = sfg.Key("dangerous").MustString("")
	config.errors_dir = sfg.Key("error").MustString("")
	config.crash_dir = sfg.Key("crash").MustString("")
	config.level_nume = sfg.Key("level").MustInt(0)
	base_dir = sfg.Key("base_dir").MustString("")
}

func SetBaseDir(dir string) bool {
	base_dir = dir
	return SetConfFile("base_dir", dir)
}
func SetLogPrintLevel(level int) bool {
	config.level_nume = level
	return SetConfFile("base_dir", strconv.Itoa(level))
}

func GetBaseDir() string      { return base_dir }
func GetDebugDir() string     { return config.debug_dir }
func GetInfoDir() string      { return config.info_dir }
func GetDangerousDir() string { return config.dangerous_dir }
func GetErrorDir() string     { return config.errors_dir }
func GetCrashDir() string     { return config.crash_dir }
func GetIntoFileLevel() int   { return config.level_nume }

func SetConfFile(key, value string) bool {
	config = new(config_info)
	cfg, err := ini.InsensitiveLoad("./conf/conf.ini")
	if err != nil {
		fmt.Println("read log conf file error : ", err.Error())
		return false
	}
	sfg := cfg.Section("log_conf")
	sfg.Key(key).SetValue(value)
	return true
}
