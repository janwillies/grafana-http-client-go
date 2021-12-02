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

type UpdateAlertNotificationWithUidCommand struct {
	OrgId int64 `json:"OrgId,omitempty"`
	Result *AlertNotification `json:"Result,omitempty"`
	DisableResolveMessage bool `json:"disableResolveMessage,omitempty"`
	Frequency string `json:"frequency,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Name string `json:"name,omitempty"`
	SecureSettings map[string]string `json:"secureSettings,omitempty"`
	SendReminder bool `json:"sendReminder,omitempty"`
	Settings *simplejson.Json `json:"settings,omitempty"`
	Type_ string `json:"type,omitempty"`
	Uid string `json:"uid,omitempty"`
}
