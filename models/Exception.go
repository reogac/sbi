/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Exception struct {
	ExcepId    ExceptionId    `json:"excepId"`
	ExcepLevel *int           `json:"excepLevel,omitempty"`
	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty"`
}
