package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"taogin/config/load"
)

var (
	TAO_SERVER load.Server
	TAO_VIPER  *viper.Viper
	TAO_LOG    *zap.Logger
)
