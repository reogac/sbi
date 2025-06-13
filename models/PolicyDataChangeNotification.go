/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
}
