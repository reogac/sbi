/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
}
