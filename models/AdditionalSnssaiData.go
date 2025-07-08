/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalSnssaiData struct {
	RequiredAuthnAuthz   *bool     `json:"requiredAuthnAuthz,omitempty"`
	SubscribedUeSliceMbr *SliceMbr `json:"subscribedUeSliceMbr,omitempty"`
	SubscribedNsSrgList  []string  `json:"subscribedNsSrgList,omitempty"`
}
