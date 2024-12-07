/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type OperatorSpecificDataContainer struct {
	ResetIds           []string `json:"resetIds,omitempty"`
	DataType           DataType `json:"dataType"`
	DataTypeDefinition string   `json:"dataTypeDefinition,omitempty"`
	Value              Value    `json:"value"`
	SupportedFeatures  string   `json:"supportedFeatures,omitempty"`
}
