/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
