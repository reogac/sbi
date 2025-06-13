/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:09 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionManagementConfiguration struct {
	Slices            []SliceConfiguration       `json:"slices"`
	TransportNetworks []string                   `json:"transportNetworks"`
	DataNetworks      []DataNetworkConfiguration `json:"dataNetworks"`
}
