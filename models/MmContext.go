/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MmContext struct {
	WagfId                  *GlobalRanNodeId       `json:"wagfId,omitempty"`
	NasUplinkCount          *int                   `json:"nasUplinkCount,omitempty"`
	NssaiMappingList        []NssaiMapping         `json:"nssaiMappingList,omitempty"`
	UeDifferentiationInfo   *UeDifferentiationInfo `json:"ueDifferentiationInfo,omitempty"`
	AnN2ApId                *int                   `json:"anN2ApId,omitempty"`
	PendingNssaiMappingList []NssaiMapping         `json:"pendingNssaiMappingList,omitempty"`
	NasSecurityMode         *NasSecurityMode       `json:"nasSecurityMode,omitempty"`
	AllowedHomeNssai        []Snssai               `json:"allowedHomeNssai,omitempty"`
	UeSecurityCapability    string                 `json:"ueSecurityCapability,omitempty"`
	AllowedNssai            []Snssai               `json:"allowedNssai,omitempty"`
	NsInstanceList          []string               `json:"nsInstanceList,omitempty"`
	ExpectedUEbehavior      *ExpectedUeBehavior    `json:"expectedUEbehavior,omitempty"`
	TngfId                  *GlobalRanNodeId       `json:"tngfId,omitempty"`
	EpsNasSecurityMode      *EpsNasSecurityMode    `json:"epsNasSecurityMode,omitempty"`
	NasDownlinkCount        *int                   `json:"nasDownlinkCount,omitempty"`
	PlmnAssiUeRadioCapId    string                 `json:"plmnAssiUeRadioCapId,omitempty"`
	ManAssiUeRadioCapId     string                 `json:"manAssiUeRadioCapId,omitempty"`
	UcmfDicEntryId          string                 `json:"ucmfDicEntryId,omitempty"`
	N3IwfId                 *GlobalRanNodeId       `json:"n3IwfId,omitempty"`
	NssaaStatusList         []NssaaStatus          `json:"nssaaStatusList,omitempty"`
	UuaaMmStatus            UuaaMmStatus           `json:"uuaaMmStatus,omitempty"`
	AccessType              AccessType             `json:"accessType"`
	S1UeNetworkCapability   string                 `json:"s1UeNetworkCapability,omitempty"`
}
