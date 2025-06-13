package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Auth   AuthConfig   `yaml:"auth"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}

type AuthConfig struct {
	Secret     string `yaml:"secret"`
	AccessExp  int    `yaml:"access_exp"`
	RefreshExp int    `yaml:"refresh_exp"`
}
