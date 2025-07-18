/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VsmfUpdatedData struct {
	QosFlowsRelList             []QosFlowItem             `json:"qosFlowsRelList,omitempty"`
	N1SmInfoFromUe              *RefToBinaryData          `json:"n1SmInfoFromUe,omitempty"`
	UeTimeZone                  string                    `json:"ueTimeZone,omitempty"`
	SecondaryRatUsageInfo       []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4InfoExt3                  *N4Information            `json:"n4InfoExt3,omitempty"`
	N4InfoExt2                  *N4Information            `json:"n4InfoExt2,omitempty"`
	QosFlowsAddModList          []QosFlowItem             `json:"qosFlowsAddModList,omitempty"`
	UnknownN1SmInfo             *RefToBinaryData          `json:"unknownN1SmInfo,omitempty"`
	UeLocation                  *UserLocation             `json:"ueLocation,omitempty"`
	AddUeLocation               *UserLocation             `json:"addUeLocation,omitempty"`
	FailedToAssignEbiList       []Arp                     `json:"failedToAssignEbiList,omitempty"`
	N4InfoExt1                  *N4Information            `json:"n4InfoExt1,omitempty"`
	QosFlowsFailedtoAddModList  []QosFlowItem             `json:"qosFlowsFailedtoAddModList,omitempty"`
	AssignedEbiList             []EbiArpMapping           `json:"assignedEbiList,omitempty"`
	ReleasedEbiList             []int                     `json:"releasedEbiList,omitempty"`
	N4Info                      *N4Information            `json:"n4Info,omitempty"`
	ModifiedEbiListNotDelivered *bool                     `json:"modifiedEbiListNotDelivered,omitempty"`
	QosFlowsFailedtoRelList     []QosFlowItem             `json:"qosFlowsFailedtoRelList,omitempty"`
	SecondaryRatUsageReport     []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
}
