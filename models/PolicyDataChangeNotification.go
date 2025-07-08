/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
}
