/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TriggerType string

// Define constant values for TriggerType
const (
	TRIGGERTYPE_QUOTA_THRESHOLD                                  TriggerType = "QUOTA_THRESHOLD"
	TRIGGERTYPE_QHT                                              TriggerType = "QHT"
	TRIGGERTYPE_FINAL                                            TriggerType = "FINAL"
	TRIGGERTYPE_QUOTA_EXHAUSTED                                  TriggerType = "QUOTA_EXHAUSTED"
	TRIGGERTYPE_VALIDITY_TIME                                    TriggerType = "VALIDITY_TIME"
	TRIGGERTYPE_OTHER_QUOTA_TYPE                                 TriggerType = "OTHER_QUOTA_TYPE"
	TRIGGERTYPE_FORCED_REAUTHORISATION                           TriggerType = "FORCED_REAUTHORISATION"
	TRIGGERTYPE_UNUSED_QUOTA_TIMER                               TriggerType = "UNUSED_QUOTA_TIMER"
	TRIGGERTYPE_UNIT_COUNT_INACTIVITY_TIMER                      TriggerType = "UNIT_COUNT_INACTIVITY_TIMER"
	TRIGGERTYPE_ABNORMAL_RELEASE                                 TriggerType = "ABNORMAL_RELEASE"
	TRIGGERTYPE_QOS_CHANGE                                       TriggerType = "QOS_CHANGE"
	TRIGGERTYPE_VOLUME_LIMIT                                     TriggerType = "VOLUME_LIMIT"
	TRIGGERTYPE_TIME_LIMIT                                       TriggerType = "TIME_LIMIT"
	TRIGGERTYPE_EVENT_LIMIT                                      TriggerType = "EVENT_LIMIT"
	TRIGGERTYPE_PLMN_CHANGE                                      TriggerType = "PLMN_CHANGE"
	TRIGGERTYPE_USER_LOCATION_CHANGE                             TriggerType = "USER_LOCATION_CHANGE"
	TRIGGERTYPE_RAT_CHANGE                                       TriggerType = "RAT_CHANGE"
	TRIGGERTYPE_SESSION_AMBR_CHANGE                              TriggerType = "SESSION_AMBR_CHANGE"
	TRIGGERTYPE_UE_TIMEZONE_CHANGE                               TriggerType = "UE_TIMEZONE_CHANGE"
	TRIGGERTYPE_TARIFF_TIME_CHANGE                               TriggerType = "TARIFF_TIME_CHANGE"
	TRIGGERTYPE_MAX_NUMBER_OF_CHANGES_IN_CHARGING_CONDITIONS     TriggerType = "MAX_NUMBER_OF_CHANGES_IN_CHARGING_CONDITIONS"
	TRIGGERTYPE_MANAGEMENT_INTERVENTION                          TriggerType = "MANAGEMENT_INTERVENTION"
	TRIGGERTYPE_CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA TriggerType = "CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA"
	TRIGGERTYPE_CHANGE_OF_3GPP_PS_DATA_OFF_STATUS                TriggerType = "CHANGE_OF_3GPP_PS_DATA_OFF_STATUS"
	TRIGGERTYPE_SERVING_NODE_CHANGE                              TriggerType = "SERVING_NODE_CHANGE"
	TRIGGERTYPE_REMOVAL_OF_UPF                                   TriggerType = "REMOVAL_OF_UPF"
	TRIGGERTYPE_ADDITION_OF_UPF                                  TriggerType = "ADDITION_OF_UPF"
	TRIGGERTYPE_INSERTION_OF_ISMF                                TriggerType = "INSERTION_OF_ISMF"
	TRIGGERTYPE_REMOVAL_OF_ISMF                                  TriggerType = "REMOVAL_OF_ISMF"
	TRIGGERTYPE_CHANGE_OF_ISMF                                   TriggerType = "CHANGE_OF_ISMF"
	TRIGGERTYPE_START_OF_SERVICE_DATA_FLOW                       TriggerType = "START_OF_SERVICE_DATA_FLOW"
	TRIGGERTYPE_ECGI_CHANGE                                      TriggerType = "ECGI_CHANGE"
	TRIGGERTYPE_TAI_CHANGE                                       TriggerType = "TAI_CHANGE"
	TRIGGERTYPE_HANDOVER_CANCEL                                  TriggerType = "HANDOVER_CANCEL"
	TRIGGERTYPE_HANDOVER_START                                   TriggerType = "HANDOVER_START"
	TRIGGERTYPE_HANDOVER_COMPLETE                                TriggerType = "HANDOVER_COMPLETE"
	TRIGGERTYPE_GFBR_GUARANTEED_STATUS_CHANGE                    TriggerType = "GFBR_GUARANTEED_STATUS_CHANGE"
	TRIGGERTYPE_ADDITION_OF_ACCESS                               TriggerType = "ADDITION_OF_ACCESS"
	TRIGGERTYPE_REMOVAL_OF_ACCESS                                TriggerType = "REMOVAL_OF_ACCESS"
	TRIGGERTYPE_START_OF_SDF_ADDITIONAL_ACCESS                   TriggerType = "START_OF_SDF_ADDITIONAL_ACCESS"
	TRIGGERTYPE_REDUNDANT_TRANSMISSION_CHANGE                    TriggerType = "REDUNDANT_TRANSMISSION_CHANGE"
	TRIGGERTYPE_CGI_SAI_CHANGE                                   TriggerType = "CGI_SAI_CHANGE"
	TRIGGERTYPE_RAI_CHANGE                                       TriggerType = "RAI_CHANGE"
	TRIGGERTYPE_VSMF_CHANGE                                      TriggerType = "VSMF_CHANGE"
)
