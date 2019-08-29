package config

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/viper"

	log "github.com/hobord/infra/log"
)

var State ConfigState

func init() {
	//load config
}

func loadConfigFiles(configDir string) {
	if configDir == "" {
		configDir = os.Getenv("CONFIGDIR")
		if configDir == "" {
			configDir = "config"
		}
	}

	var files []string
	err := filepath.Walk(configDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)

		v := viper.New()
		v.SetConfigFile(file)
		if err := v.ReadInConfig(); err != nil {
			log.Logger.Fatalf("Error reading config file, %s", err)
		}

		APIVersion := v.GetString("apiVersion")
		cfgKind := v.GetString("kind")

		if APIVersion != "InfraService/v1" {
			continue
		}
		switch cfgKind {
		case "CtxRouterConfig":
			State.loadConfig(v)
		}
	}
}
func (configState *ConfigState) loadConfig(v *viper.Viper) {
	if configState == nil {
		configState = &ConfigState{}
	}

	cfg := &RouterConfigYaml{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		// log.Fatalf("unable to decode into struct, %v", err)
		// TODO: error handling
		return
	}

	if configState.Hosts == nil {
		configState.Hosts = make(map[string]rutesByProtcols)
	}
	for _, host := range cfg.Spec.Hosts {
		var protocols []string
		if len(cfg.Spec.Protocols) > 0 {
			protocols = cfg.Spec.Protocols
		} else {
			protocols = []string{"http", "https"}
		}

		for _, protocol := range protocols {
			if configState.Hosts[host] == nil {
				configState.Hosts[host] = make(map[string][]routeRule)
			}
			for _, route := range cfg.Spec.Routes {
				reg, err := regexp.Compile(route.Path)
				if err != nil {
					continue // TODO: errorlog
				}

				rheaders := make([]rules)
				for _, header := range route.RequestHeader.Require {

				}
				newRuteRule := routeRule{
					Path:          reg,
					RequestHeader: rules{},
					CookieValues:  rules{},
					FormValues:    rules{},
					SessionValues: rules{},
					CustomLogic:   route.CustomLogic,
					Target:        routeTarget{},
				}
				configState.Hosts[host][protocol] = append(configState.Hosts[host][protocol], newRuteRule)
			}
		}
	}
}

func 