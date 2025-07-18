/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringData struct {
	WaitTime        *int     `json:"waitTime,omitempty"`
	RepPeriod       *int     `json:"repPeriod,omitempty"`
	NotifyUri       string   `json:"notifyUri,omitempty"`
	RepFreqs        []string `json:"repFreqs"`
	RepThreshDl     *int     `json:"repThreshDl,omitempty"`
	RepThreshUl     *int     `json:"repThreshUl,omitempty"`
	RepThreshRp     *int     `json:"repThreshRp,omitempty"`
	QmId            string   `json:"qmId"`
	ReqQosMonParams []string `json:"reqQosMonParams"`
	NotifyCorreId   string   `json:"notifyCorreId,omitempty"`
	DirectNotifInd  *bool    `json:"directNotifInd,omitempty"`
}
