/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicySnssaiData struct {
	SmPolicyDnnData map[string]SmPolicyDnnData `json:"smPolicyDnnData,omitempty"`
	UeSliceMbr      *SliceMbr                  `json:"ueSliceMbr,omitempty"`
	Snssai          Snssai                     `json:"snssai"`
}
