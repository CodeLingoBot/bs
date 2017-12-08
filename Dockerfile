# Copyright 2016 bs authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

FROM golang:1.9.0 as builder

COPY . /go/src/github.com/tsuru/bs
WORKDIR /go/src/github.com/tsuru/bs
RUN make build

FROM alpine:3.2
RUN  apk update && apk add conntrack-tools ca-certificates tzdata && rm -rf /var/cache/apk/*
COPY --from=builder /go/src/github.com/tsuru/bs/bs /bin/bs
ENTRYPOINT ["/bin/bs"]
