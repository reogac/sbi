/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	Event              NwdafEvent                    `json:"event"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
}
