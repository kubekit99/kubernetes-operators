/*******************************************************************************
*
* Copyright 2017 SAP SE
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You should have received a copy of the License along with this
* program. If not, you may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

package president

import (
	"fmt"
)

// Options to configure your vice president
type Options struct {
	KubeConfig          string
	VicePresidentConfig string

	ViceKeyFile string
	ViceCrtFile string

	IntermediateCertificate string
	RootCACertificate       string

	MinCertValidityDays int
	EnableValidateRemoteCertificate bool

	MetricPort                        int
	IsEnableAdditionalSymantecMetrics bool
}

// CheckOptions verifies the Options and sets default values, if necessary
func (o *Options) CheckOptions() error {
	if o.ViceCrtFile == "" {
		return fmt.Errorf("path to vice certificate not provided. Aborting")
	}
	if o.ViceKeyFile == "" {
		return fmt.Errorf("path to vice key not provided. Aborting")
	}
	if o.VicePresidentConfig == "" {
		return fmt.Errorf("path to vice config not provided. Aborting")
	}
	if o.IntermediateCertificate == "" {
		LogDebug("Intermediate certificate not provided")
	}
	if o.KubeConfig == "" {
		LogDebug("Path to kubeconfig not provided. Using Default")
	}

	if o.MinCertValidityDays <= 0 {
		LogDebug("Minimum certificate validity invalid. Using default: 30 days")
		o.MinCertValidityDays = 30
	}

	if o.MetricPort == 0 {
		o.MetricPort = 9091
		LogDebug("Metric port not provided. Using default port: 9091")
	}
	if !o.IsEnableAdditionalSymantecMetrics {
		LogDebug("Not exposing additional Symantec metrics")
	} else {
		LogDebug("Exposing additional Symantec metrics")
	}

	return nil
}
