/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Kamf       []byte                             `json:"kamf,omitempty"`
	Rand       []byte                             `json:"rand,omitempty"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	Supi       string                             `json:"supi"`
	PlmnId     PlmnId                             `json:"plmnId"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
}
