/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Rand       []byte                             `json:"rand,omitempty"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Supi       string                             `json:"supi"`
	AuthType   AuthType                           `json:"authType,omitempty"`
	PlmnId     PlmnId                             `json:"plmnId"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	Kamf       []byte                             `json:"kamf,omitempty"`
}
