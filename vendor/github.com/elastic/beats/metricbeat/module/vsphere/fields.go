// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package vsphere

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "vsphere", asset.ModuleFieldsPri, AssetVsphere); err != nil {
		panic(err)
	}
}

// AssetVsphere returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/vsphere.
func AssetVsphere() string {
	return "eJzMlsFymzAQhu9+ip3ckwfg0Es6aS9uO5OmV48Mi1GNWI20JEOfviMJPBjkJDWirQ6eMQv//3nF/tYtHLHL4NnqCg1uAFhyjRncPD/6KzcbgAJtbqRmSU0GHzYAAH0VFBVt7R4zWKOwmMEeWWwASol1YTN/8y00QuHYxC3uNGZwMNTq/krE51xoLFYIFpbpJBeXvCjblyIi018yrCnIGMZ9nhUGliN2L2SKSe0VIrc+DlRz3cGwtE4/neWDrNF2llHBTHjwzIUWueTujolFfbfvGG2UoKbm8Gf2350ieEWgErjC6Na4VZJRgjOY2884S4OYFPPBICanbC0WSSmfLBbrUOqcUzJqNDk2/F7K/vZTdRYJFVlekgaT5/99EHwmy5czINdt2BZV/Uq5LfffnkA2sJ2ojm1DAqTzDfP/DmM/0ul8/UC/YatQkVlrTLde3NnHpN+ezx5urUROhLdOEC+Fa5BfyBx37lsc7Kqh/RJk4Vz2dAaShltRK5FXsll0drmodE1uTaJvYQtez620GfkjtAG2oQ+XbSnhHk9Nv2o0gmVzgMdwiPrfQnscYocWLa8WZVTCJ2ewODQ8q3sv10T1b2qa9E3f1rMMTtZXH8bpYceRvJw1by2T2oXIihLS/ifOjqPh4m7BZN97Y4gY/9W/jd8BAAD//8WY8bw="
}
