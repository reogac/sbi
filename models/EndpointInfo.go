/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EndpointInfo struct {
	SbiUrl string `json:"sbiUrl"`
	Name   string `json:"name,omitempty"`
	GwUrl  string `json:"gwUrl"`
	GwName string `json:"gwName,omitempty"`
	Uuid   string `json:"uuid"`
}
