/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Dnai             string                `json:"dnai,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	AppId            string                `json:"appId,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
}
