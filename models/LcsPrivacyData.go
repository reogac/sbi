/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LcsPrivacyData struct {
	UnrelatedClass      *UnrelatedClass     `json:"unrelatedClass,omitempty"`
	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`
	Lpi                 *Lpi                `json:"lpi,omitempty"`
}
