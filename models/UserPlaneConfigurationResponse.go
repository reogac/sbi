/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UserPlaneConfigurationResponse struct {
	DataNetworks      []DataNetworkInfo    `json:"dataNetworks"`
	Slices            []SliceConfiguration `json:"slices"`
	TransportNetworks []string             `json:"transportNetworks"`
}
