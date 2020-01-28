/*
 * Luminate API
 *
 *  ## Introduction  Luminate API uses common RESTful resourced based URL conventions and JSON as the exchange format. <br> Properties names are case-sensitive. <br> Some of Luminate API calls omit None values from the API response.  The base-URL is `api.`&lt;`tenant-name`&gt;`.luminatesec.com`. For example, if your administration portal URL is _admin.acme.luminatesec.com_, then your API base-URL is _api.acme.luminatesec.com_.  All examples below are performed on a tenant called acme.  ## Common Operations Steps  Below you may find a list of  common operations and the relevant API calls for each. Each of these operations can also be performed by using the administrative portal at https://admin.acme.luminatesec.com.  <ol>   <li>     Creating a site and deploying a connector:     <ol type=\"a\">       <li>Creating a new site using the <a href=\"#operation/createSite\">Create site API</a>.<br></li>       <li>         Once a site is created you can use its Id (returned in the response of the Create Site request)         and call the <a href=\"#operation/createConnector\">Create connector API</a>. <br>       </li>       <li>         Deploy the Luminate connector:         <ol type=\"i\">           <li>Retrieve the deployment command using the <a href=\"#operation/getConnectorCommand\">Connector Deployment Command API.</a> <br> </li>           <li>Execute the command on the target machine.</li>         </ol>       </li>     </ol>   </li>   <li>     Creating an application:       <ol type=\"a\">         <li>           An application is always associated with a specific site in order to route the traffic to the application via the connectors associated with the same site.           In order to create the application, call the <a href=\"#operation/createApplication\">Create Application API</a>         </li>         <li>           Once the application is created, you *must* assign the application to a specific site in order to make it accessible. Assign the application to the required site           using the <a href=\"#operation/BindApplicationToSite\">Bind Application to Site API</a>.         </li>         <li>           Once the application is assigned to a site, you should update the access permissions for the application to allow specific entities (users/groups)           to access the application using the <a href=\"#operation/UpdateAccessPolicyMultipleAssignments\">Set/Update Access Policy API</a>.         </li>       </ol>   </li> </ol>  ## Object Model The object model of the API is built around the following: <ol>   <li><a href=\"#tag/Sites\">Sites</a> - Site is a representation of the physical or virtual data center your applications reside in.</li>   <li><a href=\"#tag/Connectors\">Connectors</a> - A connector is a lightweight piece of software connecting your site to the Luminate platform.</li>   <li><a href=\"#tag/Applications\">Applications</a>  - Application is the internal resource you would like to publish using Luminate. </li>   <li>     <a href=\"#tag/Policies\">Policies</a> - Secure Access Cloud continuously enforce contextual access and activity       policies to control access to resources and restrict user’s actions within resources, based on the user/device       context (such as the user’s group membership, user’s location, MFA status and managed/unmanaged       device status) and the requested resource.   <li>     Logs - Luminate internal logs for audit and forensics purposes:     <ol>       <li><a href=\"#tag/Audit-Logs\">Audit Logs</a> audit all operations done through the administration portal</li>       <li><a href=\"#tag/Web-Access-Logs\">Web Access Logs</a> audit any web access</li>       <li><a href=\"#tag/SSH-Logs\">SSH Logs</a> audit any SSH access</li>       <li><a href=\"#tag/RDP-Logs\">RDP Logs</a> audit any RDP access</li>     </ol>   </li> </ol>   ## Authentication  Authentication is done using [OAuth2](https://tools.ietf.org/html/rfc6749) with the [Bearer authentication scheme](https://tools.ietf.org/html/rfc6750).  <!-- ReDoc-Inject: <security-definitions> -->  The Luminate API is available to Luminate users who have administrative privileges in their Luminate tenant. An administrator should create an API client through the Luminate Admin portal, check the ‘Allow access to Luminate management API’ permission and copy the ‘Client Id’ and the ‘Client Secret’.  Retrieving the API access token is done using Basic-Authentication scheme, POST of a Base64 encoded Client-ID and Client-Secret: <B>   ``` curl -X POST \\  https://api.acme.luminatesec.com/v1/oauth/token \\  -u yourApiClientId:yourApiClientSecret   ``` </B>  This call returns the following JSON: {     \"access_token\":\"edfe22e3-eb4c-4c83-8ce3-3152e6a2XXX\",   \"expires_in\":86400,   \"scope\":\"luminate-scope\",   \"token_type\":\"Bearer\",   \"error\":\"\",   \"error_description\":\"\"}  All further API calls should include the ‘Authorization’ header with value “Bearer AccessToken”  For example: <B>   ```   curl -H \"Authorization: Bearer edfe22e3-eb4c-4c83-8ce3-3152e6a2XXX\" \"https://api.acme.luminatesec.com/v2/applications/d9f6ca17-9f2c-488c-aa86-51924a37092e\"   ``` </B>  ## Versioning and Compatibility  The latest Major Version is `v2`.  The Major Version is included in the URL path (e.g. /v2/applications ) and it denotes breaking changes to the API. Minor and Patch versions are transparent to the client.  ## Pagination   Some of our API responses are paginated, meaning that only a certain number of items are returned at a time.  The default number of items returned in a single page is 50.  You can override this by passing a size parameter to set the maximum number of results, but cannot exceed 100.  Specifying the page number sets the starting point for the result set, allowing you to fetch subsequent items  that are not in the initial set of results. The sort order for returned data can be controlled using the sort parameter.<br>  You can constrain the results by using a filter. <br><br>  **Note:** Most methods that support pagination use the approach specified above. However, some methods use varied   versions of pagination. The individual documentation for each API method is your source of truth for which pattern the method follows.  ## Auditing  All authentication operations and modify operations (POST, PUT, DELETE) are audited.   ## Rate-limiting The API has a rate limit of 5 requests per second. If you have hit the rate limit, then a ‘429’ status code will be returned. In such cases, you should back-off from submitting new requests for 1 second before resuming.  Note that rate-limitation applies to the accumulated requests of **all** of your clients. For example, if you have 6 clients submitting requests simultaneously at a rate of 1 request per second for each one then one of them is likely to get a 429 status code.  ## Support  For additional help you may refer to our support at https://support.luminate.io.  Each request submitted to the API returns a unique request ID that is generated by the API. The request ID will be returned in header `x-lum-request-id`. If you need to contact us about any specific request then this ID will serve as a reference to the given request.
 *
 * API version: V2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ConnectorsApiService service

/*
ConnectorsApiService Delete Connector.
Delete the connector from the contained site in your Luminate tenant. &lt;br&gt; In order to complete the action the connector container that runs in the corresponding datacenter should be removed.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connectorId Connector ID.


*/
func (a *ConnectorsApiService) ConnectorsByConnectorIdDelete(ctx context.Context, connectorId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connectors/{connector-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector-id"+"}", fmt.Sprintf("%v", connectorId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 0 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarHttpResponse, newErr
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ConnectorsApiService Get Connector
Returns the details of a Connector from your Luminate tenant.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connectorId Connector ID.

@return Connector
*/
func (a *ConnectorsApiService) ConnectorsByConnectorIdGet(ctx context.Context, connectorId string) (Connector, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Connector
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connectors/{connector-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector-id"+"}", fmt.Sprintf("%v", connectorId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Connector
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 404 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 0 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ConnectorsApiService getConnectorCommand
Returns the command for deploying Luminate connector as a docker image. This endpoint is valid for connectors of version 2.5.10 and higher. The command is generated based on the deployment_type that was set for the connector:  windows / linux / docker-compose / Kubernetes. Executing the command on the target machine is the last step in configuring the Luminate Connector.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connectorId Connector ID.

@return ConnectorDeploymentCommand
*/
func (a *ConnectorsApiService) ConnectorsCommandByConnectorIdGet(ctx context.Context, connectorId string) (ConnectorDeploymentCommand, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ConnectorDeploymentCommand
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connectors/{connector-id}/command"
	localVarPath = strings.Replace(localVarPath, "{"+"connector-id"+"}", fmt.Sprintf("%v", connectorId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v ConnectorDeploymentCommand
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 404 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 0 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ConnectorsApiService getAllConnectors
Return an array of paginated JSON objects. Each object represents a connector configured in your Luminate tenant. &lt;br&gt; The supported sort keys are either ‘name’ for sorting by connector name or ‘id’ for sorting by connector id.&lt;br&gt; Filter applies for the following fields: \&quot;name\&quot;, \&quot;internal-address\&quot;, \&quot;external-address\&quot;, \&quot;custom-domain\&quot; and \&quot;custom-root-path\&quot;. &lt;br&gt; Using the query filter&#x3D;test will return all the connectors for which one or more of the above listed fields contain \&quot;test\&quot;.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ConnectorsGetOpts - Optional Parameters:
     * @param "Sort" (optional.String) -  The value of this parameter is a comma-separated list of sort key and sort direction. By default, query results are sorted in ascending order by item id. The supported sort directions are either &#39;asc&#39; for ascending or &#39;desc&#39; for descending.
     * @param "Size" (optional.Int32) -  The number of items returned in a single page.
     * @param "Page" (optional.Int32) -  The page number.
     * @param "Filter" (optional.String) -  The string by which the results are filtered.

@return ConnectorsPage
*/

type ConnectorsGetOpts struct {
	Sort   optional.String
	Size   optional.Int32
	Page   optional.Int32
	Filter optional.String
}

func (a *ConnectorsApiService) ConnectorsGet(ctx context.Context, localVarOptionals *ConnectorsGetOpts) (ConnectorsPage, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ConnectorsPage
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connectors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v ConnectorsPage
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 0 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ConnectorsApiService createConnector
Creates a Connector in your Luminate tenant. To complete configuring the Luminate connector, you should retrieve the deployment command using &lt;a href&#x3D;\&quot;#operation/getConnectorCommand\&quot;&gt;Connector Deployment Command API&lt;/a&gt; and execute it on the target machine.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bindToSiteId The ID of the site that should contain this connector.
 * @param optional nil or *ConnectorsPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of Connector) -

@return Connector
*/

type ConnectorsPostOpts struct {
	Body optional.Interface
}

func (a *ConnectorsApiService) ConnectorsPost(ctx context.Context, bindToSiteId string, localVarOptionals *ConnectorsPostOpts) (Connector, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Connector
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connectors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("bind_to_site_id", parameterToString(bindToSiteId, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody, localVarOptionalBodyok := localVarOptionals.Body.Value().(Connector)
		if !localVarOptionalBodyok {
			return localVarReturnValue, nil, reportError("body should be Connector")
		}
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 201 {
			var v Connector
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 401 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 403 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 409 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 0 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
