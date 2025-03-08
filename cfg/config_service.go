package cfg

import (
	"github.com/pjover/coop/domain/ports"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

type configService struct {
	homeDirectory string
}

func NewConfigService() ports.ConfigService {
	service := configService{}
	service.init()
	return service
}

func (c configService) GetString(key string) string {
	return viper.GetString(key)
}

func (c configService) init() {
	c.homeDirectory = c.homeDir()
	c.setupConfig()
	c.loadEnvironmentVariables()
	c.readConfigFile()
}

func (c configService) homeDir() string {
	home := os.Getenv("COOP_HOME")
	if home == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("cannot find home directory: %s", err)
		}
		home = path.Join(homeDir, "coop")
	}

	info, err := os.Stat(home)
	if err != nil {
		log.Fatal(err)
	}
	if !info.IsDir() {
		log.Fatalf("%s is not a valid directory", home)
	}
	return home
}

func (c configService) setupConfig() {
	viper.SetConfigName("coop")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.homeDirectory)
}

func (c configService) loadEnvironmentVariables() {
	viper.AutomaticEnv()
}

func (c configService) readConfigFile() {
	log.Printf("Loading config from `%s` ...", c.homeDirectory)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error reading config file: %s", err)
	}
}

func (c configService) GetWorkingDirectory() string {
	return path.Join(c.GetString("dirs.home"), c.GetString("dirs.current"))
}
