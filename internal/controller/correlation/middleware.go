// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package correlation

import (
	"context"
	"net/http"
	"time"

	"github.com/FanQiang28/device-adsb/internal/common"
	"github.com/xinminsu/go-mod-core-contracts/clients"
	"github.com/google/uuid"
)

func ManageHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hdr := r.Header.Get(clients.CorrelationHeader)
		if hdr == "" {
			hdr = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), clients.CorrelationHeader, hdr)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func OnResponseComplete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		correlationId := FromContext(r.Context())
		if common.LoggingClient != nil {
			common.LoggingClient.Trace("Response complete", clients.CorrelationHeader, correlationId, "duration", time.Since(begin).String())
		}
	})
}

func OnRequestBegin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := FromContext(r.Context())
		if common.LoggingClient != nil {
			common.LoggingClient.Trace("Begin request", clients.CorrelationHeader, correlationId, "path", r.URL.Path)
		}
		next.ServeHTTP(w, r)
	})
}
