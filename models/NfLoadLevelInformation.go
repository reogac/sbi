/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NfLoadLevelInformation struct {
	Confidence         *int      `json:"confidence,omitempty"`
	NfType             NFType    `json:"nfType,omitempty"`
	NfSetId            string    `json:"nfSetId,omitempty"`
	NfCpuUsage         *int      `json:"nfCpuUsage,omitempty"`
	Snssai             *Snssai   `json:"snssai,omitempty"`
	NfLoadLevelAverage *int      `json:"nfLoadLevelAverage,omitempty"`
	NfLoadLevelpeak    *int      `json:"nfLoadLevelpeak,omitempty"`
	NfLoadAvgInAoi     *int      `json:"nfLoadAvgInAoi,omitempty"`
	NfInstanceId       string    `json:"nfInstanceId,omitempty"`
	NfStatus           *NfStatus `json:"nfStatus,omitempty"`
	NfMemoryUsage      *int      `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage     *int      `json:"nfStorageUsage,omitempty"`
}
