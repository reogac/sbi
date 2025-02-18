/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:05:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models
type N2SmInfoUplinkContent struct {
	 N2SmInfoType	N2SmInfoType	`json:"n2SmInfoType,omitempty"`
	 SessionId	int16	`json:"sessionId"`
	 N2SmInfo	[]byte	`json:"n2SmInfo,omitempty"`
}
