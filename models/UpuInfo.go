/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpuInfo struct {
	UpuTransparentContainer string    `json:"upuTransparentContainer,omitempty"`
	UpuDataList             []UpuData `json:"upuDataList,omitempty"`
	UpuRegInd               *bool     `json:"upuRegInd,omitempty"`
	UpuAckInd               *bool     `json:"upuAckInd,omitempty"`
	UpuMacIausf             string    `json:"upuMacIausf,omitempty"`
	CounterUpu              string    `json:"counterUpu,omitempty"`
	ProvisioningTime        string    `json:"provisioningTime"`
}
