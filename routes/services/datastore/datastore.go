// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package datastore contains the API implementation for managing the data stores
package datastore

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cuttle-ai/data-integration-service/config"
	"github.com/cuttle-ai/data-integration-service/routes"
	"github.com/cuttle-ai/data-integration-service/routes/response"
	"github.com/cuttle-ai/db-toolkit/datastores/services"
)

//GetServices api returns the list of data store services available in the system
func GetServices(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * We will get the app context
	 * We will get the list of services that are available
	 * We will return the response
	 */

	//getting the app context
	appCtx := ctx.Value(routes.AppContextKey).(*config.AppContext)
	appCtx.Log.Info("Got a request to fetch the list of datastore services", appCtx.Session.User.ID)

	//fetching the list of datastore services
	services, err := services.GetAll(appCtx.Db)
	if err != nil {
		//error while getting the list
		appCtx.Log.Error("error while getting the list", err.Error())
		response.WriteError(res, response.Error{Err: "Couldn't fetch the list"}, http.StatusInternalServerError)
		return
	}
	for i := 0; i < len(services); i++ {
		services[i].Password = ""
	}

	appCtx.Log.Info("Successfully fetched the list of datastore services of length", len(services))
	response.Write(res, response.Message{Message: "Successfully fetched the list", Data: services})
}

//UpdateService api creates a data store service in the system
func UpdateService(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * We will get the app context
	 * We will parse the request item
	 * We will validate the service
	 * We will update the service
	 * We will return the response
	 */

	//getting the app context
	appCtx := ctx.Value(routes.AppContextKey).(*config.AppContext)
	appCtx.Log.Info("Got a request to update a datastore service", appCtx.Session.User.ID)

	//parse the request param service
	s := &services.Service{}
	err := json.NewDecoder(req.Body).Decode(s)
	if err != nil {
		//bad request
		appCtx.Log.Error("error while parsing the datastore service", err.Error())
		response.WriteError(res, response.Error{Err: "Invalid Params " + err.Error()}, http.StatusBadRequest)
		return
	}

	//validating the service
	err = s.Validate()
	if err != nil {
		//error while validating the service
		appCtx.Log.Error("error while validating the datastore service", err.Error())
		response.WriteError(res, response.Error{Err: "Invalid Params " + err.Error()}, http.StatusBadRequest)
		return
	}

	//updating the service
	err = s.Update(appCtx.Db)
	if err != nil {
		//error while updating the service
		appCtx.Log.Error("error while updating the service", err.Error())
		response.WriteError(res, response.Error{Err: "Couldn't update the service"}, http.StatusInternalServerError)
		return
	}

	appCtx.Log.Info("Successfully updated the datastore service", s.ID)
	s.Password = ""
	response.Write(res, response.Message{Message: "Successfully updated the service", Data: s})
}

//CreateService api creates a data store service in the system
func CreateService(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * We will get the app context
	 * We will parse the request item
	 * We will validate the service
	 * We will create the service
	 * We will return the response
	 */

	//getting the app context
	appCtx := ctx.Value(routes.AppContextKey).(*config.AppContext)
	appCtx.Log.Info("Got a request to create a datastore service", appCtx.Session.User.ID)

	//parse the request param service
	s := &services.Service{}
	err := json.NewDecoder(req.Body).Decode(s)
	if err != nil {
		//bad request
		appCtx.Log.Error("error while parsing the datastore service", err.Error())
		response.WriteError(res, response.Error{Err: "Invalid Params " + err.Error()}, http.StatusBadRequest)
		return
	}

	//validating the service
	err = s.Validate()
	if err != nil {
		//error while validating the service
		appCtx.Log.Error("error while validating the datastore service", err.Error())
		response.WriteError(res, response.Error{Err: "Invalid Params " + err.Error()}, http.StatusBadRequest)
		return
	}

	//creating the service
	err = s.Create(appCtx.Db)
	if err != nil {
		//error while creating the service
		appCtx.Log.Error("error while creating the service", err.Error())
		response.WriteError(res, response.Error{Err: "Couldn't create the service"}, http.StatusInternalServerError)
		return
	}

	appCtx.Log.Info("Successfully created the datastore service", s.ID)
	s.Password = ""
	response.Write(res, response.Message{Message: "Successfully created the service", Data: s})
}

//DeleteService api deletes the data store service in the system
func DeleteService(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * We will get the app context
	 * We will parse the request item
	 * We will validate the service
	 * We will delete the service
	 * We will return the response
	 */

	//getting the app context
	appCtx := ctx.Value(routes.AppContextKey).(*config.AppContext)
	appCtx.Log.Info("Got a request to delete a datastore service", appCtx.Session.User.ID)

	//parse the request param service
	s := &services.Service{}
	err := json.NewDecoder(req.Body).Decode(s)
	if err != nil {
		//bad request
		appCtx.Log.Error("error while parsing the datastore service", err.Error())
		response.WriteError(res, response.Error{Err: "Invalid Params " + err.Error()}, http.StatusBadRequest)
		return
	}

	//validating the service
	if s.ID == 0 {
		//error while validating the service
		appCtx.Log.Error("error while validating the datastore service. Couldn't find the id of the service")
		response.WriteError(res, response.Error{Err: "Invalid Params. No ID found for the service."}, http.StatusBadRequest)
		return
	}

	//deleting the service
	err = s.Delete(appCtx.Db)
	if err != nil {
		//error while deleting the service
		appCtx.Log.Error("error while deleting the service", err.Error())
		response.WriteError(res, response.Error{Err: "Couldn't delete the service"}, http.StatusInternalServerError)
		return
	}

	appCtx.Log.Info("Successfully deleted the datastore service", s.ID)
	response.Write(res, response.Message{Message: "Successfully deleted the service"})
}

func init() {
	routes.AddRoutes(routes.Route{
		Version:     "v1",
		HandlerFunc: GetServices,
		Pattern:     "/services/datastore/list",
	})
	routes.AddRoutes(routes.Route{
		Version:     "v1",
		HandlerFunc: UpdateService,
		Pattern:     "/services/datastore/update",
	})
	routes.AddRoutes(routes.Route{
		Version:     "v1",
		HandlerFunc: CreateService,
		Pattern:     "/services/datastore/create",
	})
	routes.AddRoutes(routes.Route{
		Version:     "v1",
		HandlerFunc: DeleteService,
		Pattern:     "/services/datastore/delete",
	})
}
