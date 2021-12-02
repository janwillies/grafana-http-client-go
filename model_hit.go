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

type Hit struct {
	FolderId int64 `json:"folderId,omitempty"`
	FolderTitle string `json:"folderTitle,omitempty"`
	FolderUid string `json:"folderUid,omitempty"`
	FolderUrl string `json:"folderUrl,omitempty"`
	Id int64 `json:"id,omitempty"`
	IsStarred bool `json:"isStarred,omitempty"`
	Slug string `json:"slug,omitempty"`
	SortMeta int64 `json:"sortMeta,omitempty"`
	SortMetaName string `json:"sortMetaName,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Title string `json:"title,omitempty"`
	Type_ *HitType `json:"type,omitempty"`
	Uid string `json:"uid,omitempty"`
	Uri string `json:"uri,omitempty"`
	Url string `json:"url,omitempty"`
}
