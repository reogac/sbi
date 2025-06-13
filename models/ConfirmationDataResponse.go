/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfirmationDataResponse struct {
	Supi       string                             `json:"supi,omitempty"`
	Kseaf      string                             `json:"kseaf,omitempty"`
	PvsInfo    []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthResult AuthResult                         `json:"authResult"`
}
