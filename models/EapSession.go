/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:41 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	KSeaf             string                             `json:"kSeaf,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
