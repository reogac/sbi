/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:57 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	Links             map[string]Link                    `json:"_links,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
}
