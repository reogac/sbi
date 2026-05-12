/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InformationTransferReqData struct {
	TaiList           []Tai             `json:"taiList,omitempty"`
	RatSelector       RatSelector       `json:"ratSelector,omitempty"`
	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`
	N2Information     N2InfoContainer   `json:"n2Information"`
	SupportedFeatures string            `json:"supportedFeatures,omitempty"`
}
