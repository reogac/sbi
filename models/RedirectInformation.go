/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedirectInformation struct {
	RedirectServerAddress string              `json:"redirectServerAddress,omitempty"`
	RedirectEnabled       *bool               `json:"redirectEnabled,omitempty"`
	RedirectAddressType   RedirectAddressType `json:"redirectAddressType,omitempty"`
}
