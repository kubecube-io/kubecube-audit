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

package env

import "os"

const (
	//todo config service es.kubecube-system.svc
	defaultEsHost  = "http://10.219.192.164:30007"
	defaultEsIndex = "audit"
	defaultEsType  = "logs"
	defaultPort    = "8888"
)

type EsWebhook struct {
	Host  string
	Index string
	Type  string
}

func Webhook() *EsWebhook {
	host := os.Getenv("AUDIT_WEBHOOK_HOST")
	index := os.Getenv("AUDIT_WEBHOOK_INDEX")
	types := os.Getenv("AUDIT_WEBHOOK_TYPE")
	if host == "" || index == "" || types == "" {
		return &EsWebhook{defaultEsHost, defaultEsIndex, defaultEsType}
	}
	return &EsWebhook{host, index, types}
}

func JwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func Port() string {
	p := os.Getenv("PORT")
	if p == "" {
		return defaultPort
	}
	return p
}
