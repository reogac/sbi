/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Dnn              string                `json:"dnn,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Dnai             string                `json:"dnai,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
}
