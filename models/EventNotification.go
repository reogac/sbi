/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventNotification struct {
	NfLoadLevelInfos   []NfLoadLevelInformation       `json:"nfLoadLevelInfos,omitempty"`
	QosSustainInfos    []QosSustainabilityInfo        `json:"qosSustainInfos,omitempty"`
	AbnorBehavrs       []AbnormalBehaviour            `json:"abnorBehavrs,omitempty"`
	DisperInfos        []DispersionInfo               `json:"disperInfos,omitempty"`
	RedTransInfos      []RedundantTransmissionExpInfo `json:"redTransInfos,omitempty"`
	Start              string                         `json:"start,omitempty"`
	Expiry             string                         `json:"expiry,omitempty"`
	AnaMetaInfo        *AnalyticsMetadataInfo         `json:"anaMetaInfo,omitempty"`
	UeMobs             []UeMobility                   `json:"ueMobs,omitempty"`
	WlanInfos          []WlanPerformanceInfo          `json:"wlanInfos,omitempty"`
	SmccExps           []SmcceInfo                    `json:"smccExps,omitempty"`
	Event              NwdafEvent                     `json:"event"`
	FailNotifyCode     NwdafFailureCode               `json:"failNotifyCode,omitempty"`
	NsiLoadLevelInfos  []NsiLoadLevelInfo             `json:"nsiLoadLevelInfos,omitempty"`
	TimeStampGen       string                         `json:"timeStampGen,omitempty"`
	UserDataCongInfos  []UserDataCongestionInfo       `json:"userDataCongInfos,omitempty"`
	UeComms            []UeCommunication              `json:"ueComms,omitempty"`
	NwPerfs            []NetworkPerfInfo              `json:"nwPerfs,omitempty"`
	DnPerfInfos        []DnPerfInfo                   `json:"dnPerfInfos,omitempty"`
	RvWaitTime         *int                           `json:"rvWaitTime,omitempty"`
	SliceLoadLevelInfo *SliceLoadLevelInformation     `json:"sliceLoadLevelInfo,omitempty"`
	SvcExps            []ServiceExperienceInfo        `json:"svcExps,omitempty"`
}
