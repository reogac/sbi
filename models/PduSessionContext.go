/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	Dnn                        string             `json:"dnn"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SNssai                     Snssai             `json:"sNssai"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
}
