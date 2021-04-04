package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var DB = struct {
	Name     string `default:"sample_db"`
	Host     string `default:"db"`
	User     string `required:"true" env:"MYSQL_USER"`
	Password string `required:"true" env:"MYSQL_PASS"`
	Port     string `default:"3306"`
}{}

func Setup() {
	configor.Load(&DB, []string{}...)
	fmt.Printf("config: %#v", DB)
}
