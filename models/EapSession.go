/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:29 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
}
