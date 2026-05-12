/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:28 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	PlmnId     PlmnId                             `json:"plmnId"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	NgKsi      NgKsi                              `json:"ngKsi"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Supi       string                             `json:"supi"`
	Kamf       []byte                             `json:"kamf,omitempty"`
}
