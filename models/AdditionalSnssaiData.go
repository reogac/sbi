/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalSnssaiData struct {
	SubscribedNsSrgList  []string  `json:"subscribedNsSrgList,omitempty"`
	RequiredAuthnAuthz   *bool     `json:"requiredAuthnAuthz,omitempty"`
	SubscribedUeSliceMbr *SliceMbr `json:"subscribedUeSliceMbr,omitempty"`
}
