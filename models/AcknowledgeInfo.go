/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:37 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AcknowledgeInfo struct {
	SorTransparentContainer string `json:"sorTransparentContainer,omitempty"`
	UeNotReachable          *bool  `json:"ueNotReachable,omitempty"`
	UpuTransparentContainer string `json:"upuTransparentContainer,omitempty"`
	SorMacIue               string `json:"sorMacIue,omitempty"`
	UpuMacIue               string `json:"upuMacIue,omitempty"`
	ProvisioningTime        string `json:"provisioningTime"`
}
