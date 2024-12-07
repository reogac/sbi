/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	AppId            string                `json:"appId,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
}
