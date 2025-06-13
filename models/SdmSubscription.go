/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
}
