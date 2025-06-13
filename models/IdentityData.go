/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IdentityData struct {
	GpsiList           []string          `json:"gpsiList,omitempty"`
	AllowedAfIds       []string          `json:"allowedAfIds,omitempty"`
	ApplicationPortIds map[string]string `json:"applicationPortIds,omitempty"`
	AfIdGpsis          map[string]string `json:"afIdGpsis,omitempty"`
	MtcProviderGpsis   map[string]string `json:"mtcProviderGpsis,omitempty"`
	SupiList           []string          `json:"supiList,omitempty"`
}
