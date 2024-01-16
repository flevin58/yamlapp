package cfg

// Modify below to reflect the contents of the application.yaml file
// Please note that all values are internal to the application, the user will
// not be able to change them.
// This file is embedded at compile time and so every time the app starts, the values will
// be those in the yaml/application.yaml file.
// Use this file to store constants that can be accessed as a sort of 'global context' by cfg.App
//
// Remember to use uppercase letters as the names of the fileds, otherwise Viper will not
// be able to unmarshal the data to the User structure.
// Also, if the name of a field in the user.yaml file is different from the one you define here,
// remember to use `mapstructure:"name-in-yaml"` for each one of those fields.
type ApplicationYaml struct {
	// Please don't modify the four fields below
	ModuleName string
	Name       string
	Version    string
	UserYaml   bool `mapstructure:"userconfig"`
	// Add here below as needed
}
