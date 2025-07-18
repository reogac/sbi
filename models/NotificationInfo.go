/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NotificationInfo struct {
	NotifId     string `json:"notifId"`
	NotifUri    string `json:"notifUri"`
	UpBufferInd *bool  `json:"upBufferInd,omitempty"`
}
