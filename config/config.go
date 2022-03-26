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

func GetProductionStatus() bool {
	env := New()
	if env.GetString("ENV") == "production" {
		return true
	}

	return false
}
