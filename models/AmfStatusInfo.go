/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfStatusInfo struct {
	GuamiList        []Guami      `json:"guamiList"`
	StatusChange     StatusChange `json:"statusChange"`
	TargetAmfRemoval string       `json:"targetAmfRemoval,omitempty"`
	TargetAmfFailure string       `json:"targetAmfFailure,omitempty"`
}
