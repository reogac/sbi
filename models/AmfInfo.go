/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfInfo struct {
	Guami         Guami      `json:"guami"`
	AccessType    AccessType `json:"accessType,omitempty"`
	AmfInstanceId string     `json:"amfInstanceId"`
}
