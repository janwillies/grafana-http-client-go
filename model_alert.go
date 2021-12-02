/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go_client

import (
	"time"
)

type Alert struct {
	Created time.Time `json:"Created,omitempty"`
	DashboardId int64 `json:"DashboardId,omitempty"`
	EvalData *simplejson.Json `json:"EvalData,omitempty"`
	ExecutionError string `json:"ExecutionError,omitempty"`
	For_ *Duration `json:"For,omitempty"`
	Frequency int64 `json:"Frequency,omitempty"`
	Handler int64 `json:"Handler,omitempty"`
	Id int64 `json:"Id,omitempty"`
	Message string `json:"Message,omitempty"`
	Name string `json:"Name,omitempty"`
	NewStateDate time.Time `json:"NewStateDate,omitempty"`
	OrgId int64 `json:"OrgId,omitempty"`
	PanelId int64 `json:"PanelId,omitempty"`
	Settings *simplejson.Json `json:"Settings,omitempty"`
	Severity string `json:"Severity,omitempty"`
	Silenced bool `json:"Silenced,omitempty"`
	State *AlertStateType `json:"State,omitempty"`
	StateChanges int64 `json:"StateChanges,omitempty"`
	Updated time.Time `json:"Updated,omitempty"`
	Version int64 `json:"Version,omitempty"`
}
