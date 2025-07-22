/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Supis            []string              `json:"supis,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
}
