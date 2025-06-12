/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type StatusInfo struct {
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType            AccessType         `json:"anType,omitempty"`
	ResourceStatus    ResourceStatus     `json:"resourceStatus"`
	Cause             Cause              `json:"cause,omitempty"`
}
