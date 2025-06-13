/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
}
