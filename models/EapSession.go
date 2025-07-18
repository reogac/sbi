/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EapSession struct {
	AuthResult        AuthResult                         `json:"authResult,omitempty"`
	AmData            *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	KSeaf             string                             `json:"kSeaf,omitempty"`
	Links             map[string]Link                    `json:"_links,omitempty"`
	PvsInfo           []ServerAddressingInfo             `json:"pvsInfo,omitempty"`
	Msk               string                             `json:"msk,omitempty"`
	EapPayload        string                             `json:"eapPayload"`
	Supi              string                             `json:"supi,omitempty"`
	SupportedFeatures string                             `json:"supportedFeatures,omitempty"`
}
