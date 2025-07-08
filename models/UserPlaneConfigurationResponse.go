/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UserPlaneConfigurationResponse struct {
	TransportNetworks []string             `json:"transportNetworks"`
	DataNetworks      []DataNetworkInfo    `json:"dataNetworks"`
	Slices            []SliceConfiguration `json:"slices"`
}
