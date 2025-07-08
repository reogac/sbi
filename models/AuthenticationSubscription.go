/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
}
