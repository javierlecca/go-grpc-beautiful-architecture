package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"grcp"`
	Port    int32  `default:"8000"`
	DB      struct {
		Use   string `default:"mysql"`
		Mysql []struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"3306"`
			UserName string `default:"root"`
			Password string `default:"xxxxxxxx"`
			Database string `default:"xxxxxxxx"`
		}
	}
	Contacts struct {
		Name  string `default:"javier Lecca"`
		Email string `default:"leccajavier@gmail.com"`
	}
}

func (c *Config) NewConfig() (*Config, error) {
	err := configor.Load(c, "config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
