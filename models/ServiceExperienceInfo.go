/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Supis            []string              `json:"supis,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
}
