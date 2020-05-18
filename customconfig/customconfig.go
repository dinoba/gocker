package customconfig

import (
	"sync"

	"github.com/spf13/viper"
)

var Config *viper.Viper
var once sync.Once

// start loggeando
func GetInstance() (*viper.Viper, error) {
	var err error
	once.Do(func() {
		Config, err = createConfig()
	})
	return Config, err
}

func createConfig() (*viper.Viper, error) {
	Config := viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath(".")
	err := Config.ReadInConfig()
	return Config, err
}
