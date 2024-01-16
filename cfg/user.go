package cfg

// Modify below to reflect the contents of the user.yaml file
// Please note that all values can be changed by the user by modifying
// the configuration file $HOME/.config/<appname>/<appname>.yaml
// that is initially copied with the default values that you defined in the file yaml/user.yaml
//
// Remember to use uppercase letters as the names of the fileds, otherwise Viper will not
// be able to unmarshal the data to the User structure.
// Also, if the name of a field in the user.yaml file is different from the one you define here,
// remember to use `mapstructure:"name-in-yaml"` for each one of those fields.
type UserYaml struct {
	Greetings string
}
