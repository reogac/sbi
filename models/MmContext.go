/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MmContext struct {
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	AccessType              AccessType             `json:"accessType"`
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	UuaaMmStatus            UuaaMmStatus           `json:"uuaaMmStatus,omitempty"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
}
