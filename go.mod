module audit

go 1.15

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator/v10 v10.6.0 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/kubecube-io/kubecube v1.0.0
	github.com/olivere/elastic/v7 v7.0.24
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.1-0.20210326183817-17c1766b6349
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210506145944-38f3c27a63bf // indirect
	golang.org/x/net v0.0.0-20210508051633-16afe75a6701 // indirect
	golang.org/x/sys v0.0.0-20210507161434-a76c4d0a0096 // indirect
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
	k8s.io/apiserver v0.20.5
	k8s.io/client-go v0.20.5
	sigs.k8s.io/controller-runtime v0.8.3
)
