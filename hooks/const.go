// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2021 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package hooks

// ConfToEnv defines mappings from snap config keys to EdgeX environment variable
// names that are used to override individual device-mqtt's [Driver]  configuration
// values via a .env file read by the snap service wrapper.
//
// The syntax to set a configuration key is:
//
// env.<section>.<keyname>
//
var ConfToEnv = map[string]string{
	// [Driver]
	"driver.incoming-schema":         "DRIVER_INCOMINGSCHEMA",
	"driver.incoming-host":           "DRIVER_INCOMINGHOST",
	"driver.incoming-port":           "DRIVER_INCOMINGPORT",
	"driver.incoming-password":       "DRIVER_INCOMINGPASSWORD",
	"driver.incoming-qos":            "DRIVER_INCOMINGQOS",
	"driver.incoming-keep-alive":     "DRIVER_INCOMINGKEEPALIVE",
	"driver.incoming-client-id":      "DRIVER_INCOMINGCLIENTID",
	"driver.incoming-topic":          "DRIVER_INCOMINGTOPIC",
	"driver.response-schema":         "DRIVER_RESPONSESCHEMA",
	"driver.response-host":           "DRIVER_RESPONSEHOST",
	"driver.response-port":           "DRIVER_RESPONSEPORT",
	"driver.response-password":       "DRIVER_RESPONSEPASSWORD",
	"driver.response-qos":            "DRIVER_RESPONSEQOS",
	"driver.response-keep-alive":     "DRIVER_RESPONSEKEEPALIVE",
	"driver.response-client-id":      "DRIVER_RESPONSECLIENTID",
	"driver.response-topic":          "DRIVER_RESPONSETOPIC",
	"driver.conn-establishing-retry": "DRIVER_CONNESTABLISHINGRETRY",
	"driver.conn-retry-wait-time":    "DRIVER_CONNRETRYWAITTIME",
}
