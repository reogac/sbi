/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	KSeaf             string                             `json:"kSeaf,omitempty"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
