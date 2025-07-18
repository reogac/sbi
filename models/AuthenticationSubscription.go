/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
}
