/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowNotifyItem struct {
	NotificationCause      NotificationCause `json:"notificationCause"`
	CurrentQosProfileIndex *int              `json:"currentQosProfileIndex,omitempty"`
	NullQoSProfileIndex    *bool             `json:"nullQoSProfileIndex,omitempty"`
	Qfi                    int               `json:"qfi"`
}
