/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeGroupProfileData struct {
	RestrictedEventTypes []string               `json:"restrictedEventTypes,omitempty"`
	AllowedMtcProvider   map[string]MtcProvider `json:"allowedMtcProvider,omitempty"`
	SupportedFeatures    string                 `json:"supportedFeatures,omitempty"`
	IwkEpcRestricted     *bool                  `json:"iwkEpcRestricted,omitempty"`
	ExtGroupId           string                 `json:"extGroupId,omitempty"`
	HssGroupId           string                 `json:"hssGroupId,omitempty"`
}
