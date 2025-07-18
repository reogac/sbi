/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfirmationDataResponse struct {
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthResult AuthResult                         `json:"authResult"`
	Supi       string                             `json:"supi,omitempty"`
	Kseaf      string                             `json:"kseaf,omitempty"`
	PvsInfo    []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
}
