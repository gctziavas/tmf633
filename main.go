/*
 * Service Catalog Management
 *
 * ## TMF API Reference: TMF633 - Service Catalog Management   Version 4.0   The Service Catalog API is one of Catalog Management API Family. Service Catalog API goal is to provide a catalog of services.  Service Catalog API performs the following operations on the resource : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity (for administration purposes) - Manage notification of events. .  Copyright © TM Forum 2020. All Rights Reserved.
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	tmf633 "github.com/gctziavas/tmf633/go"
)

func main() {
	log.Printf("Server started at http://localhost:8633")

	EventsSubscriptionApiService := tmf633.NewEventsSubscriptionApiService()
	EventsSubscriptionApiController := tmf633.NewEventsSubscriptionApiController(EventsSubscriptionApiService)

	ExportJobApiService := tmf633.NewExportJobApiService()
	ExportJobApiController := tmf633.NewExportJobApiController(ExportJobApiService)

	ImportJobApiService := tmf633.NewImportJobApiService()
	ImportJobApiController := tmf633.NewImportJobApiController(ImportJobApiService)

	NotificationListenersClientSideApiService := tmf633.NewNotificationListenersClientSideApiService()
	NotificationListenersClientSideApiController := tmf633.NewNotificationListenersClientSideApiController(NotificationListenersClientSideApiService)

	ServiceCandidateApiService := tmf633.NewServiceCandidateApiService()
	ServiceCandidateApiController := tmf633.NewServiceCandidateApiController(ServiceCandidateApiService)

	ServiceCatalogApiService := tmf633.NewServiceCatalogApiService()
	ServiceCatalogApiController := tmf633.NewServiceCatalogApiController(ServiceCatalogApiService)

	ServiceCategoryApiService := tmf633.NewServiceCategoryApiService()
	ServiceCategoryApiController := tmf633.NewServiceCategoryApiController(ServiceCategoryApiService)

	ServiceSpecificationApiService := tmf633.NewServiceSpecificationApiService()
	ServiceSpecificationApiController := tmf633.NewServiceSpecificationApiController(ServiceSpecificationApiService)

	router := tmf633.NewRouter(EventsSubscriptionApiController, ExportJobApiController, ImportJobApiController, NotificationListenersClientSideApiController, ServiceCandidateApiController, ServiceCatalogApiController, ServiceCategoryApiController, ServiceSpecificationApiController)

	log.Fatal(http.ListenAndServe(":8633", router))
}