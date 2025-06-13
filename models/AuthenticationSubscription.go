/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	Supi                          string          `json:"supi,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
}
