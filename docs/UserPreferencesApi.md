# \UserPreferencesApi

All URIs are relative to *http://localhost:10081/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserPreferences**](UserPreferencesApi.md#GetUserPreferences) | **Get** /user/preferences | Get user preferences.
[**UpdateUserPreferences**](UserPreferencesApi.md#UpdateUserPreferences) | **Put** /user/preferences | Update user preferences.


# **GetUserPreferences**
> Prefs GetUserPreferences(ctx, )
Get user preferences.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Prefs**](Prefs.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserPreferences**
> SuccessResponseBody UpdateUserPreferences(ctx, body)
Update user preferences.

Omitting a key (`theme`, `homeDashboardId`, `timezone`) will cause the current value to be replaced with the system default value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePrefsCmd**](UpdatePrefsCmd.md)|  | 

### Return type

[**SuccessResponseBody**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

