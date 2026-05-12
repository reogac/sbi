/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeGroupProfileData struct {
	ExtGroupId           string                 `json:"extGroupId,omitempty"`
	HssGroupId           string                 `json:"hssGroupId,omitempty"`
	RestrictedEventTypes []string               `json:"restrictedEventTypes,omitempty"`
	AllowedMtcProvider   map[string]MtcProvider `json:"allowedMtcProvider,omitempty"`
	SupportedFeatures    string                 `json:"supportedFeatures,omitempty"`
	IwkEpcRestricted     *bool                  `json:"iwkEpcRestricted,omitempty"`
}
