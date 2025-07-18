/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AcknowledgeInfo struct {
	UpuTransparentContainer string `json:"upuTransparentContainer,omitempty"`
	SorMacIue               string `json:"sorMacIue,omitempty"`
	UpuMacIue               string `json:"upuMacIue,omitempty"`
	ProvisioningTime        string `json:"provisioningTime"`
	SorTransparentContainer string `json:"sorTransparentContainer,omitempty"`
	UeNotReachable          *bool  `json:"ueNotReachable,omitempty"`
}
