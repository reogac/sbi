/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type OperatorSpecificDataContainer struct {
	Value              Value    `json:"value"`
	SupportedFeatures  string   `json:"supportedFeatures,omitempty"`
	ResetIds           []string `json:"resetIds,omitempty"`
	DataType           DataType `json:"dataType"`
	DataTypeDefinition string   `json:"dataTypeDefinition,omitempty"`
}
