/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type OperatorSpecificDataContainer struct {
	DataTypeDefinition string   `json:"dataTypeDefinition,omitempty"`
	Value              Value    `json:"value"`
	SupportedFeatures  string   `json:"supportedFeatures,omitempty"`
	ResetIds           []string `json:"resetIds,omitempty"`
	DataType           DataType `json:"dataType"`
}
