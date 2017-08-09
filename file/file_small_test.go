// +build small

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"testing"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

var mockMts = []plugin.Metric{
	{
		Namespace: plugin.NewNamespace("foo"),
		Timestamp: time.Now(),
		Version:   99,
		Data:      "foo",
	},
}

func TestFilePublisher(t *testing.T) {
	Convey("Create a File Publisher", t, func() {
		fp := New()
		Convey("So file publisher should not be nil", func() {
			So(fp, ShouldNotBeNil)
		})
		Convey("So file publisher should be of publisher plugin type", func() {
			So(fp, ShouldHaveSameTypeAs, &filePublisher{})
		})

		Convey("Test GetConfigPolicy()", func() {
			configPolicy, err := fp.GetConfigPolicy()

			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So getting a config policy should not return an error", func() {
				So(err, ShouldBeNil)
			})

			Convey("So config policy should be a cpolicy.ConfigPolicy type", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})

		Convey("Publish content to file", func() {
			config := plugin.Config{}
			config["file"] = "/tmp/pub.out"
			err := fp.Publish(mockMts, config)
			So(err, ShouldBeNil)
		})
	})
}

func TestFormatMetricTypes(t *testing.T) {
	Convey("FormatMetricTypes returns metrics to publish", t, func() {
		metrics := formatMetricTypes(mockMts)
		So(metrics, ShouldNotBeEmpty)
		// formatted metric has namespace represented as a single string
		So(metrics[0].Namespace, ShouldEqual, mockMts[0].Namespace.String())
	})
}
