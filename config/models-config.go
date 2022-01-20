package config

//Config ...
type Config struct {
	APIPort int `yaml:"api port"`
	Linx    SQL `yaml:"linx"`
	App     SQL `yaml:"fin"`
}

//SQL ...
type SQL struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth `yaml:",inline"`

	Db string `yaml:"db"`
}

//Auth ...
type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
