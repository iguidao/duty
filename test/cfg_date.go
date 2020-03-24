package main

import (
	"fmt"
	"log"
	"oncall/src/cfg"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

type Datetime struct {
	start string
	end   string
}

func main() {
	if err := cfg.Init(""); err != nil {
		panic(err)
	}
	nextday := "2020-05-05"
	holiday := Get_holiday("holiday")
	for _, aaa := range holiday {
		if nextday == aaa {
			log.Println("cunzai")
		}
	}
	log.Println("bucunzai")
	// laborday := holiday["laborday"]
	// nationalday := holiday["nationalday"]
	// springfestivalday := holiday["springfestivalday"]
	//springfestivalday_end := springfestivalday.(Datetime).end
	//log.Println(springfestivalday_end)
	// log.Println(laborday, nationalday, springfestivalday)
}

func Get_holiday(get_local string) []string {
	local_holiday := viper.GetStringSlice("local.holiday")
	return local_holiday

}

func Get_Local(get_local string) string {
	switch get_local {
	case "addr":
		local_addr := viper.GetString("local.addr")
		return local_addr
	case "secretkey":
		local_secretkey := viper.GetString("local.secretkey")
		return local_secretkey
	default:
		return "noconfig"
	}
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("yaml")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}
