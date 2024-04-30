package preInitLoader

import log "github.com/sirupsen/logrus"

func ExecutePreInit() {
	onPreInitLoadLoggingConfiguration()
}

func onPreInitLoadLoggingConfiguration() {
	log.SetFormatter(&log.JSONFormatter{})
}
