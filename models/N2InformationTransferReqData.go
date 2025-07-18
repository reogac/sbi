/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationTransferReqData struct {
	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`
	N2Information     N2InfoContainer   `json:"n2Information"`
	SupportedFeatures string            `json:"supportedFeatures,omitempty"`
	TaiList           []Tai             `json:"taiList,omitempty"`
	RatSelector       RatSelector       `json:"ratSelector,omitempty"`
}
