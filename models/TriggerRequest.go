/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TriggerRequest struct {
	Supi        string        `json:"supi"`
	FailedPcscf *PcscfAddress `json:"failedPcscf,omitempty"`
}
