/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
}
