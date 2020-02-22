package contract

// AppKey is the key in container
const AppKey = "app"

// App define application structure
type App interface {
	// application version
	Version() string
	// base path which is the base folder
	BasePath() string
	// config folder which contains config
	ConfigPath() string
	// environmentPath which contain .env
	EnvironmentPath() string
	// storagePath define storage folder
	StoragePath() string
	// logPath define logPath
	LogPath() string
}
