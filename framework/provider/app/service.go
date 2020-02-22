package app

type HadeApp struct{
}

func NewHadeApp(params ...interface{}) (interface{}, error) {
	return &HadeApp{}, nil
}

// application version
func (app *HadeApp) Version() string {
	return "0.0.1"
}

// base path which is the base folder
func (app *HadeApp) BasePath() string {
	return "/Users/didi/Documents/workspace/hade/" 
}

// config folder which contains config
func (app *HadeApp) ConfigPath() string {
	return app.BasePath() + "config/"
}

// environmentPath which contain .env
func (app *HadeApp) EnvironmentPath() string {
	return app.BasePath()
}

// storagePath define storage folder
func (app *HadeApp) StoragePath() string {
	return app.BasePath() + "storage/"
}

// logPath define logPath
func (app *HadeApp) LogPath() string {
	return app.StoragePath() + "logs/"
}

