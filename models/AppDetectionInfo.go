/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AppDetectionInfo struct {
	AppId           string            `json:"appId"`
	InstanceId      string            `json:"instanceId,omitempty"`
	SdfDescriptions []FlowInformation `json:"sdfDescriptions,omitempty"`
}
