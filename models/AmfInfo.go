/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfInfo struct {
	AccessType    AccessType `json:"accessType,omitempty"`
	AmfInstanceId string     `json:"amfInstanceId"`
	Guami         Guami      `json:"guami"`
}
