/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorData struct {
	SorXmacIue         string         `json:"sorXmacIue,omitempty"`
	SorMacIue          string         `json:"sorMacIue,omitempty"`
	MeSupportOfSorCmci *bool          `json:"meSupportOfSorCmci,omitempty"`
	ProvisioningTime   string         `json:"provisioningTime"`
	UeUpdateStatus     UeUpdateStatus `json:"ueUpdateStatus"`
}
