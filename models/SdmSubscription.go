/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	NfInstanceId               string                       `json:"nfInstanceId"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
}
