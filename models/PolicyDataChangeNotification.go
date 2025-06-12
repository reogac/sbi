/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
}
