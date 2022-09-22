package cmd

import "go.uber.org/zap"

func newLogger(debug bool) *zap.SugaredLogger {
	var cfg zap.Config
	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
		cfg.Encoding = "console"
	}

	return zap.Must(cfg.Build()).Sugar()
}
