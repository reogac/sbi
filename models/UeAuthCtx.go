/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Rand       []byte                             `json:"rand,omitempty"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	Supi       string                             `json:"supi"`
	Kamf       []byte                             `json:"kamf,omitempty"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
