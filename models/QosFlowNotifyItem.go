/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowNotifyItem struct {
	CurrentQosProfileIndex *int              `json:"currentQosProfileIndex,omitempty"`
	NullQoSProfileIndex    *bool             `json:"nullQoSProfileIndex,omitempty"`
	Qfi                    int               `json:"qfi"`
	NotificationCause      NotificationCause `json:"notificationCause"`
}
