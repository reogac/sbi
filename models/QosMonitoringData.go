/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringData struct {
	QmId            string   `json:"qmId"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	WaitTime        *int     `json:"waitTime,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
}
