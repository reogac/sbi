/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthenticationSubscription struct {
	Supi                          string          `json:"supi,omitempty"`
	EncPermanentKey               string          `json:"encPermanentKey,omitempty"`
	ProtectionParameterId         string          `json:"protectionParameterId,omitempty"`
	AuthenticationManagementField string          `json:"authenticationManagementField,omitempty"`
	EncOpcKey                     string          `json:"encOpcKey,omitempty"`
	VectorGenerationInHss         *bool           `json:"vectorGenerationInHss,omitempty"`
	AuthenticationMethod          AuthMethod      `json:"authenticationMethod"`
	EncTopcKey                    string          `json:"encTopcKey,omitempty"`
	HssGroupId                    string          `json:"hssGroupId,omitempty"`
	AkmaAllowed                   *bool           `json:"akmaAllowed,omitempty"`
	RoutingId                     string          `json:"routingId,omitempty"`
	RgAuthenticationInd           *bool           `json:"rgAuthenticationInd,omitempty"`
	SequenceNumber                *SequenceNumber `json:"sequenceNumber,omitempty"`
	AlgorithmId                   string          `json:"algorithmId,omitempty"`
	N5gcAuthMethod                AuthMethod      `json:"n5gcAuthMethod,omitempty"`
}
