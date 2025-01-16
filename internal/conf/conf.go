package conf

import log "unknwon.dev/clog/v2"

func init() {
	// Initialize the primary logger until logging service is up.
	err := log.NewConsole()
	if err != nil {
		panic("init console logger: " + err.Error())
	}
}
