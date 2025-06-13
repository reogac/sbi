/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	AuthType   AuthType                           `json:"authType,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	PlmnId     PlmnId                             `json:"plmnId"`
	Eap        []byte                             `json:"eap,omitempty"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Supi       string                             `json:"supi"`
	Kamf       []byte                             `json:"kamf,omitempty"`
	Rand       []byte                             `json:"rand,omitempty"`
}
