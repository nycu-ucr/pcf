/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smpolicy

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/http_wrapper"
	"github.com/nycu-ucr/openapi"
	"github.com/nycu-ucr/openapi/models"
	"github.com/nycu-ucr/pcf/logger"
	"github.com/nycu-ucr/pcf/producer"
)

// SmPoliciesPost -
func HTTPSmPoliciesPost(c *gin.Context) {
	var smPolicyContextData models.SmPolicyContextData
	// step 1: retrieve http request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.SMpolicylog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	// step 2: convert requestBody to openapi models
	err = openapi.Deserialize(&smPolicyContextData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.SMpolicylog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, smPolicyContextData)
	rsp := producer.HandleCreateSmPolicyRequest(req)

	// step 5: response
	for key, val := range rsp.Header { // header response is optional
		c.Header(key, val[0])
	}
	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.SMpolicylog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// SmPoliciesSmPolicyIdDeletePost -
func HTTPSmPoliciesSmPolicyIdDeletePost(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["smPolicyId"] = c.Params.ByName("smPolicyId")

	rsp := producer.HandleDeleteSmPolicyContextRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.SMpolicylog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// SmPoliciesSmPolicyIdGet -
func HTTPSmPoliciesSmPolicyIDGet(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["smPolicyId"] = c.Params.ByName("smPolicyId")

	rsp := producer.HandleGetSmPolicyContextRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.SMpolicylog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}

// SmPoliciesSmPolicyIdUpdatePost -
func HTTPSmPoliciesSmPolicyIdUpdatePost(c *gin.Context) {
	var smPolicyUpdateContextData models.SmPolicyUpdateContextData
	// step 1: retrieve http request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.SMpolicylog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	// step 2: convert requestBody to openapi models
	err = openapi.Deserialize(&smPolicyUpdateContextData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.SMpolicylog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, smPolicyUpdateContextData)
	req.Params["smPolicyId"] = c.Params.ByName("smPolicyId")

	rsp := producer.HandleUpdateSmPolicyContextRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.SMpolicylog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
