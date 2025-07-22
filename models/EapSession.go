/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
}
