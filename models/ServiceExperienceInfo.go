/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Supis            []string              `json:"supis,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
}
