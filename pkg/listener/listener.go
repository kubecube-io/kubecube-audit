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

package listener

import (
	"audit/pkg/backend"
	"github.com/kubecube-io/kubecube/pkg/clog"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"strings"
	"time"
)

func Listener() {
	stop := make(chan struct{})

	clientSet := getClientSet()
	for clientSet == nil {
		clientSet = getClientSet()
	}

	watcher := cache.NewListWatchFromClient(clientSet.CoreV1().RESTClient(), string(v1.ResourceConfigMaps), "kubecube-system",
		fields.OneTermEqualSelector("metadata.name", "hotplug"))
	_, configMapController := cache.NewInformer(
		watcher,
		&v1.ConfigMap{},
		time.Second*2,
		cache.ResourceEventHandlerFuncs{
			UpdateFunc: func(oldObj, newObj interface{}) {
				clog.Debug("type: %+v, data:%+v", reflect.TypeOf(newObj), newObj)
				configmap, ok := newObj.(*v1.ConfigMap)
				if !ok {
					clog.Error("watch an error obj: %+v", newObj)
					return
				}
				// parse configmap
				if auditProperties := configmap.Data["audit.properties"]; auditProperties != "" {
					propertyList := strings.Split(auditProperties, "\n")
					for _, i := range propertyList {
						kv := strings.Split(i, "=")
						if kv[0] == "enabled" {
							if kv[1] == "true" {
								backend.EnableAudit = true
							} else {
								backend.EnableAudit = false
							}
						}
					}
				}
			},
		},
	)
	go configMapController.Run(stop)
}

func getClientSet() *kubernetes.Clientset {
	config, err := ctrl.GetConfig()
	if err != nil {
		clog.Warn("get kubeconfig failed: %v, only allowed when testing", err)
		return nil
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		clog.Error("create clientset failed: %v", err)
		return nil
	}
	return clientSet
}
