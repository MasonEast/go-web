package config

type Config struct {

	Mysql  Mysql `json:"mysql" yaml:"mysql"`
	JWT		JWT  `json:"jwt" yaml:"jwt"`

}
