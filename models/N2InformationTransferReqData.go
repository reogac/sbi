/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationTransferReqData struct {
	SupportedFeatures string            `json:"supportedFeatures,omitempty"`
	TaiList           []Tai             `json:"taiList,omitempty"`
	RatSelector       RatSelector       `json:"ratSelector,omitempty"`
	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`
	N2Information     N2InfoContainer   `json:"n2Information"`
}
