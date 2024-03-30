package config

import "time"

type Config struct {
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	Secret         string
	AccessExpTime  time.Duration
	RefreshExpTime time.Duration
	DBName         string
	CollectionName string
}
