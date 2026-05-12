/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeProfileData struct {
	AllowedMtcProvider   map[string]MtcProvider `json:"allowedMtcProvider,omitempty"`
	IwkEpcRestricted     *bool                  `json:"iwkEpcRestricted,omitempty"`
	Imsi                 string                 `json:"imsi,omitempty"`
	HssGroupId           string                 `json:"hssGroupId,omitempty"`
	RestrictedEventTypes []string               `json:"restrictedEventTypes,omitempty"`
	SupportedFeatures    string                 `json:"supportedFeatures,omitempty"`
}
