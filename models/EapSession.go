/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	Msk               string                             `json:"msk,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
}
