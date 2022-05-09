package config

import "google.golang.org/protobuf/types/known/durationpb"

type Config struct {
	Http     HttpServer `json:"http" yaml:"http"`
	Database Database   `json:"database" yaml:"database"`
	Redis    Redis      `json:"redis" yaml:"redis"`
}

type HttpServer struct {
	Path string `json:"path" yaml:"path"`
	Port string `json:"port" yaml:"port"`
}

type Database struct {
	Driver string `json:"driver" yaml:"driver"`
	Source string `json:"source" yaml:"source"`
}

type Redis struct {
	Network      string               `json:"network" yaml:"network"`
	Address      string               `json:"address" yaml:"address"`
	ReadTimeout  *durationpb.Duration `json:"read_timeout" yaml:"readTimeout"`
	WriteTimeout *durationpb.Duration `json:"write_timeout" yaml:"writeTimeout"`
}
