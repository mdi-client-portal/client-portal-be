package config

import (
    "go.uber.org/zap"
)

var Log *zap.Logger

func LoggerInit() {
    Log, _ = zap.NewProduction()
}
