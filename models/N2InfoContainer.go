/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InfoContainer struct {
	N2InformationClass N2InformationClass `json:"n2InformationClass"`
	SmInfo             *N2SmInformation   `json:"smInfo,omitempty"`
	RanInfo            *N2RanInformation  `json:"ranInfo,omitempty"`
	NrppaInfo          *NrppaInformation  `json:"nrppaInfo,omitempty"`
	PwsInfo            *PwsInformation    `json:"pwsInfo,omitempty"`
	V2xInfo            *V2xInformation    `json:"v2xInfo,omitempty"`
	ProseInfo          *ProSeInformation  `json:"proseInfo,omitempty"`
}
