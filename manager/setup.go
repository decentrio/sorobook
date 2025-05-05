package manager

import (
	"github.com/decentrio/sorobook/aggregation"
	"github.com/decentrio/sorobook/config"
)

func DefaultNewManager(cfg *config.ManagerConfig) *Manager {
	as := aggregation.NewAggregation(cfg.AggregationCfg)
	return NewManager(cfg, as)
}
