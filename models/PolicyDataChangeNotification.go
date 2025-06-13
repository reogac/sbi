/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	DelResources            []string                                 `json:"delResources,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
}
