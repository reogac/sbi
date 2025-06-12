/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
}
