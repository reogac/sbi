/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationNotification struct {
	RanNodeId              *GlobalRanNodeId   `json:"ranNodeId,omitempty"`
	ToReleaseSessionList   []int              `json:"toReleaseSessionList,omitempty"`
	N2InfoContainer        *N2InfoContainer   `json:"n2InfoContainer,omitempty"`
	LcsCorrelationId       string             `json:"lcsCorrelationId,omitempty"`
	NotifyReason           N2InfoNotifyReason `json:"notifyReason,omitempty"`
	SmfChangeInfoList      []SmfChangeInfo    `json:"smfChangeInfoList,omitempty"`
	InitialAmfName         string             `json:"initialAmfName,omitempty"`
	AnN2IPv4Addr           string             `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr           string             `json:"anN2IPv6Addr,omitempty"`
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	NotifySourceNgRan      *bool              `json:"notifySourceNgRan,omitempty"`
	Guami                  *Guami             `json:"guami,omitempty"`
}
