/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalSnssaiData struct {
	RequiredAuthnAuthz   *bool     `json:"requiredAuthnAuthz,omitempty"`
	SubscribedUeSliceMbr *SliceMbr `json:"subscribedUeSliceMbr,omitempty"`
	SubscribedNsSrgList  []string  `json:"subscribedNsSrgList,omitempty"`
}
