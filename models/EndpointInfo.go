/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:04 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EndpointInfo struct {
	Uuid   string `json:"uuid"`
	SbiUrl string `json:"sbiUrl"`
	Name   string `json:"name,omitempty"`
	GwUrl  string `json:"gwUrl"`
	GwName string `json:"gwName,omitempty"`
}
