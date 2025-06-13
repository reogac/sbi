/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AfCoordinationInfo struct {
	NotificationInfoList []NotificationInfo `json:"notificationInfoList,omitempty"`
	SourceDnai           string             `json:"sourceDnai,omitempty"`
	SourceUeIpv4Addr     string             `json:"sourceUeIpv4Addr,omitempty"`
	SourceUeIpv6Prefix   string             `json:"sourceUeIpv6Prefix,omitempty"`
}
