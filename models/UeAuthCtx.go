/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	NgKsi      NgKsi                              `json:"ngKsi"`
	Rand       []byte                             `json:"rand,omitempty"`
	Eap        []byte                             `json:"eap,omitempty"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
	EapSuccess bool                               `json:"eapSuccess"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Supi       string                             `json:"supi"`
	Kamf       []byte                             `json:"kamf,omitempty"`
}
