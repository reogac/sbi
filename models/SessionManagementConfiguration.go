/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionManagementConfiguration struct {
	Slices            []SliceConfiguration       `json:"slices"`
	TransportNetworks []string                   `json:"transportNetworks"`
	DataNetworks      []DataNetworkConfiguration `json:"dataNetworks"`
}
