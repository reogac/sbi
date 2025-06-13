/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedirectInformation struct {
	RedirectServerAddress string              `json:"redirectServerAddress,omitempty"`
	RedirectEnabled       *bool               `json:"redirectEnabled,omitempty"`
	RedirectAddressType   RedirectAddressType `json:"redirectAddressType,omitempty"`
}
