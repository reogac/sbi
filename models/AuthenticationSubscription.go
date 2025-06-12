/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
}
