package dto

type RunningConfig struct {
	WatcherConfig watcherConfig
}

type watcherConfig struct {
	LoopDelay int64
}

func NewRunningConfig() *RunningConfig {
	return &RunningConfig{
		WatcherConfig: watcherConfig{
			LoopDelay: 30,
		},
	}
}
