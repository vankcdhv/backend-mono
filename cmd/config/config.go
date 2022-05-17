package config

type Config struct {
	Http     HttpServer `json:"http" yaml:"http"`
	Database Database   `json:"database" yaml:"database"`
}

type HttpServer struct {
	Path string `json:"path" yaml:"path"`
	Port string `json:"port" yaml:"port"`
}

type Database struct {
	Driver string `json:"driver" yaml:"driver"`
	Source string `json:"source" yaml:"source"`
}
