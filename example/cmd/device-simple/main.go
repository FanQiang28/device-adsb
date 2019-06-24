// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
	"github.com/FanQiang28/device-adsb"
	"github.com/FanQiang28/device-adsb/example/driver"
	"github.com/FanQiang28/device-adsb/pkg/startup"
)

const (
	serviceName string = "Device-adsb"
)

func main() {
	sd := driver.SimpleDriver{}
	startup.Bootstrap(serviceName, device.Version, &sd)
}
