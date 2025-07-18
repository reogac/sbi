/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
}
