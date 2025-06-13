/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:21 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AcknowledgeInfo struct {
	UeNotReachable          *bool  `json:"ueNotReachable,omitempty"`
	UpuTransparentContainer string `json:"upuTransparentContainer,omitempty"`
	SorMacIue               string `json:"sorMacIue,omitempty"`
	UpuMacIue               string `json:"upuMacIue,omitempty"`
	ProvisioningTime        string `json:"provisioningTime"`
	SorTransparentContainer string `json:"sorTransparentContainer,omitempty"`
}
