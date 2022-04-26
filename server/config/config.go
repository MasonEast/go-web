package config

type Server struct {

	// gorm
	Mysql  Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

}
