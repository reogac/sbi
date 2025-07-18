/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
