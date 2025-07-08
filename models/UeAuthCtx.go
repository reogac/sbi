/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Supi       string                             `json:"supi"`
	Kamf       []byte                             `json:"kamf,omitempty"`
	Rand       []byte                             `json:"rand,omitempty"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
	EapSuccess bool                               `json:"eapSuccess"`
	Eap        []byte                             `json:"eap,omitempty"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
