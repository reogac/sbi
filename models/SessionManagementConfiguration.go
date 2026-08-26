/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionManagementConfiguration struct {
	DataNetworks      []DataNetworkConfiguration `json:"dataNetworks"`
	Slices            []SliceConfiguration       `json:"slices"`
	TransportNetworks []string                   `json:"transportNetworks"`
}
