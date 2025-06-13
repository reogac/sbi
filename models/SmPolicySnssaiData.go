/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicySnssaiData struct {
	Snssai          Snssai                     `json:"snssai"`
	SmPolicyDnnData map[string]SmPolicyDnnData `json:"smPolicyDnnData,omitempty"`
	UeSliceMbr      *SliceMbr                  `json:"ueSliceMbr,omitempty"`
}
