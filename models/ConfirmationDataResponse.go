/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfirmationDataResponse struct {
	PvsInfo    []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthResult AuthResult                         `json:"authResult"`
	Supi       string                             `json:"supi,omitempty"`
	Kseaf      string                             `json:"kseaf,omitempty"`
}
