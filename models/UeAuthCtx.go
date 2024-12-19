/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Kamf       []byte                             `json:"kamf,omitempty"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	Supi       string                             `json:"supi"`
	Rand       []byte                             `json:"rand,omitempty"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
