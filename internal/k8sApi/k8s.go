/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package k8sApi

import (
	"github.com/apache/dubbo-go/config_center"
)

import (
	"github.com/dubbogo/v3router/internal/k8sCRD"
)

const (
	GroupName    = "service.dubbo.apache.org"
	GroupVersion = "v1alpha1"
	Namespace    = "dubbo-workplace"
)

func SetK8sEventListener(listener config_center.ConfigurationListener) error {
	vsUniformRouterListenerHandler := newVirtualServiceListenerHandler(listener)
	drUniformRouterListenerHandler := newDestRuleListenerHandler(listener)
	k8sCRDClient, err := k8sCRD.NewK8sCRDClient(GroupName, GroupVersion, Namespace, vsUniformRouterListenerHandler, drUniformRouterListenerHandler)
	if err != nil {
		return err
	}
	k8sCRDClient.WatchResources()
	return nil
}
