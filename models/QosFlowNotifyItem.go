/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowNotifyItem struct {
	Qfi                    int               `json:"qfi"`
	NotificationCause      NotificationCause `json:"notificationCause"`
	CurrentQosProfileIndex *int              `json:"currentQosProfileIndex,omitempty"`
	NullQoSProfileIndex    *bool             `json:"nullQoSProfileIndex,omitempty"`
}
