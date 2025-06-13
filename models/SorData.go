/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorData struct {
	ProvisioningTime   string         `json:"provisioningTime"`
	UeUpdateStatus     UeUpdateStatus `json:"ueUpdateStatus"`
	SorXmacIue         string         `json:"sorXmacIue,omitempty"`
	SorMacIue          string         `json:"sorMacIue,omitempty"`
	MeSupportOfSorCmci *bool          `json:"meSupportOfSorCmci,omitempty"`
}
