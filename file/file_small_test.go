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
	"bytes"
	"encoding/gob"
	"errors"
	"testing"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

var mockMts = []plugin.MetricType{
	*plugin.NewMetricType(core.NewNamespace("foo"), time.Now(), nil, "", 99),
}

func TestMetaData(t *testing.T) {
	Convey("Meta returns proper metadata", t, func() {
		meta := Meta()
		So(meta, ShouldNotBeNil)
		So(meta.Name, ShouldResemble, PluginName)
		So(meta.Version, ShouldResemble, PluginVersion)
		So(meta.Type, ShouldResemble, PluginType)
	})
}

func TestFilePublisher(t *testing.T) {
	Convey("Create a File Publisher", t, func() {
		fp := NewFilePublisher()
		Convey("so file publisher should not be nil", func() {
			So(fp, ShouldNotBeNil)
		})
		Convey("so file publisher should be of publisher plugin type", func() {
			So(fp, ShouldHaveSameTypeAs, &filePublisher{})
		})

		configPolicy, err := fp.GetConfigPolicy()

		Convey("Test GetConfigPolicy()", func() {
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So getting a config policy should not return an error", func() {
				So(err, ShouldBeNil)
			})

			Convey("So config policy should be a cpolicy.ConfigPolicy type", func() {
				So(configPolicy, ShouldHaveSameTypeAs, &cpolicy.ConfigPolicy{})
			})
		})
		Convey("Publish content to file", func() {
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			enc.Encode(mockMts)

			config := make(map[string]ctypes.ConfigValue)
			config["file"] = ctypes.ConfigValueStr{Value: "/tmp/pub.out"}

			Convey("invalid contentType", func() {
				err := fp.Publish("", buf.Bytes(), config)
				So(err, ShouldResemble, errors.New("Unknown content type ''"))
			})
			Convey("empty content", func() {
				err = fp.Publish(plugin.SnapGOBContentType, []byte{}, config)
				So(err, ShouldNotBeNil)
			})
			Convey("successful publishing", func() {
				err = fp.Publish(plugin.SnapGOBContentType, buf.Bytes(), config)
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestFormatMetricTypes(t *testing.T) {
	Convey("FormatMetricTypes returns metrics to publish", t, func() {
		metrics := formatMetricTypes(mockMts)
		So(metrics, ShouldNotBeEmpty)
		// formatted metric has namespace represented as a single string
		So(metrics[0].Namespace, ShouldEqual, mockMts[0].Namespace().String())
	})
}
