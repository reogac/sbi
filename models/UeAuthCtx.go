/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Supi       string                             `json:"supi"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Kamf       []byte                             `json:"kamf,omitempty"`
	Rand       []byte                             `json:"rand,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
	Eap        []byte                             `json:"eap,omitempty"`
	NgKsi      NgKsi                              `json:"ngKsi"`
}
