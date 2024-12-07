/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedirectInformation struct {
	RedirectEnabled       *bool               `json:"redirectEnabled,omitempty"`
	RedirectAddressType   RedirectAddressType `json:"redirectAddressType,omitempty"`
	RedirectServerAddress string              `json:"redirectServerAddress,omitempty"`
}
