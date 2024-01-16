package cfg

import (
	"bytes"
	"embed"
	_ "embed"
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

//
// Please DO NOT modify this file unless you know what you are doing!
//
// Modify application.go / application.yaml to change your design-time constants
// Modify user.go / user.yaml to give the user the possibility to change some parameters
//

var (
	// Global variables available to all packages
	App  ApplicationYaml
	User UserYaml
)

//go:embed yaml
var yamlDir embed.FS

func init() {
	applicationCfgInit()
	if App.UserYaml {
		userCfgInit()
	}
}

//
// ----------------------[ APPLICATION.YAML CONFIGURATION ]-----------------------------
//

// This function is internal to the package and it's called by cfg on its own init()
// It ensures that application.yaml is read before user.yaml
func applicationCfgInit() {
	// The application.yaml file is MANDATORY. In case of errors we fail!
	internalYaml, err := yamlDir.ReadFile("yaml/application.yaml")
	if err != nil {
		log.Fatalln("couldn't load internal.yaml:", err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewReader(internalYaml))
	if err != nil {
		log.Fatalln("couldn't read internal.yaml:", err)
	}
	viper.Unmarshal(&App)
}

//
// --------------------------[ USER.YAML CONFIGURATION ]--------------------------------
//

func userCfgInit() {
	// Construct the user pathnames in his/her home folder
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("could not get the HOME path:", err)
	}
	appName := filepath.Base(App.ModuleName)
	userConfigFileName := appName + ".yaml"
	userConfigDir := path.Join(userHomeDir, ".config", appName)
	userConfigFile := path.Join(userConfigDir, userConfigFileName)

	// If the file exists, use it!
	_, err = os.Stat(userConfigFile)
	if !errors.Is(err, os.ErrNotExist) {
		loadUserConfiguration(userConfigFile)
		return
	}

	// If it doesn't, create it from the embedded file
	createUserConfiguration(userConfigDir, userConfigFile)
}

// Common function to unmarshal the yaml file into the User structure
func unmarshalUserConfiguration() {
	err := viper.Unmarshal(&User)
	if err != nil {
		log.Fatalln("couldn't unmarshal user.yaml:", err)
	}
}

func loadUserConfiguration(configFile string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("error loading config.yaml:", err)
	}
	unmarshalUserConfiguration()
}

func createUserConfiguration(userConfigDir, userConfigFile string) {

	// If the user folder does not exist, create it
	// Eventual errors depend on the user's PC, so we log it but continue
	err := os.MkdirAll(userConfigDir, 0755)
	if err != nil {
		log.Println("could not create folder:", err)
		useDefaultUserConfig()
		return
	}

	// Read the embedded default user.yaml configuration file
	userYaml, err := yamlDir.ReadFile("yaml/user.yaml")
	if err != nil {
		log.Fatalln("could not read user.yaml:", err)
	}

	// Copy the default configuration file to the destination
	err = os.WriteFile(userConfigFile, userYaml, 0644)
	if err != nil {
		log.Println("could not create user file:", err)
		useDefaultUserConfig()
		return
	}

	// Use it!
	loadUserConfiguration(userConfigFile)
}

func useDefaultUserConfig() {
	viper.SetConfigType("yaml")
	userYaml, err := yamlDir.ReadFile("yaml/user.yaml")
	if err != nil {
		log.Println("could not read :", err)
		return
	}
	err = viper.ReadConfig(bytes.NewReader(userYaml))
	if err != nil {
		log.Fatalln("couldn't read internal.yaml:", err)
	}

	unmarshalUserConfiguration()
	log.Println("using default user configuration")
}
