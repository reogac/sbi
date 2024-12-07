/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
}
