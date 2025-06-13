/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Supi       string                             `json:"supi"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	Kamf       []byte                             `json:"kamf,omitempty"`
	Rand       []byte                             `json:"rand,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
}
