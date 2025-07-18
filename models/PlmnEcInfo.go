/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PlmnEcInfo struct {
	EcRestrictionDataWb *EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`
	EcRestrictionDataNb *bool                `json:"ecRestrictionDataNb,omitempty"`
	PlmnId              PlmnId               `json:"plmnId"`
}
