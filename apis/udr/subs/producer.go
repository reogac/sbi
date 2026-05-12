/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package subs

import (
	"context"
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnCreateIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateIndividualAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateIndividualAuthenticationStatus(ctx.Context(), &params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryProvisionedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryProvisionedDataParams

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)), nil)
			return
		}
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryProvisionedData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmfRegList(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmfRegListParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuerySmfRegList(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleModifyIpSmGwContext(ctx.Context(), ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5mbsData(ctx sbi.RequestContext, prod Producer) {
	var params Query5mbsDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQuery5mbsData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryRoamingInformation(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryRoamingInformation(ctx.Context(), ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.IpSmGwRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateIpSmGwContext(ctx.Context(), ueId, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateCagUpdateAck(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateCagUpdateAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.CagAckData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateCagUpdateAck(ctx.Context(), &params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveSmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveSmfGroupSubscriptions(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AuthEvent)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationStatus(ctx.Context(), ueId, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetOdbData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetOdbData(ctx.Context(), ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryV2xData(ctx sbi.RequestContext, prod Producer) {
	var params QueryV2xDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp := prod.HandleQueryV2xData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryLcsBcaData(ctx sbi.RequestContext, prod Producer) {
	var params QueryLcsBcaDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryLcsBcaData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryPeiInformation(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryPeiInformation(ctx.Context(), ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfGroupSubscriptions(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDelete5GmbsGroup(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDelete5GmbsGroup(ctx.Context(), externalGroupId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryAuthSoR(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthSoR(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmfContext3gppParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryAmfContext3gpp(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmsData(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmsDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQuerySmsData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnGetHssSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetHssSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSubscriptionInfo(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdateEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateEeGroupSubscriptions(ctx.Context(), &params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetIdentityData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetIdentityDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'app-port-id'
	appPortIdStr := ctx.Param("app-port-id")
	if len(appPortIdStr) > 0 {
		if params.AppPortId, err = models.AppPortIdFromString(appPortIdStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse app-port-id failed: %+v", err)), nil)
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// call application handler
	headers, rsp, prob := prod.HandleGetIdentityData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQuery5GVnGroupInternal(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"), nil)
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroupInternal(ctx.Context(), internalGroupIds)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAmfContextNon3gpp(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetPPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var params GetPPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetPPDataEntry(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQuery5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuery5GVnGroup(ctx.Context(), gpsis)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyAuthenticationSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAuthenticationSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAuthenticationSubscription(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateOperSpecData(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveAmfSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveAmfSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveAmfSubscriptionsInfo(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryEEData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryEEDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryEEData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryLcsMoData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryLcsMoDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryLcsMoData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.ServiceSpecificAuthorizationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateServiceSpecificAuthorizationInfo(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAuthenticationStatusParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthenticationStatus(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySubsToNotify(ctx sbi.RequestContext, prod Producer) {
	var params QuerySubsToNotifyParams

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySubsToNotify(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetSSAuData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSSAuDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"), nil)
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
		return
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"), nil)
		return
	}

	// read 'mtc-provider-information'
	params.MtcProviderInformation = ctx.Param("mtc-provider-information")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// call application handler
	prob := prod.HandleGetSSAuData(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnGetAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params GetAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfGroupSubscriptions(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmsfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerySmsfContextNon3gpp(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryIpSmGwContext(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryIpSmGwContextParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryIpSmGwContext(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.MessageWaitingData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateMessageWaitingData(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetHssSDMSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetHssSDMSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetHssSDMSubscriptionInfo(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryUeSubscribedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryUeSubscribedDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) > 0 {
		if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)), nil)
			return
		}
	}

	// read 'serving-plmn'
	params.ServingPlmn = ctx.Param("serving-plmn")

	// call application handler
	rsp := prod.HandleQueryUeSubscribedData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnDeleteSmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContext3gpp(ctx.Context(), ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyHssSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyHssSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSubscriptionInfo(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyNiddAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyNiddAuthorizationInfo(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5GMbsGroupInternal(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'internal-group-ids'
	var internalGroupIds []string
	internalGroupIdsStr := ctx.Param("internal-group-ids")
	if len(internalGroupIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "internal-group-ids is required"), nil)
		return
	}

	if internalGroupIds, err = models.ArrayOfStringFromString(internalGroupIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse internal-group-ids failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupInternal(ctx.Context(), internalGroupIds)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.Amf3GppAccessRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateAmfContext3gpp(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params DeleteSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmfRegistration(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryCagAck(ctx sbi.RequestContext, prod Producer) {
	var params QueryCagAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryCagAck(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateEeSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateEeSubscriptions(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnRemoveSmfSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveSmfSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveSmfSubscriptionsInfo(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDelete5GVnGroup(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDelete5GVnGroup(ctx.Context(), externalGroupId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetNiddAuData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetNiddAuDataParams

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "single-nssai is required"), nil)
		return
	}

	if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
		return
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")
	if len(params.Dnn) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "dnn is required"), nil)
		return
	}

	// read 'mtc-provider-information'
	params.MtcProviderInformation = ctx.Param("mtc-provider-information")
	if len(params.MtcProviderInformation) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "mtc-provider-information is required"), nil)
		return
	}

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleGetNiddAuData(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnCreateAuthenticationUPU(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAuthenticationUPUParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.UpuData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationUPU(ctx.Context(), &params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnAmfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params AmfContext3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContext3gpp(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyEesubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyEesubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEesubscription(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreate5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.FiveGVnGroupConfiguration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GVnGroup(ctx.Context(), externalGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQueryPorseData(ctx sbi.RequestContext, prod Producer) {
	var params QueryPorseDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	headers, rsp := prod.HandleQueryPorseData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateOrUpdatePeiInformation(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PeiUpdateInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdatePeiInformation(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreatePPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreatePPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.PpDataEntry)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreatePPDataEntry(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnSubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SubscriptionDataSubscriptions)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleSubscriptionDataSubscriptions(ctx.Context(), body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnModify5GVnGroup(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Modify5GVnGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GVnGroup(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryUeLocation(ctx sbi.RequestContext, prod Producer) {
	var params QueryUeLocationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryUeLocation(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetNiddAuthorizationInfo(ctx.Context(), ueId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateSmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmsfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContextNon3gpp(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetAmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetAmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetAmfSubscriptionInfo(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetGroupIdentifiers(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetGroupIdentifiersParams

	// read 'int-group-id'
	params.IntGroupId = ctx.Param("int-group-id")

	// read 'ue-id-ind'
	ueIdIndStr := ctx.Param("ue-id-ind")
	if len(ueIdIndStr) > 0 {
		var ueIdIndTmp bool
		if ueIdIndTmp, err = models.BoolFromString(ueIdIndStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ue-id-ind failed: %+v", err)), nil)
			return
		}
		params.UeIdInd = &ueIdIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ext-group-id'
	params.ExtGroupId = ctx.Param("ext-group-id")

	// call application handler
	rsp, prob := prod.HandleGetGroupIdentifiers(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnQueryCoverageRestrictionData(ctx sbi.RequestContext, prod Producer) {
	var params QueryCoverageRestrictionDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryCoverageRestrictionData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnCreateAuthenticationSoR(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAuthenticationSoRParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SorData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateAuthenticationSoR(ctx.Context(), &params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmfRegistrationParams

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuerySmfRegistration(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnModifyPpData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyPpDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyPpData(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetSharedData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params GetSharedDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'shared-data-ids'
	sharedDataIdsStr := ctx.Param("shared-data-ids")
	if len(sharedDataIdsStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "shared-data-ids is required"), nil)
		return
	}

	if params.SharedDataIds, err = models.ArrayOfStringFromString(sharedDataIdsStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse shared-data-ids failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSharedData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateSmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmsfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmsfContext3gpp(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteSmsfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteSmsfContextNon3gpp(ctx.Context(), ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateHSSSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateHSSSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HssSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateHSSSubscriptions(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateAmfGroupSubscriptions(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteAuthenticationStatus(ctx.Context(), ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params AmfContextNon3gppParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleAmfContextNon3gpp(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdateSmfContext(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateSmfContextParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSmfContext(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeletePPDataEntry(ctx sbi.RequestContext, prod Producer) {
	var params DeletePPDataEntryParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'afInstanceId'
	params.AfInstanceId = ctx.Param("afInstanceId")
	if len(params.AfInstanceId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "afInstanceId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeletePPDataEntry(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModify5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Modify5GmbsGroupParams

	// read 'externalGroupId'
	params.ExternalGroupId = ctx.Param("externalGroupId")
	if len(params.ExternalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModify5GmbsGroup(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var params DeleteIndividualAuthenticationStatusParams

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteIndividualAuthenticationStatus(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemovesdmSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemovesdmSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleRemovesdmSubscriptions(ctx.Context(), &params)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveMultipleSubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params RemoveMultipleSubscriptionDataSubscriptionsParams

	// read 'ue-id'
	params.UeId = ctx.Param("ue-id")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ue-id is required"), nil)
		return
	}

	// read 'nf-instance-id'
	params.NfInstanceId = ctx.Param("nf-instance-id")

	// read 'delete-all-nfs'
	deleteAllNfsStr := ctx.Param("delete-all-nfs")
	if len(deleteAllNfsStr) > 0 {
		var deleteAllNfsTmp bool
		if deleteAllNfsTmp, err = models.BoolFromString(deleteAllNfsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse delete-all-nfs failed: %+v", err)), nil)
			return
		}
		params.DeleteAllNfs = &deleteAllNfsTmp
	}

	// read 'implicit-unsubscribe-indication'
	implicitUnsubscribeIndicationStr := ctx.Param("implicit-unsubscribe-indication")
	if len(implicitUnsubscribeIndicationStr) > 0 {
		var implicitUnsubscribeIndicationTmp bool
		if implicitUnsubscribeIndicationTmp, err = models.BoolFromString(implicitUnsubscribeIndicationStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse implicit-unsubscribe-indication failed: %+v", err)), nil)
			return
		}
		params.ImplicitUnsubscribeIndication = &implicitUnsubscribeIndicationTmp
	}

	// call application handler
	prod.HandleRemoveMultipleSubscriptionDataSubscriptions(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryAuthUPU(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthUPUParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryAuthUPU(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmsfContext3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmsfContext3gppParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuerySmsfContext3gpp(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnDeleteIpSmGwContext(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteIpSmGwContext(ctx.Context(), ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnUpdatesdmsubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdatesdmsubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SdmSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdatesdmsubscriptions(ctx.Context(), &params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryContextData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryContextDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'context-dataset-names'
	contextDatasetNamesStr := ctx.Param("context-dataset-names")
	if len(contextDatasetNamesStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "context-dataset-names is required"), nil)
		return
	}

	if params.ContextDatasetNames, err = models.ArrayOfStringFromString(contextDatasetNamesStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse context-dataset-names failed: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryContextData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnGetIndividualSharedData(ctx sbi.RequestContext, prod Producer) {
	var params GetIndividualSharedDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'sharedDataId'
	params.SharedDataId = ctx.Param("sharedDataId")
	if len(params.SharedDataId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "sharedDataId is required"), nil)
		return
	}

	// call application handler
	headers, rsp, prob := prod.HandleGetIndividualSharedData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnModifySmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfGroupSubscriptions(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryIndividualAuthenticationStatus(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryIndividualAuthenticationStatusParams

	// read 'servingNetworkName'
	params.ServingNetworkName = ctx.Param("servingNetworkName")
	if len(params.ServingNetworkName) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingNetworkName is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryIndividualAuthenticationStatus(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateOrUpdateSmfRegistration(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOrUpdateSmfRegistrationParams

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "pduSessionId is required"), nil)
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateOrUpdateSmfRegistration(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteMessageWaitingData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleDeleteMessageWaitingData(ctx.Context(), ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetMultiplePPDataEntries(ctx sbi.RequestContext, prod Producer) {
	var params GetMultiplePPDataEntriesParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetMultiplePPDataEntries(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnModifyAmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyAmfSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyAmfSubscriptionInfo(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryEeGroupSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QueryEeGroupSubscriptionParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQueryEeGroupSubscription(ctx.Context(), &params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnQuerySubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQuerySubscriptionDataSubscriptions(ctx.Context(), subsId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdateRoamingInformation(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.RoamingInfoUpdate)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleUpdateRoamingInformation(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateAMFSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateAMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateAMFSubscriptions(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params QueryEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryEeGroupSubscriptions(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueGroupId'
	var ueGroupId string
	ueGroupId = ctx.Param("ueGroupId")
	if len(ueGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateEeGroupSubscriptions(ctx.Context(), ueGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnGetSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params GetSmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfGroupSubscriptions(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuery5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'gpsis'
	var gpsis []string
	gpsisStr := ctx.Param("gpsis")
	if len(gpsisStr) > 0 {
		if gpsis, err = models.ArrayOfStringFromString(gpsisStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsis failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleQuery5GmbsGroup(ctx.Context(), gpsis)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnRemoveEeGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveEeGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveEeGroupSubscriptions(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerysdmsubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params QuerysdmsubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQuerysdmsubscriptions(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnUpdateAuthenticationSoR(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateAuthenticationSoRParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateAuthenticationSoR(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmDataParams

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.VarSnssaiFromString(singleNssaiStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	headers, rsp := prod.HandleQuerySmData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnModifySmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifySmfSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifySmfSubscriptionInfo(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryTraceData(ctx sbi.RequestContext, prod Producer) {
	var params QueryTraceDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQueryTraceData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQuery5GVNGroupPPData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Query5GVNGroupPPDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQuery5GVNGroupPPData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateNIDDAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NiddAuthorizationInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateNIDDAuthorizationInfo(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuerySmfSelectData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QuerySmfSelectDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	headers, rsp := prod.HandleQuerySmfSelectData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnQueryOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryOperSpecDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	headers, rsp := prod.HandleQueryOperSpecData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnModifyEeGroupSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyEeGroupSubscriptionParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyEeGroupSubscription(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifysdmSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysdmSubscription(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemovesubscriptionDataSubscriptions(ctx sbi.RequestContext, prod Producer) {

	// read 'subsId'
	var subsId string
	subsId = ctx.Param("subsId")
	if len(subsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemovesubscriptionDataSubscriptions(ctx.Context(), subsId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryLcsPrivacyData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryLcsPrivacyDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryLcsPrivacyData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnRemoveAmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveAmfGroupSubscriptionsParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveAmfGroupSubscriptions(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryPPData(ctx sbi.RequestContext, prod Producer) {
	var params QueryPPDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryPPData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryAuthSubsData(ctx sbi.RequestContext, prod Producer) {
	var params QueryAuthSubsDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleQueryAuthSubsData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryNssaiAck(ctx sbi.RequestContext, prod Producer) {
	var params QueryNssaiAckParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryNssaiAck(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryeesubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryeesubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'event-types'
	eventTypesStr := ctx.Param("event-types")
	if len(eventTypesStr) > 0 {
		if params.EventTypes, err = models.ArrayOfStringFromString(eventTypesStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse event-types failed: %+v", err)), nil)
			return
		}
	}

	// read 'nf-identifiers'
	nfIdentifiersStr := ctx.Param("nf-identifiers")
	if len(nfIdentifiersStr) > 0 {
		if params.NfIdentifiers, err = models.ArrayOfNfIdentifierFromString(nfIdentifiersStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse nf-identifiers failed: %+v", err)), nil)
			return
		}
	}

	// call application handler
	rsp := prod.HandleQueryeesubscriptions(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemoveServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveServiceSpecificAuthorizationInfo(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyServiceSpecificAuthorizationInfoParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyServiceSpecificAuthorizationInfo(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryeeSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QueryeeSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQueryeeSubscription(ctx.Context(), &params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnModifyHssSDMSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyHssSDMSubscriptionInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyHssSDMSubscriptionInfo(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryMessageWaitingDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryMessageWaitingData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnRemoveeeSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var params RemoveeeSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveeeSubscriptions(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateSmfGroupSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateSmfGroupSubscriptionsParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSmfGroupSubscriptions(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQuery5GMbsGroupPPData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params Query5GMbsGroupPPDataParams

	// read 'ext-group-ids'
	extGroupIdsStr := ctx.Param("ext-group-ids")
	if len(extGroupIdsStr) > 0 {
		if params.ExtGroupIds, err = models.ArrayOfStringFromString(extGroupIdsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ext-group-ids failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleQuery5GMbsGroupPPData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateOrUpdateNssaiAck(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateOrUpdateNssaiAckParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.NssaiAckData)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateOrUpdateNssaiAck(ctx.Context(), &params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryAmData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params QueryAmDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'fields'
	fieldsStr := ctx.Param("fields")
	if len(fieldsStr) > 0 {
		if params.Fields, err = models.ArrayOfStringFromString(fieldsStr); err != nil {
			ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse fields failed: %+v", err)), nil)
			return
		}
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQueryAmData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnGetppData(ctx sbi.RequestContext, prod Producer) {
	var params GetppDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp, prob := prod.HandleGetppData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnGetSmfSubscriptionInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetSmfSubscriptionInfoParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetSmfSubscriptionInfo(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateSMFSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateSMFSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SmfSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp := prod.HandleCreateSMFSubscriptions(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnQueryGroupEEData(ctx sbi.RequestContext, prod Producer) {
	var params QueryGroupEEDataParams

	// read 'ueGroupId'
	params.UeGroupId = ctx.Param("ueGroupId")
	if len(params.UeGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueGroupId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp := prod.HandleQueryGroupEEData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnCreateSdmSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.SdmSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateSdmSubscriptions(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

}

func OnQuerysdmSubscription(ctx sbi.RequestContext, prod Producer) {
	var params QuerysdmSubscriptionParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleQuerysdmSubscription(ctx.Context(), &params)

	// success
	ctx.WriteResponse(200, nil, nil)

}

func OnGet5GVnGroupConfiguration(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGet5GVnGroupConfiguration(ctx.Context(), externalGroupId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQuerySmsMngData(ctx sbi.RequestContext, prod Producer) {
	var params QuerySmsMngDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'servingPlmnId'
	params.ServingPlmnId = ctx.Param("servingPlmnId")
	if len(params.ServingPlmnId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "servingPlmnId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// call application handler
	headers, rsp := prod.HandleQuerySmsMngData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnUpdateEesubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params UpdateEesubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.EeSubscription)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleUpdateEesubscriptions(ctx.Context(), &params, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreateHSSSDMSubscriptions(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params CreateHSSSDMSubscriptionsParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.HssSubscriptionInfo)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prod.HandleCreateHSSSDMSubscriptions(ctx.Context(), &params, body)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveNiddAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveNiddAuthorizationInfo(ctx.Context(), ueId)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnGetServiceSpecificAuthorizationInfo(ctx sbi.RequestContext, prod Producer) {
	var params GetServiceSpecificAuthorizationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'serviceType'
	params.ServiceType = ctx.Param("serviceType")
	if len(params.ServiceType) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "serviceType is required"), nil)
		return
	}

	// call application handler
	rsp := prod.HandleGetServiceSpecificAuthorizationInfo(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

}

func OnQueryUserConsentData(ctx sbi.RequestContext, prod Producer) {
	var params QueryUserConsentDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ucPurpose'
	params.UcPurpose = ctx.Param("ucPurpose")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Header("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Header("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleQueryUserConsentData(ctx.Context(), &params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, headers)
		return
	}

}

func OnGetMulticastMbsGroupMemb(ctx sbi.RequestContext, prod Producer) {

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetMulticastMbsGroupMemb(ctx.Context(), externalGroupId)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

func OnCreateAmfContextNon3gpp(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.AmfNon3GppAccessRegistration)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	headers, rsp := prod.HandleCreateAmfContextNon3gpp(ctx.Context(), ueId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, headers)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnDeleteOperSpecData(ctx sbi.RequestContext, prod Producer) {

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// call application handler
	prob := prod.HandleDeleteOperSpecData(ctx.Context(), ueId)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyOperSpecData(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifyOperSpecDataParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifyOperSpecData(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifyMessageWaitingData(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	prob := prod.HandleModifyMessageWaitingData(ctx.Context(), ueId, body)

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveHssSDMSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveHssSDMSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveHssSDMSubscriptionsInfo(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnRemoveHssSubscriptionsInfo(ctx sbi.RequestContext, prod Producer) {
	var params RemoveHssSubscriptionsInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "ueId is required"), nil)
		return
	}

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// call application handler
	prod.HandleRemoveHssSubscriptionsInfo(ctx.Context(), &params)

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnModifysubscriptionDataSubscription(ctx sbi.RequestContext, prod Producer) {
	var err error
	var params ModifysubscriptionDataSubscriptionParams

	// read 'subsId'
	params.SubsId = ctx.Param("subsId")
	if len(params.SubsId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "subsId is required"), nil)
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// decode request body
	contentLength, content := ctx.RequestBody()
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleModifysubscriptionDataSubscription(ctx.Context(), &params, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

	// success
	ctx.WriteResponse(204, nil, nil)

}

func OnCreate5GmbsGroup(ctx sbi.RequestContext, prod Producer) {
	var err error

	// read 'externalGroupId'
	var externalGroupId string
	externalGroupId = ctx.Param("externalGroupId")
	if len(externalGroupId) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "externalGroupId is required"), nil)
		return
	}

	// decode request body
	contentLength, content := ctx.RequestBody()
	body := new(models.MulticastMbsGroupMemb)
	if err = sbi.Decode(contentLength, content, body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)), nil)
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreate5GmbsGroup(ctx.Context(), externalGroupId, body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp, nil)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob, nil)
		return
	}

}

type Producer interface {
	HandleCreateIndividualAuthenticationStatus(context.Context, *CreateIndividualAuthenticationStatusParams, *models.AuthEvent)

	HandleQueryProvisionedData(context.Context, *QueryProvisionedDataParams) *models.ProvisionedDataSets

	HandleQuerySmfRegList(context.Context, *QuerySmfRegListParams) *[]models.SmfRegistration

	HandleModifyIpSmGwContext(context.Context, string, *[]models.PatchItem) *models.ProblemDetails

	HandleQuery5mbsData(context.Context, *Query5mbsDataParams) (map[string]string, *models.MbsSubscriptionData)

	HandleQueryRoamingInformation(context.Context, string) *models.RoamingInfoUpdate

	HandleCreateIpSmGwContext(context.Context, string, *models.IpSmGwRegistration)

	HandleCreateCagUpdateAck(context.Context, *CreateCagUpdateAckParams, *models.CagAckData)

	HandleRemoveSmfGroupSubscriptions(context.Context, *RemoveSmfGroupSubscriptionsParams)

	HandleCreateAuthenticationStatus(context.Context, string, *models.AuthEvent)

	HandleGetOdbData(context.Context, string) *models.OdbData

	HandleQueryV2xData(context.Context, *QueryV2xDataParams) (map[string]string, *models.V2xSubscriptionData)

	HandleQueryLcsBcaData(context.Context, *QueryLcsBcaDataParams) (map[string]string, *models.LcsBroadcastAssistanceTypesData)

	HandleQueryPeiInformation(context.Context, string) *models.PeiUpdateInfo

	HandleModifyAmfGroupSubscriptions(context.Context, *ModifyAmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleDelete5GmbsGroup(context.Context, string) *models.ProblemDetails

	HandleQueryAuthSoR(context.Context, *QueryAuthSoRParams) *models.SorData

	HandleQueryAmfContext3gpp(context.Context, *QueryAmfContext3gppParams) *models.Amf3GppAccessRegistration

	HandleQuerySmsData(context.Context, *QuerySmsDataParams) (map[string]string, *models.SmsSubscriptionData)

	HandleGetHssSubscriptionInfo(context.Context, *GetHssSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleUpdateEeGroupSubscriptions(context.Context, *UpdateEeGroupSubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleGetIdentityData(context.Context, *GetIdentityDataParams) (map[string]string, *models.IdentityData, *models.ProblemDetails)

	HandleQuery5GVnGroupInternal(context.Context, []string) *map[string]models.FiveGVnGroupConfiguration

	HandleQueryAmfContextNon3gpp(context.Context, *QueryAmfContextNon3gppParams) *models.AmfNon3GppAccessRegistration

	HandleGetPPDataEntry(context.Context, *GetPPDataEntryParams) (*models.PpDataEntry, *models.ProblemDetails)

	HandleQuery5GVnGroup(context.Context, []string) *map[string]models.FiveGVnGroupConfiguration

	HandleModifyAuthenticationSubscription(context.Context, *ModifyAuthenticationSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreateOperSpecData(context.Context, *CreateOperSpecDataParams, *map[string]models.OperatorSpecificDataContainer) (*map[string]models.OperatorSpecificDataContainer, *models.ProblemDetails)

	HandleRemoveAmfSubscriptionsInfo(context.Context, *RemoveAmfSubscriptionsInfoParams)

	HandleQueryEEData(context.Context, *QueryEEDataParams) *models.EeProfileData

	HandleQueryLcsMoData(context.Context, *QueryLcsMoDataParams) (map[string]string, *models.LcsMoData)

	HandleCreateServiceSpecificAuthorizationInfo(context.Context, *CreateServiceSpecificAuthorizationInfoParams, *models.ServiceSpecificAuthorizationInfo) *models.ServiceSpecificAuthorizationInfo

	HandleQueryAuthenticationStatus(context.Context, *QueryAuthenticationStatusParams) *models.AuthEvent

	HandleQuerySubsToNotify(context.Context, *QuerySubsToNotifyParams) *[]models.SubscriptionDataSubscriptions

	HandleGetSSAuData(context.Context, *GetSSAuDataParams) *models.ProblemDetails

	HandleGetAmfGroupSubscriptions(context.Context, *GetAmfGroupSubscriptionsParams) *[]models.AmfSubscriptionInfo

	HandleQuerySmsfContextNon3gpp(context.Context, *QuerySmsfContextNon3gppParams) *models.SmsfRegistration

	HandleQueryIpSmGwContext(context.Context, *QueryIpSmGwContextParams) *models.IpSmGwRegistration

	HandleCreateMessageWaitingData(context.Context, string, *models.MessageWaitingData) *models.MessageWaitingData

	HandleGetHssSDMSubscriptionInfo(context.Context, *GetHssSDMSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleQueryUeSubscribedData(context.Context, *QueryUeSubscribedDataParams) *models.UeSubscribedDataSets

	HandleDeleteSmsfContext3gpp(context.Context, string)

	HandleModifyHssSubscriptionInfo(context.Context, *ModifyHssSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifyNiddAuthorizationInfo(context.Context, *ModifyNiddAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQuery5GMbsGroupInternal(context.Context, []string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleCreateAmfContext3gpp(context.Context, string, *models.Amf3GppAccessRegistration) (map[string]string, *models.Amf3GppAccessRegistration)

	HandleDeleteSmfRegistration(context.Context, *DeleteSmfRegistrationParams)

	HandleQueryCagAck(context.Context, *QueryCagAckParams) *models.CagAckData

	HandleCreateEeSubscriptions(context.Context, string, *models.EeSubscription) (map[string]string, *models.EeSubscription)

	HandleRemoveSmfSubscriptionsInfo(context.Context, *RemoveSmfSubscriptionsInfoParams)

	HandleDelete5GVnGroup(context.Context, string)

	HandleGetNiddAuData(context.Context, *GetNiddAuDataParams) *models.ProblemDetails

	HandleCreateAuthenticationUPU(context.Context, *CreateAuthenticationUPUParams, *models.UpuData)

	HandleAmfContext3gpp(context.Context, *AmfContext3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifyEesubscription(context.Context, *ModifyEesubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleCreate5GVnGroup(context.Context, string, *models.FiveGVnGroupConfiguration) (*models.FiveGVnGroupConfiguration, *models.ProblemDetails)

	HandleQueryPorseData(context.Context, *QueryPorseDataParams) (map[string]string, *models.ProseSubscriptionData)

	HandleCreateOrUpdatePeiInformation(context.Context, string, *models.PeiUpdateInfo) *models.PeiUpdateInfo

	HandleCreatePPDataEntry(context.Context, *CreatePPDataEntryParams, *models.PpDataEntry) (*models.PpDataEntry, *models.ProblemDetails)

	HandleSubscriptionDataSubscriptions(context.Context, *models.SubscriptionDataSubscriptions) (map[string]string, *models.SubscriptionDataSubscriptions)

	HandleModify5GVnGroup(context.Context, *Modify5GVnGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryUeLocation(context.Context, *QueryUeLocationParams) *models.LocationInfo

	HandleGetNiddAuthorizationInfo(context.Context, string) *models.NiddAuthorizationInfo

	HandleCreateSmsfContextNon3gpp(context.Context, string, *models.SmsfRegistration) *models.SmsfRegistration

	HandleGetAmfSubscriptionInfo(context.Context, *GetAmfSubscriptionInfoParams) *[]models.AmfSubscriptionInfo

	HandleGetGroupIdentifiers(context.Context, *GetGroupIdentifiersParams) (*models.GroupIdentifiers, *models.ProblemDetails)

	HandleQueryCoverageRestrictionData(context.Context, *QueryCoverageRestrictionDataParams) (map[string]string, *models.EnhancedCoverageRestrictionData)

	HandleCreateAuthenticationSoR(context.Context, *CreateAuthenticationSoRParams, *models.SorData)

	HandleQuerySmfRegistration(context.Context, *QuerySmfRegistrationParams) *models.SmfRegistration

	HandleModifyPpData(context.Context, *ModifyPpDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleGetSharedData(context.Context, *GetSharedDataParams) (*[]models.SharedData, *models.ProblemDetails)

	HandleCreateSmsfContext3gpp(context.Context, string, *models.SmsfRegistration) *models.SmsfRegistration

	HandleDeleteSmsfContextNon3gpp(context.Context, string)

	HandleCreateHSSSubscriptions(context.Context, *CreateHSSSubscriptionsParams, *models.HssSubscriptionInfo) *models.HssSubscriptionInfo

	HandleCreateAmfGroupSubscriptions(context.Context, *CreateAmfGroupSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleDeleteAuthenticationStatus(context.Context, string)

	HandleAmfContextNon3gpp(context.Context, *AmfContextNon3gppParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleUpdateSmfContext(context.Context, *UpdateSmfContextParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleDeletePPDataEntry(context.Context, *DeletePPDataEntryParams) *models.ProblemDetails

	HandleModify5GmbsGroup(context.Context, *Modify5GmbsGroupParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleDeleteIndividualAuthenticationStatus(context.Context, *DeleteIndividualAuthenticationStatusParams)

	HandleRemovesdmSubscriptions(context.Context, *RemovesdmSubscriptionsParams) *models.ProblemDetails

	HandleRemoveMultipleSubscriptionDataSubscriptions(context.Context, *RemoveMultipleSubscriptionDataSubscriptionsParams)

	HandleQueryAuthUPU(context.Context, *QueryAuthUPUParams) *models.UpuData

	HandleQuerySmsfContext3gpp(context.Context, *QuerySmsfContext3gppParams) *models.SmsfRegistration

	HandleDeleteIpSmGwContext(context.Context, string)

	HandleUpdatesdmsubscriptions(context.Context, *UpdatesdmsubscriptionsParams, *models.SdmSubscription) *models.ProblemDetails

	HandleQueryContextData(context.Context, *QueryContextDataParams) *models.ContextDataSets

	HandleGetIndividualSharedData(context.Context, *GetIndividualSharedDataParams) (map[string]string, *models.SharedData, *models.ProblemDetails)

	HandleModifySmfGroupSubscriptions(context.Context, *ModifySmfGroupSubscriptionsParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryIndividualAuthenticationStatus(context.Context, *QueryIndividualAuthenticationStatusParams) *models.AuthEvent

	HandleCreateOrUpdateSmfRegistration(context.Context, *CreateOrUpdateSmfRegistrationParams, *models.SmfRegistration) *models.SmfRegistration

	HandleDeleteMessageWaitingData(context.Context, string)

	HandleGetMultiplePPDataEntries(context.Context, *GetMultiplePPDataEntriesParams) (*models.PpDataEntryList, *models.ProblemDetails)

	HandleModifyAmfSubscriptionInfo(context.Context, *ModifyAmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryEeGroupSubscription(context.Context, *QueryEeGroupSubscriptionParams)

	HandleQuerySubscriptionDataSubscriptions(context.Context, string) *models.SubscriptionDataSubscriptions

	HandleUpdateRoamingInformation(context.Context, string, *models.RoamingInfoUpdate) *models.RoamingInfoUpdate

	HandleCreateAMFSubscriptions(context.Context, *CreateAMFSubscriptionsParams, *[]models.AmfSubscriptionInfo) *[]models.AmfSubscriptionInfo

	HandleQueryEeGroupSubscriptions(context.Context, *QueryEeGroupSubscriptionsParams) *[]models.EeSubscription

	HandleCreateEeGroupSubscriptions(context.Context, string, *models.EeSubscription) (map[string]string, *models.EeSubscription)

	HandleGetSmfGroupSubscriptions(context.Context, *GetSmfGroupSubscriptionsParams) *models.SmfSubscriptionInfo

	HandleQuery5GmbsGroup(context.Context, []string) (*map[string]models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleRemoveEeGroupSubscriptions(context.Context, *RemoveEeGroupSubscriptionsParams)

	HandleQuerysdmsubscriptions(context.Context, *QuerysdmsubscriptionsParams) *[]models.SdmSubscription

	HandleUpdateAuthenticationSoR(context.Context, *UpdateAuthenticationSoRParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQuerySmData(context.Context, *QuerySmDataParams) (map[string]string, *models.SmSubsData)

	HandleModifySmfSubscriptionInfo(context.Context, *ModifySmfSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryTraceData(context.Context, *QueryTraceDataParams) (map[string]string, *models.TraceData)

	HandleQuery5GVNGroupPPData(context.Context, *Query5GVNGroupPPDataParams) *models.Pp5gVnGroupProfileData

	HandleCreateNIDDAuthorizationInfo(context.Context, string, *models.NiddAuthorizationInfo) *models.NiddAuthorizationInfo

	HandleQuerySmfSelectData(context.Context, *QuerySmfSelectDataParams) (map[string]string, *models.SmfSelectionSubscriptionData)

	HandleQueryOperSpecData(context.Context, *QueryOperSpecDataParams) (map[string]string, *map[string]models.OperatorSpecificDataContainer)

	HandleModifyEeGroupSubscription(context.Context, *ModifyEeGroupSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifysdmSubscription(context.Context, *ModifysdmSubscriptionParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleRemovesubscriptionDataSubscriptions(context.Context, string)

	HandleQueryLcsPrivacyData(context.Context, *QueryLcsPrivacyDataParams) (map[string]string, *models.LcsPrivacyData)

	HandleRemoveAmfGroupSubscriptions(context.Context, *RemoveAmfGroupSubscriptionsParams)

	HandleQueryPPData(context.Context, *QueryPPDataParams) *models.PpProfileData

	HandleQueryAuthSubsData(context.Context, *QueryAuthSubsDataParams) *models.AuthenticationSubscription

	HandleQueryNssaiAck(context.Context, *QueryNssaiAckParams) *models.NssaiAckData

	HandleQueryeesubscriptions(context.Context, *QueryeesubscriptionsParams) *[]models.EeSubscriptionExt

	HandleRemoveServiceSpecificAuthorizationInfo(context.Context, *RemoveServiceSpecificAuthorizationInfoParams)

	HandleModifyServiceSpecificAuthorizationInfo(context.Context, *ModifyServiceSpecificAuthorizationInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryeeSubscription(context.Context, *QueryeeSubscriptionParams)

	HandleModifyHssSDMSubscriptionInfo(context.Context, *ModifyHssSDMSubscriptionInfoParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleQueryMessageWaitingData(context.Context, *QueryMessageWaitingDataParams) *models.MessageWaitingData

	HandleRemoveeeSubscriptions(context.Context, *RemoveeeSubscriptionsParams)

	HandleCreateSmfGroupSubscriptions(context.Context, *CreateSmfGroupSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleQuery5GMbsGroupPPData(context.Context, *Query5GMbsGroupPPDataParams) (*models.Pp5gMbsGroupProfileData, *models.ProblemDetails)

	HandleCreateOrUpdateNssaiAck(context.Context, *CreateOrUpdateNssaiAckParams, *models.NssaiAckData)

	HandleQueryAmData(context.Context, *QueryAmDataParams) (map[string]string, *models.AccessAndMobilitySubscriptionData)

	HandleGetppData(context.Context, *GetppDataParams) (map[string]string, *models.PpData, *models.ProblemDetails)

	HandleGetSmfSubscriptionInfo(context.Context, *GetSmfSubscriptionInfoParams) *models.SmfSubscriptionInfo

	HandleCreateSMFSubscriptions(context.Context, *CreateSMFSubscriptionsParams, *models.SmfSubscriptionInfo) *models.SmfSubscriptionInfo

	HandleQueryGroupEEData(context.Context, *QueryGroupEEDataParams) *models.EeGroupProfileData

	HandleCreateSdmSubscriptions(context.Context, string, *models.SdmSubscription) (map[string]string, *models.SdmSubscription)

	HandleQuerysdmSubscription(context.Context, *QuerysdmSubscriptionParams)

	HandleGet5GVnGroupConfiguration(context.Context, string) *models.FiveGVnGroupConfiguration

	HandleQuerySmsMngData(context.Context, *QuerySmsMngDataParams) (map[string]string, *models.SmsManagementSubscriptionData)

	HandleUpdateEesubscriptions(context.Context, *UpdateEesubscriptionsParams, *models.EeSubscription) *models.ProblemDetails

	HandleCreateHSSSDMSubscriptions(context.Context, *CreateHSSSDMSubscriptionsParams, *models.HssSubscriptionInfo)

	HandleRemoveNiddAuthorizationInfo(context.Context, string)

	HandleGetServiceSpecificAuthorizationInfo(context.Context, *GetServiceSpecificAuthorizationInfoParams) *models.ServiceSpecificAuthorizationInfo

	HandleQueryUserConsentData(context.Context, *QueryUserConsentDataParams) (map[string]string, *models.UcSubscriptionData)

	HandleGetMulticastMbsGroupMemb(context.Context, string) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)

	HandleCreateAmfContextNon3gpp(context.Context, string, *models.AmfNon3GppAccessRegistration) (map[string]string, *models.Amf3GppAccessRegistration)

	HandleDeleteOperSpecData(context.Context, string) *models.ProblemDetails

	HandleModifyOperSpecData(context.Context, *ModifyOperSpecDataParams, *[]models.PatchItem) (*models.PatchResult, *models.ProblemDetails)

	HandleModifyMessageWaitingData(context.Context, string, *[]models.PatchItem) *models.ProblemDetails

	HandleRemoveHssSDMSubscriptionsInfo(context.Context, *RemoveHssSDMSubscriptionsInfoParams)

	HandleRemoveHssSubscriptionsInfo(context.Context, *RemoveHssSubscriptionsInfoParams)

	HandleModifysubscriptionDataSubscription(context.Context, *ModifysubscriptionDataSubscriptionParams, *[]models.PatchItem) (*models.Schema, *models.ProblemDetails)

	HandleCreate5GmbsGroup(context.Context, string, *models.MulticastMbsGroupMemb) (*models.MulticastMbsGroupMemb, *models.ProblemDetails)
}
