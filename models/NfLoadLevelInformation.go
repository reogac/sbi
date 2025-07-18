/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NfLoadLevelInformation struct {
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	NfType             NFType    `json:"nfType,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	Confidence         *int      `json:"confidence,omitempty"`
}
