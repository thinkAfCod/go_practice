package conf

import (
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"simple-explore/common"
)

var Config MyConfiguration

type MyConfiguration struct {
	iris.Configuration
	Server     Server     `json:"server,omitempty" yaml:"server" toml:"server"`
	DataSource DataSource `json:"dataSource,omitempty" yaml:"dataSource" toml:"dataSource"`
}

type DataSource struct {
	DriverName string `json:"driverName,omitempty" yaml:"driverName" toml:"driverName"`
	Username   string `json:"username,omitempty" yaml:"username" toml:"username"`
	Password   string `json:"password,omitempty" yaml:"password" toml:"password"`
	Domain     string `json:"domain,omitempty" yaml:"domain" toml:"domain"`
	Path       string `json:"path,omitempty" yaml:"path" toml:"path"`
}

type Server struct {
	Host string `json:"host,omitempty" yaml:"host" toml:"host"`
	Port string `json:"port,omitempty" yaml:"port" toml:"port"`
	Name string `json:"name,omitempty" yaml:"name" toml:"name"`
}

func DefaultConfiguration() MyConfiguration {
	return MyConfiguration{
		Server: Server{
			Host: "",
			Port: "8080",
			Name: "",
		},
		DataSource: DataSource{
			DriverName: "",
			Username:   "",
			Password:   "",
			Domain:     "",
			Path:       "",
		},
		Configuration: iris.DefaultConfiguration(),
	}
}

func YAML(filename string) {
	Config = DefaultConfiguration()
	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	yamlAbsPath, err := filepath.Abs(filename)
	common.PanicErr(err)

	// read the raw contents of the file
	data, err := ioutil.ReadFile(yamlAbsPath)
	common.PanicErr(err)

	// put the file's contents as yaml to the default configuration(c)
	err = yaml.Unmarshal(data, &Config)
	common.PanicErr(err)
}
