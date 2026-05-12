/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NotificationInfo struct {
	NotifId     string `json:"notifId"`
	NotifUri    string `json:"notifUri"`
	UpBufferInd *bool  `json:"upBufferInd,omitempty"`
}
