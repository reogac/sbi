/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type StatusInfo struct {
	ResourceStatus    ResourceStatus     `json:"resourceStatus"`
	Cause             Cause              `json:"cause,omitempty"`
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType            AccessType         `json:"anType,omitempty"`
}
