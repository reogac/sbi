/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	Supi                          string          `json:"supi,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
}
