/*
Copyright 2021 KubeCube Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"audit/pkg/audit"
	"audit/pkg/backend"
	"audit/pkg/listener"
	"audit/pkg/utils/env"
	"github.com/gin-gonic/gin"
	"github.com/kubecube-io/kubecube/pkg/clients"
	"github.com/kubecube-io/kubecube/pkg/clog"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const apiPathAuditRoot = "/api/v1/cube/audit"

// @title Swagger KubeCube-Audit API
// @version 1.0
// @description This is KubeCube-Audit api documentation.
func main() {

	clients.InitCubeClientSetWithOpts(nil)

	listener.Listener()

	router := gin.Default()

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.POST(apiPathAuditRoot+"/k8s", audit.HandleK8sAuditLog)
	router.POST(apiPathAuditRoot+"/cube", audit.HandleCubeAuditLog)
	router.GET(apiPathAuditRoot, audit.SearchAuditLog)
	router.GET(apiPathAuditRoot+"/export", audit.ExportAuditLog)

	b := backend.NewBackend()
	go b.Run()

	err := router.Run(":" + env.Port())
	if err != nil {
		clog.Error("%s", err)
	}
	return
}
