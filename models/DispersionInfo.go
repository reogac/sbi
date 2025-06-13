/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionInfo struct {
	DisperCollects []DispersionCollection `json:"disperCollects"`
	DisperType     DispersionType         `json:"disperType"`
	TsStart        string                 `json:"tsStart"`
	TsDuration     int                    `json:"tsDuration"`
}
