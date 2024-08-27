package viper

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Use() {
	// init source with Time UTC +0
	os.Setenv("TZ", "GMT+0")
	viper.WatchConfig()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
