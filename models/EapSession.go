/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	Supi              string                             `json:"supi,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	Links             map[string]Link                    `json:"_links,omitempty"`
}
