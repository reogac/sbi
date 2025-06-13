/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MmeCapabilities struct {
	EthernetSupported *bool `json:"ethernetSupported,omitempty"`
	UpipSupported     *bool `json:"upipSupported,omitempty"`
	NonIpSupported    *bool `json:"nonIpSupported,omitempty"`
}
