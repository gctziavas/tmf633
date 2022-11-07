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

// ServiceCategoryApiController binds http requests to an api service and writes the service results to the http response
type ServiceCategoryApiController struct {
	service ServiceCategoryApiServicer
	errorHandler ErrorHandler
}

// ServiceCategoryApiOption for how the controller is set up.
type ServiceCategoryApiOption func(*ServiceCategoryApiController)

// WithServiceCategoryApiErrorHandler inject ErrorHandler into controller
func WithServiceCategoryApiErrorHandler(h ErrorHandler) ServiceCategoryApiOption {
	return func(c *ServiceCategoryApiController) {
		c.errorHandler = h
	}
}

// NewServiceCategoryApiController creates a default api controller
func NewServiceCategoryApiController(s ServiceCategoryApiServicer, opts ...ServiceCategoryApiOption) Router {
	controller := &ServiceCategoryApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ServiceCategoryApiController
func (c *ServiceCategoryApiController) Routes() Routes {
	return Routes{ 
		{
			"CreateServiceCategory",
			strings.ToUpper("Post"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCategory",
			c.CreateServiceCategory,
		},
		{
			"DeleteServiceCategory",
			strings.ToUpper("Delete"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCategory/{id}",
			c.DeleteServiceCategory,
		},
		{
			"ListServiceCategory",
			strings.ToUpper("Get"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCategory",
			c.ListServiceCategory,
		},
		{
			"PatchServiceCategory",
			strings.ToUpper("Patch"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCategory/{id}",
			c.PatchServiceCategory,
		},
		{
			"RetrieveServiceCategory",
			strings.ToUpper("Get"),
			"/tmf-api/serviceCatalogManagement/v4/serviceCategory/{id}",
			c.RetrieveServiceCategory,
		},
	}
}

// CreateServiceCategory - Creates a ServiceCategory
func (c *ServiceCategoryApiController) CreateServiceCategory(w http.ResponseWriter, r *http.Request) {
	serviceCategoryParam := ServiceCategoryCreate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&serviceCategoryParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertServiceCategoryCreateRequired(serviceCategoryParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateServiceCategory(r.Context(), serviceCategoryParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteServiceCategory - Deletes a ServiceCategory
func (c *ServiceCategoryApiController) DeleteServiceCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	
	result, err := c.service.DeleteServiceCategory(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ListServiceCategory - List or find ServiceCategory objects
func (c *ServiceCategoryApiController) ListServiceCategory(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListServiceCategory(r.Context(), fieldsParam, offsetParam, limitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PatchServiceCategory - Updates partially a ServiceCategory
func (c *ServiceCategoryApiController) PatchServiceCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	
	serviceCategoryParam := ServiceCategoryUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&serviceCategoryParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertServiceCategoryUpdateRequired(serviceCategoryParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PatchServiceCategory(r.Context(), idParam, serviceCategoryParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RetrieveServiceCategory - Retrieves a ServiceCategory by ID
func (c *ServiceCategoryApiController) RetrieveServiceCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	idParam := params["id"]
	
	fieldsParam := query.Get("fields")
	result, err := c.service.RetrieveServiceCategory(r.Context(), idParam, fieldsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}