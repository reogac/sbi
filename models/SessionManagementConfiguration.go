/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Mar 25 10:55:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionManagementConfiguration struct {
	DataNetworks      []DataNetworkConfiguration `json:"dataNetworks"`
	TransportNetworks []string                   `json:"transportNetworks"`
}
