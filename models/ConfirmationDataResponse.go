/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:41 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfirmationDataResponse struct {
	AuthResult AuthResult                         `json:"authResult"`
	Supi       string                             `json:"supi,omitempty"`
	Kseaf      string                             `json:"kseaf,omitempty"`
	PvsInfo    []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
