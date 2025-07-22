/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NotificationInfo struct {
	NotifId     string `json:"notifId"`
	NotifUri    string `json:"notifUri"`
	UpBufferInd *bool  `json:"upBufferInd,omitempty"`
}
