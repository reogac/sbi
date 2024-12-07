/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2InfoContainer struct {
	SmInfo             *N2SmInformation   `json:"smInfo,omitempty"`
	RanInfo            *N2RanInformation  `json:"ranInfo,omitempty"`
	NrppaInfo          *NrppaInformation  `json:"nrppaInfo,omitempty"`
	PwsInfo            *PwsInformation    `json:"pwsInfo,omitempty"`
	V2xInfo            *V2xInformation    `json:"v2xInfo,omitempty"`
	ProseInfo          *ProSeInformation  `json:"proseInfo,omitempty"`
	N2InformationClass N2InformationClass `json:"n2InformationClass"`
}
