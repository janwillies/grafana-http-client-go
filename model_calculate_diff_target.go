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

type CalculateDiffTarget struct {
	DashboardId int64 `json:"dashboardId,omitempty"`
	UnsavedDashboard *simplejson.Json `json:"unsavedDashboard,omitempty"`
	Version int64 `json:"version,omitempty"`
}
