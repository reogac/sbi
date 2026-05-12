/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Supis            []string              `json:"supis,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
}
