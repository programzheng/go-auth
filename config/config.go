package config

type Config interface {
	GetString(name string) string
}

func New() Config {
	config := &Instance{
		Package: NewViper(),
	}

	return config
}
