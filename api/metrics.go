package api

import (
	"github.com/dukex/mixpanel"
)

type metricsResource struct {
	MX *mixpanel.Mixpanel
}

func NewMetricsResource(userUUID string) *metricsResource {
	m := &metricsResource{}
	client := mixpanel.NewMixpanel("mytoken")
	client.Identify(userUUID)
	m.MX = client
	return m
}
