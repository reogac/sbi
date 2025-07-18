/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorData struct {
	MeSupportOfSorCmci *bool          `json:"meSupportOfSorCmci,omitempty"`
	ProvisioningTime   string         `json:"provisioningTime"`
	UeUpdateStatus     UeUpdateStatus `json:"ueUpdateStatus"`
	SorXmacIue         string         `json:"sorXmacIue,omitempty"`
	SorMacIue          string         `json:"sorMacIue,omitempty"`
}
