package initLoader

import "github.com/glieske/kube-nanny/internal/dto"

type initHandler interface {
	execute(*dto.MainConfig)
	setNext(next initHandler)
}

func ExecuteInit(config *dto.MainConfig) {
	initChain := &LoadLoggerInitHandler{}
	initChain.execute(config)
}
