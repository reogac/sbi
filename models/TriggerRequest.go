/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TriggerRequest struct {
	Supi        string        `json:"supi"`
	FailedPcscf *PcscfAddress `json:"failedPcscf,omitempty"`
}
