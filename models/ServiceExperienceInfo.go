/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Dnai             string                `json:"dnai,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
}
