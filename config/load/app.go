package load

type App struct {
	AppName string `mapstructure:"app-name" json:"app-name" yaml:"app-name"`
	Debug   string `mapstructure:"debug" json:"debug" yaml:"debug"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
}
