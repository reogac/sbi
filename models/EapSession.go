/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
}
