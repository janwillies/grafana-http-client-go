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

type ItemDto struct {
	AlertId int64 `json:"alertId,omitempty"`
	AlertName string `json:"alertName,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Created int64 `json:"created,omitempty"`
	DashboardId int64 `json:"dashboardId,omitempty"`
	Data *simplejson.Json `json:"data,omitempty"`
	Email string `json:"email,omitempty"`
	Id int64 `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	NewState string `json:"newState,omitempty"`
	PanelId int64 `json:"panelId,omitempty"`
	PrevState string `json:"prevState,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Text string `json:"text,omitempty"`
	Time int64 `json:"time,omitempty"`
	TimeEnd int64 `json:"timeEnd,omitempty"`
	Updated int64 `json:"updated,omitempty"`
	UserId int64 `json:"userId,omitempty"`
}
