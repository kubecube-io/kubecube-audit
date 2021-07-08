# Copyright 2021 The KubeCube Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.16 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN git config --global url."https://JiahuiZhao11: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@github.com".insteadOf "https://github.com"
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -o audit main.go

FROM alpine:3.13.4
WORKDIR /
COPY --from=builder /workspace/audit .

ENTRYPOINT ["/audit"]