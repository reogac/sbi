/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	Expires                    string                       `json:"expires,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
}
