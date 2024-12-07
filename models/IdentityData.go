/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IdentityData struct {
	AfIdGpsis          map[string]string `json:"afIdGpsis,omitempty"`
	MtcProviderGpsis   map[string]string `json:"mtcProviderGpsis,omitempty"`
	SupiList           []string          `json:"supiList,omitempty"`
	GpsiList           []string          `json:"gpsiList,omitempty"`
	AllowedAfIds       []string          `json:"allowedAfIds,omitempty"`
	ApplicationPortIds map[string]string `json:"applicationPortIds,omitempty"`
}
