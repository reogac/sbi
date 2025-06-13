/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
}
