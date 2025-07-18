/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeProfileData struct {
	RestrictedEventTypes []string               `json:"restrictedEventTypes,omitempty"`
	SupportedFeatures    string                 `json:"supportedFeatures,omitempty"`
	AllowedMtcProvider   map[string]MtcProvider `json:"allowedMtcProvider,omitempty"`
	IwkEpcRestricted     *bool                  `json:"iwkEpcRestricted,omitempty"`
	Imsi                 string                 `json:"imsi,omitempty"`
	HssGroupId           string                 `json:"hssGroupId,omitempty"`
}
