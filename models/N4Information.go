/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N4Information struct {
	N4DnaiInfo       *DnaiInformation `json:"n4DnaiInfo,omitempty"`
	PsaUpfId         string           `json:"psaUpfId,omitempty"`
	UlClBpId         string           `json:"ulClBpId,omitempty"`
	N9UlPdrIdList    []int            `json:"n9UlPdrIdList,omitempty"`
	N4MessageType    N4MessageType    `json:"n4MessageType"`
	N4MessagePayload RefToBinaryData  `json:"n4MessagePayload"`
}
