/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
}
