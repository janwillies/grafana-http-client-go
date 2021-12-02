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

type UserSearchHitDto struct {
	AuthLabels []string `json:"authLabels,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Email string `json:"email,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsAdmin bool `json:"isAdmin,omitempty"`
	IsDisabled bool `json:"isDisabled,omitempty"`
	LastSeenAt time.Time `json:"lastSeenAt,omitempty"`
	LastSeenAtAge string `json:"lastSeenAtAge,omitempty"`
	Login string `json:"login,omitempty"`
	Name string `json:"name,omitempty"`
}
