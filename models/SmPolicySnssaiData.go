/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicySnssaiData struct {
	Snssai          Snssai                     `json:"snssai"`
	SmPolicyDnnData map[string]SmPolicyDnnData `json:"smPolicyDnnData,omitempty"`
	UeSliceMbr      *SliceMbr                  `json:"ueSliceMbr,omitempty"`
}
