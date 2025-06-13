/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationNotification struct {
	SmfChangeInfoList      []SmfChangeInfo    `json:"smfChangeInfoList,omitempty"`
	AnN2IPv6Addr           string             `json:"anN2IPv6Addr,omitempty"`
	N2InfoContainer        *N2InfoContainer   `json:"n2InfoContainer,omitempty"`
	ToReleaseSessionList   []int              `json:"toReleaseSessionList,omitempty"`
	NotifyReason           N2InfoNotifyReason `json:"notifyReason,omitempty"`
	InitialAmfName         string             `json:"initialAmfName,omitempty"`
	AnN2IPv4Addr           string             `json:"anN2IPv4Addr,omitempty"`
	Guami                  *Guami             `json:"guami,omitempty"`
	NotifySourceNgRan      *bool              `json:"notifySourceNgRan,omitempty"`
	N2NotifySubscriptionId string             `json:"n2NotifySubscriptionId"`
	LcsCorrelationId       string             `json:"lcsCorrelationId,omitempty"`
	RanNodeId              *GlobalRanNodeId   `json:"ranNodeId,omitempty"`
}
