/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionManagementConfiguration struct {
	TransportNetworks []string                   `json:"transportNetworks"`
	DataNetworks      []DataNetworkConfiguration `json:"dataNetworks"`
	Slices            []SliceConfiguration       `json:"slices"`
}
