/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyControlRequestTrigger string

// Define constant values for PolicyControlRequestTrigger
const (
	POLICYCONTROLREQUESTTRIGGER_PLMN_CH                          PolicyControlRequestTrigger = "PLMN_CH"
	POLICYCONTROLREQUESTTRIGGER_RES_MO_RE                        PolicyControlRequestTrigger = "RES_MO_RE"
	POLICYCONTROLREQUESTTRIGGER_AC_TY_CH                         PolicyControlRequestTrigger = "AC_TY_CH"
	POLICYCONTROLREQUESTTRIGGER_UE_IP_CH                         PolicyControlRequestTrigger = "UE_IP_CH"
	POLICYCONTROLREQUESTTRIGGER_UE_MAC_CH                        PolicyControlRequestTrigger = "UE_MAC_CH"
	POLICYCONTROLREQUESTTRIGGER_AN_CH_COR                        PolicyControlRequestTrigger = "AN_CH_COR"
	POLICYCONTROLREQUESTTRIGGER_US_RE                            PolicyControlRequestTrigger = "US_RE"
	POLICYCONTROLREQUESTTRIGGER_APP_STA                          PolicyControlRequestTrigger = "APP_STA"
	POLICYCONTROLREQUESTTRIGGER_APP_STO                          PolicyControlRequestTrigger = "APP_STO"
	POLICYCONTROLREQUESTTRIGGER_AN_INFO                          PolicyControlRequestTrigger = "AN_INFO"
	POLICYCONTROLREQUESTTRIGGER_CM_SES_FAIL                      PolicyControlRequestTrigger = "CM_SES_FAIL"
	POLICYCONTROLREQUESTTRIGGER_PS_DA_OFF                        PolicyControlRequestTrigger = "PS_DA_OFF"
	POLICYCONTROLREQUESTTRIGGER_DEF_QOS_CH                       PolicyControlRequestTrigger = "DEF_QOS_CH"
	POLICYCONTROLREQUESTTRIGGER_SE_AMBR_CH                       PolicyControlRequestTrigger = "SE_AMBR_CH"
	POLICYCONTROLREQUESTTRIGGER_QOS_NOTIF                        PolicyControlRequestTrigger = "QOS_NOTIF"
	POLICYCONTROLREQUESTTRIGGER_NO_CREDIT                        PolicyControlRequestTrigger = "NO_CREDIT"
	POLICYCONTROLREQUESTTRIGGER_REALLO_OF_CREDIT                 PolicyControlRequestTrigger = "REALLO_OF_CREDIT"
	POLICYCONTROLREQUESTTRIGGER_PRA_CH                           PolicyControlRequestTrigger = "PRA_CH"
	POLICYCONTROLREQUESTTRIGGER_SAREA_CH                         PolicyControlRequestTrigger = "SAREA_CH"
	POLICYCONTROLREQUESTTRIGGER_SCNN_CH                          PolicyControlRequestTrigger = "SCNN_CH"
	POLICYCONTROLREQUESTTRIGGER_RE_TIMEOUT                       PolicyControlRequestTrigger = "RE_TIMEOUT"
	POLICYCONTROLREQUESTTRIGGER_RES_RELEASE                      PolicyControlRequestTrigger = "RES_RELEASE"
	POLICYCONTROLREQUESTTRIGGER_SUCC_RES_ALLO                    PolicyControlRequestTrigger = "SUCC_RES_ALLO"
	POLICYCONTROLREQUESTTRIGGER_RAI_CH                           PolicyControlRequestTrigger = "RAI_CH"
	POLICYCONTROLREQUESTTRIGGER_RAT_TY_CH                        PolicyControlRequestTrigger = "RAT_TY_CH"
	POLICYCONTROLREQUESTTRIGGER_REF_QOS_IND_CH                   PolicyControlRequestTrigger = "REF_QOS_IND_CH"
	POLICYCONTROLREQUESTTRIGGER_NUM_OF_PACKET_FILTER             PolicyControlRequestTrigger = "NUM_OF_PACKET_FILTER"
	POLICYCONTROLREQUESTTRIGGER_UE_STATUS_RESUME                 PolicyControlRequestTrigger = "UE_STATUS_RESUME"
	POLICYCONTROLREQUESTTRIGGER_UE_TZ_CH                         PolicyControlRequestTrigger = "UE_TZ_CH"
	POLICYCONTROLREQUESTTRIGGER_AUTH_PROF_CH                     PolicyControlRequestTrigger = "AUTH_PROF_CH"
	POLICYCONTROLREQUESTTRIGGER_QOS_MONITORING                   PolicyControlRequestTrigger = "QOS_MONITORING"
	POLICYCONTROLREQUESTTRIGGER_SCELL_CH                         PolicyControlRequestTrigger = "SCELL_CH"
	POLICYCONTROLREQUESTTRIGGER_USER_LOCATION_CH                 PolicyControlRequestTrigger = "USER_LOCATION_CH"
	POLICYCONTROLREQUESTTRIGGER_EPS_FALLBACK                     PolicyControlRequestTrigger = "EPS_FALLBACK"
	POLICYCONTROLREQUESTTRIGGER_MA_PDU                           PolicyControlRequestTrigger = "MA_PDU"
	POLICYCONTROLREQUESTTRIGGER_TSN_BRIDGE_INFO                  PolicyControlRequestTrigger = "TSN_BRIDGE_INFO"
	POLICYCONTROLREQUESTTRIGGER_5G_RG_JOIN                       PolicyControlRequestTrigger = "5G_RG_JOIN"
	POLICYCONTROLREQUESTTRIGGER_5G_RG_LEAVE                      PolicyControlRequestTrigger = "5G_RG_LEAVE"
	POLICYCONTROLREQUESTTRIGGER_DDN_FAILURE                      PolicyControlRequestTrigger = "DDN_FAILURE"
	POLICYCONTROLREQUESTTRIGGER_DDN_DELIVERY_STATUS              PolicyControlRequestTrigger = "DDN_DELIVERY_STATUS"
	POLICYCONTROLREQUESTTRIGGER_GROUP_ID_LIST_CHG                PolicyControlRequestTrigger = "GROUP_ID_LIST_CHG"
	POLICYCONTROLREQUESTTRIGGER_DDN_FAILURE_CANCELLATION         PolicyControlRequestTrigger = "DDN_FAILURE_CANCELLATION"
	POLICYCONTROLREQUESTTRIGGER_DDN_DELIVERY_STATUS_CANCELLATION PolicyControlRequestTrigger = "DDN_DELIVERY_STATUS_CANCELLATION"
	POLICYCONTROLREQUESTTRIGGER_VPLMN_QOS_CH                     PolicyControlRequestTrigger = "VPLMN_QOS_CH"
	POLICYCONTROLREQUESTTRIGGER_SUCC_QOS_UPDATE                  PolicyControlRequestTrigger = "SUCC_QOS_UPDATE"
	POLICYCONTROLREQUESTTRIGGER_SAT_CATEGORY_CHG                 PolicyControlRequestTrigger = "SAT_CATEGORY_CHG"
	POLICYCONTROLREQUESTTRIGGER_PCF_UE_NOTIF_IND                 PolicyControlRequestTrigger = "PCF_UE_NOTIF_IND"
	POLICYCONTROLREQUESTTRIGGER_NWDAF_DATA_CHG                   PolicyControlRequestTrigger = "NWDAF_DATA_CHG"
)
