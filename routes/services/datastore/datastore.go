// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package datastore contains the API implementation for managing the data stores
package datastore

import (
	"context"
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

	appCtx.Log.Info("Successfully fetched the list of datastore services of length", len(services))
	response.Write(res, response.Message{Message: "Successfully fetched the list", Data: services})
}

func init() {
	routes.AddRoutes(routes.Route{
		Version:     "v1",
		HandlerFunc: GetServices,
		Pattern:     "/services/datastore/get",
	})
}
