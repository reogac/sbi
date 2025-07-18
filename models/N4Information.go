/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N4Information struct {
	N4MessageType    N4MessageType    `json:"n4MessageType"`
	N4MessagePayload RefToBinaryData  `json:"n4MessagePayload"`
	N4DnaiInfo       *DnaiInformation `json:"n4DnaiInfo,omitempty"`
	PsaUpfId         string           `json:"psaUpfId,omitempty"`
	UlClBpId         string           `json:"ulClBpId,omitempty"`
	N9UlPdrIdList    []int            `json:"n9UlPdrIdList,omitempty"`
}
