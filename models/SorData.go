/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorData struct {
	UeUpdateStatus     UeUpdateStatus `json:"ueUpdateStatus"`
	SorXmacIue         string         `json:"sorXmacIue,omitempty"`
	SorMacIue          string         `json:"sorMacIue,omitempty"`
	MeSupportOfSorCmci *bool          `json:"meSupportOfSorCmci,omitempty"`
	ProvisioningTime   string         `json:"provisioningTime"`
}
