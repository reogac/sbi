/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 11:15:54 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IpSegmentKey struct {
	Family          IpFamily `json:"family"`
	IpIndex         *int32   `json:"ipIndex,omitempty"`
	Dnn             string   `json:"dnn"`
	NetworkInstance string   `json:"networkInstance"`
	Dnai            string   `json:"dnai,omitempty"`
}
