/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SdmSubscription struct {
	ContextInfo                *ContextInfo                 `json:"contextInfo,omitempty"`
	ResetIds                   []string                     `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                       `json:"dataRestorationCallbackUri,omitempty"`
	ImplicitUnsubscribe        *bool                        `json:"implicitUnsubscribe,omitempty"`
	SubscriptionId             string                       `json:"subscriptionId,omitempty"`
	SupportedFeatures          string                       `json:"supportedFeatures,omitempty"`
	PlmnId                     *PlmnId                      `json:"plmnId,omitempty"`
	Report                     *ImmediateReport             `json:"report,omitempty"`
	NfChangeFilter             *bool                        `json:"nfChangeFilter,omitempty"`
	ImmediateReport            *bool                        `json:"immediateReport,omitempty"`
	UeConSmfDataSubFilter      *UeContextInSmfDataSubFilter `json:"ueConSmfDataSubFilter,omitempty"`
	UdrRestartInd              *bool                        `json:"udrRestartInd,omitempty"`
	Expires                    string                       `json:"expires,omitempty"`
	CallbackReference          string                       `json:"callbackReference"`
	MonitoredResourceUris      []string                     `json:"monitoredResourceUris"`
	Dnn                        string                       `json:"dnn,omitempty"`
	UniqueSubscription         *bool                        `json:"uniqueSubscription,omitempty"`
	NfInstanceId               string                       `json:"nfInstanceId"`
	AmfServiceName             ServiceName                  `json:"amfServiceName,omitempty"`
	SingleNssai                *Snssai                      `json:"singleNssai,omitempty"`
}
