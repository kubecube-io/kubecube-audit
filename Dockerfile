# Copyright 2021 The KubeCube Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.13.5 AS builder

WORKDIR /go/src/kubecube-audit
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o audit main.go

ENTRYPOINT ["/audit"]