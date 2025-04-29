/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Apr 29 09:38:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package conf

import (
	"fmt"
	"github.com/reogac/sbi"
	"github.com/reogac/sbi/models"
)

func OnGetUdmConfiguration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// call application handler
	rsp, prob := prod.HandleGetUdmConfiguration()

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetNssfConfiguration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// call application handler
	rsp, prob := prod.HandleGetNssfConfiguration()

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetSessionManagementConfiguration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error
	var params GetSessionManagementConfigurationParams

	// read 'uuid'
	params.Uuid = ctx.Param("uuid")
	if len(params.Uuid) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "uuid is required"))
		return
	}

	// read 'slice'
	sliceStr := ctx.Param("slice")
	if len(sliceStr) == 0 {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, "slice is required"))
		return
	}

	if params.Slice, err = models.SnssaiFromString(sliceStr); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("parse slice failed: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSessionManagementConfiguration(&params)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetUserPlaneConfiguration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.UserPlaneConfigurationRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		ctx.WriteResponse(400, models.CreateProblemDetails(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetUserPlaneConfiguration(body)

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

func OnGetUdrConfiguration(ctx sbi.RequestContext, handler any) {
	prod := handler.(Producer)

	// call application handler
	rsp, prob := prod.HandleGetUdrConfiguration()

	// check for success response
	if rsp != nil {
		ctx.WriteResponse(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		ctx.WriteResponse(prob.Status, prob)
		return
	}

}

type Producer interface {
	HandleGetUdmConfiguration() (*models.UdmConfiguration, *models.ProblemDetails)

	HandleGetNssfConfiguration() (*models.NssfConfiguration, *models.ProblemDetails)

	HandleGetSessionManagementConfiguration(*GetSessionManagementConfigurationParams) (*models.SessionManagementConfiguration, *models.ProblemDetails)

	HandleGetUserPlaneConfiguration(*models.UserPlaneConfigurationRequest) (*models.UserPlaneConfigurationResponse, *models.ProblemDetails)

	HandleGetUdrConfiguration() (*models.UdrConfiguration, *models.ProblemDetails)
}
