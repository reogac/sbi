/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	NfInstanceId               string                       `json:"nfInstanceId"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	Dnn                        string                       `json:"dnn,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
}
