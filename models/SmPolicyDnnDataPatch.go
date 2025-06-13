/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnDataPatch struct {
	Dnn       string            `json:"dnn"`
	BdtRefIds map[string]string `json:"bdtRefIds,omitempty"`
}
