/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved. 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf633

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ExportJobApiController binds http requests to an api service and writes the service results to the http response
type ExportJobApiController struct {
	service ExportJobApiServicer
	errorHandler ErrorHandler
}

// ExportJobApiOption for how the controller is set up.
type ExportJobApiOption func(*ExportJobApiController)

// WithExportJobApiErrorHandler inject ErrorHandler into controller
func WithExportJobApiErrorHandler(h ErrorHandler) ExportJobApiOption {
	return func(c *ExportJobApiController) {
		c.errorHandler = h
	}
}

// NewExportJobApiController creates a default api controller
func NewExportJobApiController(s ExportJobApiServicer, opts ...ExportJobApiOption) Router {
	controller := &ExportJobApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ExportJobApiController
func (c *ExportJobApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateExportJob",
			strings.ToUpper("Post"),
			"/tmf-api/serviceCatalogManagement/v4/exportJob",
			c.CreateExportJob,
		},
		{
			"DeleteExportJob",
			strings.ToUpper("Delete"),
			"/tmf-api/serviceCatalogManagement/v4/exportJob/{id}",
			c.DeleteExportJob,
		},
		{
			"ListExportJob",
			strings.ToUpper("Get"),
			"/tmf-api/serviceCatalogManagement/v4/exportJob",
			c.ListExportJob,
		},
		{
			"RetrieveExportJob",
			strings.ToUpper("Get"),
			"/tmf-api/serviceCatalogManagement/v4/exportJob/{id}",
			c.RetrieveExportJob,
		},
	}
}

// CreateExportJob - Creates a ExportJob
func (c *ExportJobApiController) CreateExportJob(w http.ResponseWriter, r *http.Request) {
	exportJobParam := ExportJobCreate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&exportJobParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertExportJobCreateRequired(exportJobParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateExportJob(r.Context(), exportJobParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteExportJob - Deletes a ExportJob
func (c *ExportJobApiController) DeleteExportJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	
	result, err := c.service.DeleteExportJob(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListExportJob - List or find ExportJob objects
func (c *ExportJobApiController) ListExportJob(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fieldsParam := query.Get("fields")
	offsetParam, err := parseInt32Parameter(query.Get("offset"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ListExportJob(r.Context(), fieldsParam, offsetParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RetrieveExportJob - Retrieves a ExportJob by ID
func (c *ExportJobApiController) RetrieveExportJob(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	idParam := params["id"]
	
	fieldsParam := query.Get("fields")
	result, err := c.service.RetrieveExportJob(r.Context(), idParam, fieldsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
