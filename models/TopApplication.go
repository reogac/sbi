/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TopApplication struct {
	Ratio           *int      `json:"ratio,omitempty"`
	AppId           string    `json:"appId,omitempty"`
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty"`
}
