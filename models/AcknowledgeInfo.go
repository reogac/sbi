/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AcknowledgeInfo struct {
	SorMacIue               string `json:"sorMacIue,omitempty"`
	UpuMacIue               string `json:"upuMacIue,omitempty"`
	ProvisioningTime        string `json:"provisioningTime"`
	SorTransparentContainer string `json:"sorTransparentContainer,omitempty"`
	UeNotReachable          *bool  `json:"ueNotReachable,omitempty"`
	UpuTransparentContainer string `json:"upuTransparentContainer,omitempty"`
}
