// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafka

import (
	"github.com/kiegroup/kogito-cloud-operator/pkg/apis/app/v1alpha1"
	kafkabetav1 "github.com/kiegroup/kogito-cloud-operator/pkg/apis/kafka/v1beta1"
)

func updateAppPropsInStatus(kafkaInstance *kafkabetav1.Kafka, instance *v1alpha1.KogitoInfra) error {
	appProps, err := getKafkaAppProps(kafkaInstance)
	if err != nil {
		return err
	}
	instance.Status.AppProps = appProps
	return nil
}

func updateEnvVarsInStatus(kafkaInstance *kafkabetav1.Kafka, instance *v1alpha1.KogitoInfra) error {
	envVars, err := getKafkaEnvVars(kafkaInstance)
	if err != nil {
		return err
	}
	instance.Status.Env = envVars
	return nil
}
