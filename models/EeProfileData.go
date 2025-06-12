/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeProfileData struct {
	HssGroupId           string                 `json:"hssGroupId,omitempty"`
	RestrictedEventTypes []string               `json:"restrictedEventTypes,omitempty"`
	SupportedFeatures    string                 `json:"supportedFeatures,omitempty"`
	AllowedMtcProvider   map[string]MtcProvider `json:"allowedMtcProvider,omitempty"`
	IwkEpcRestricted     *bool                  `json:"iwkEpcRestricted,omitempty"`
	Imsi                 string                 `json:"imsi,omitempty"`
}
