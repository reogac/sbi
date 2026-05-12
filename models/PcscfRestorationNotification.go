/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:39 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PcscfRestorationNotification struct {
	Supi        string        `json:"supi"`
	FailedPcscf *PcscfAddress `json:"failedPcscf,omitempty"`
}
