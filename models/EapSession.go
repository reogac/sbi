/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:43 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	Supi              string                             `json:"supi,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
}
