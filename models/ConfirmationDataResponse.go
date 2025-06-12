/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfirmationDataResponse struct {
	Kseaf      string                             `json:"kseaf,omitempty"`
	PvsInfo    []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthResult AuthResult                         `json:"authResult"`
	Supi       string                             `json:"supi,omitempty"`
}
