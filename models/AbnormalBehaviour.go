/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AbnormalBehaviour struct {
	Supis        []string               `json:"supis,omitempty"`
	Excep        Exception              `json:"excep"`
	Dnn          string                 `json:"dnn,omitempty"`
	Snssai       *Snssai                `json:"snssai,omitempty"`
	Ratio        *int                   `json:"ratio,omitempty"`
	Confidence   *int                   `json:"confidence,omitempty"`
	AddtMeasInfo *AdditionalMeasurement `json:"addtMeasInfo,omitempty"`
}
