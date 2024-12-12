package n42

import (
	"github.com/reogac/sbi/models"
	"time"
)

type HeartbeatRequest struct {
	Nonce int64 //for matching request/response
	Time  time.Time
}

type HeartbeatResponse struct {
	Nonce int64 //must be equal the value in request
	Time  time.Time
}
type RegistrationRequest struct {
	SbiIp   string          `json:"sbiip"`
	SbiPort int             `json:"sbiport"`
	Slices  []models.Snssai `json:"slices"`
	Infs    []NetInfConfig  `json:"infs"`
	Dnns    []DnnInfoConfig `json:"dnns"`
	Time    int64           `json:"time"`
	Active  bool            `json:"active"`
}

type RegistrationResponse struct {
	UpfId string
	Time  int64
}

type NetInfConfig struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Mtu  uint32 `json:"mtu"`
}

type DnnInfoConfig struct {
	Dnn  string `json:"dnn"`
	Cidr string `json:"cidr"`
	Addr string `json:"addr"`
}
