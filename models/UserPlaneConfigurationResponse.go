/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 17:37:55 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UserPlaneConfigurationResponse struct {
	Slices            []SliceConfiguration `json:"slices"`
	TransportNetworks []string             `json:"transportNetworks"`
	DataNetworks      []DataNetworkInfo    `json:"dataNetworks"`
}
