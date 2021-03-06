// Copyright (c)  WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package envoy

const (
	mgBasePath                   = "/api/mgw/adapter/0.1"
	mgDeployResourcePath         = "/apis"
	mgDeleteAPIResourcePath      = "/apis"
	envoyMgwConfName             = "envoy-mgw-configs"
	envoyMgwSecretName           = "envoymgw-adapter-secret"
	mgwAdapterHostConst          = "mgwAdapterHost"
	mgwInsecureSkipVerifyConst   = "mgwInsecureSkipVerify"
	HeaderAuthorization          = "Authorization"
	HeaderAccept                 = "Accept"
	HeaderContentType            = "Content-Type"
	HeaderConnection             = "Connection"
	HeaderValueAuthBasicPrefix   = "Basic"
	HeaderValueKeepAlive         = "keep-alive"
	DefaultHttpRequestTimeout    = 10000
	mgwCertSecretName            = "mgwCertSecretName"
	usernameProperty             = "username"
	passwordProperty             = "password"
	apiNameProperty              = "apiName"
	versionProperty              = "version"
	apiOperatorConfigHome        = "API_OPERATOR_CONFIG_HOME"
	apiOperatorDefaultConfigHome = "/usr/local/bin"
)

// constants related to API-CTL project
// TODO: use API-CTL code to init project
const (
	apiYamlFile           = "api.yaml"
	swaggerDefinitionFile = "Definitions/swagger.yaml"
	deploymentEnvFile     = "deployment_environments.yaml"
	// TODO: use API-CTL code to init project
	deploymentEnvFileData = `type: deployment_environments
version: v4.0.0
data:
 -
   displayOnDevportal: true
   deploymentEnvironment: Default
`
)
