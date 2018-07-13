package config

import (
	"github.com/brocaar/lora-gateway-bridge/internal/legacy/backend/mqttpubsub"
	"github.com/brocaar/lora-gateway-bridge/internal/legacy/gateway"
)

// Config defines the configuration structure.
type Config struct {
	General struct {
		LogLevel int `mapstructure:"log_level"`
	}

	PacketForwarder struct {
		UDPBind      string `mapstructure:"udp_bind"`
		SkipCRCCheck bool   `mapstructure:"skip_crc_check"`

		Configuration []gateway.Configuration `mapstructure:"configuration"`
	} `mapstructure:"packet_forwarder"`

	Backend struct {
		MQTT mqttpubsub.BackendConfig
	}
	Metrics struct {
		Prometheus struct {
			EndpointEnabled bool `mapstructure:"endpoint_enabled"`
			Bind            string
		}
	}
}

// C holds the global configuration.
var C Config
