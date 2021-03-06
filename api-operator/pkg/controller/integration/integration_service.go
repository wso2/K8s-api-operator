/*
 * Copyright (c) 2021 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package integration

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"strconv"
)

// serviceForIntegration returns a service object
func (r *ReconcileIntegration) serviceForIntegration(config EIConfigNew) *corev1.Service {

	m := config.integration

	//set HTTP and HTTPS ports for as ServiceSpec ports
	exposeServicePorts := []corev1.ServicePort{
		{
			Name:       m.Name + strconv.Itoa(int(m.Spec.Expose.PassthroPort)),
			Port:       m.Spec.Expose.PassthroPort,
			TargetPort: intstr.FromInt(int(m.Spec.Expose.PassthroPort)),
		},
	}

	// check inbound endpoint port is exist and append to the container port
	for _, port := range m.Spec.Expose.InboundPorts {
		exposeServicePorts = append(
			exposeServicePorts,
			corev1.ServicePort{
				Name: m.Name + strconv.Itoa(int(port)),
				Port: port,
				TargetPort: intstr.IntOrString{
					Type:   Int,
					IntVal: port,
				},
			},
		)
	}

	labels := labelsForIntegration(m.Name)

	service := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      nameForService(&m),
			Namespace: m.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports:    exposeServicePorts,
		},
	}
	// Set Integration instance as the owner and controller
	controllerutil.SetControllerReference(&m, service, r.scheme)
	return service
}
