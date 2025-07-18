/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	RoutingId                     string          `json:"routingId,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
}
