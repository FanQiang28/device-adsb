// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package correlation

import (
	"context"

	"github.com/xinminsu/go-mod-core-contracts/clients"
)

func FromContext(ctx context.Context) string {
	hdr, ok := ctx.Value(clients.CorrelationHeader).(string)
	if !ok {
		hdr = ""
	}
	return hdr
}
