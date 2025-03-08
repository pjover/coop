package ports

type ConfigService interface {
	GetString(key string) string
	GetWorkingDirectory() string
}
