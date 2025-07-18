/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	EapPayload        string                             `json:"eapPayload"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	Links             map[string]Link                    `json:"_links,omitempty"`
}
