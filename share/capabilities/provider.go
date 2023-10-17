package capabilities

import (
	"github.com/renatomb/open-rport/server/chconfig"
	chshare "github.com/renatomb/open-rport/share"
	"github.com/renatomb/open-rport/share/models"
)

func NewServerCapabilities(cfg *chconfig.MonitoringConfig) *models.Capabilities {
	caps := models.Capabilities{
		ServerVersion:      chshare.BuildVersion,
		MonitoringVersion:  chshare.MonitoringVersion,
		IPAddressesVersion: chshare.IPAddressesVersion,
	}

	if !cfg.Enabled {
		caps.MonitoringVersion = 0
	}
	return &caps
}
