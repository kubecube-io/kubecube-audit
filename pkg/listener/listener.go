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
	"context"
	"github.com/kubecube-io/kubecube/pkg/apis"
	hotplugv1 "github.com/kubecube-io/kubecube/pkg/apis/hotplug/v1"
	"github.com/kubecube-io/kubecube/pkg/clog"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	toolcache "k8s.io/client-go/tools/cache"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
)

const (
	hotPlugNameCommon         = "common"
	hotPlugComponentNameAudit = "audit"
	hotPlugAuditEnabled       = "enabled"
)

func Listener() {

	config, err := ctrl.GetConfig()
	if err != nil {
		panic(err)
	}

	scheme := runtime.NewScheme()
	utilruntime.Must(apis.AddToScheme(scheme))

	c, err := cache.New(config, cache.Options{Scheme: scheme})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	hotPlug := &hotplugv1.Hotplug{}
	hotPlug.Name = hotPlugNameCommon
	informer, err := c.GetInformer(ctx, hotPlug)
	if err != nil {
		panic(err)
	}

	informer.AddEventHandler(toolcache.ResourceEventHandlerFuncs{
		UpdateFunc: func(oldObj, newObj interface{}) {
			hotplug, ok := newObj.(*hotplugv1.Hotplug)
			if !ok {
				clog.Error("watch an error obj: %+v", newObj)
				return
			}
			components := hotplug.Spec.Component
			for _, component := range components {
				if component.Name == hotPlugComponentNameAudit {
					if component.Status == hotPlugAuditEnabled {
						backend.EnableAudit = true
					} else {
						backend.EnableAudit = false
					}

				}
			}
		},
	})

	err = c.Start(ctx)
	if err != nil {
		panic(err)
	}
}
