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

type LibraryElementDtoMeta struct {
	ConnectedDashboards int64 `json:"connectedDashboards,omitempty"`
	Created time.Time `json:"created,omitempty"`
	CreatedBy *LibraryElementDtoMetaUser `json:"createdBy,omitempty"`
	FolderName string `json:"folderName,omitempty"`
	FolderUid string `json:"folderUid,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	UpdatedBy *LibraryElementDtoMetaUser `json:"updatedBy,omitempty"`
}