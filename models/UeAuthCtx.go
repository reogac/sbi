/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeAuthCtx struct {
	Supi       string                             `json:"supi"`
	PlmnId     PlmnId                             `json:"plmnId"`
	Eap        []byte                             `json:"eap,omitempty"`
	EapSuccess bool                               `json:"eapSuccess"`
	AmData     *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	Kamf       []byte                             `json:"kamf,omitempty"`
	NgKsi      NgKsi                              `json:"ngKsi"`
}
