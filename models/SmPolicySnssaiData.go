/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicySnssaiData struct {
	UeSliceMbr      *SliceMbr                  `json:"ueSliceMbr,omitempty"`
	Snssai          Snssai                     `json:"snssai"`
	SmPolicyDnnData map[string]SmPolicyDnnData `json:"smPolicyDnnData,omitempty"`
}
