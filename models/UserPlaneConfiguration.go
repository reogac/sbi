/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:59:14 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UserPlaneConfiguration struct {
	TransportNetworks []string             `json:"transportNetworks"`
	DataNetworks      []DataNetworkInfo    `json:"dataNetworks"`
	Slices            []SliceConfiguration `json:"slices"`
}
