/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IdentityData struct {
	SupiList           []string          `json:"supiList,omitempty"`
	GpsiList           []string          `json:"gpsiList,omitempty"`
	AllowedAfIds       []string          `json:"allowedAfIds,omitempty"`
	ApplicationPortIds map[string]string `json:"applicationPortIds,omitempty"`
	AfIdGpsis          map[string]string `json:"afIdGpsis,omitempty"`
	MtcProviderGpsis   map[string]string `json:"mtcProviderGpsis,omitempty"`
}
