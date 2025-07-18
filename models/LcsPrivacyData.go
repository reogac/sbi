/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LcsPrivacyData struct {
	Lpi                 *Lpi                `json:"lpi,omitempty"`
	UnrelatedClass      *UnrelatedClass     `json:"unrelatedClass,omitempty"`
	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`
}
