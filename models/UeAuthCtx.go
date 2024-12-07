/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:17 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Supi       string                             `json:"supi"`
	Kamf       []byte                             `json:"kamf,omitempty"`
	Rand       []byte                             `json:"rand,omitempty"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
