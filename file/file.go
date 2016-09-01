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
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/ctypes"
)

const (
	PluginName    = "file"
	PluginVersion = 3
	PluginType    = plugin.PublisherPluginType
)

type filePublisher struct {
}

type MetricToPublish struct {
	// The timestamp from when the metric was created.
	Timestamp time.Time         `json:"timestamp"`
	Namespace string            `json:"namespace"`
	Data      interface{}       `json:"data"`
	Unit      string            `json:"unit"`
	Tags      map[string]string `json:"tags"`
	Version_  int               `json:"version"`
	// Last advertised time is the last time the snap agent was told about a metric.
	LastAdvertisedTime time.Time `json:"last_advertised_time"`
}

//NewFilePublisher returns an instance of filePublisher
func NewFilePublisher() *filePublisher {
	return &filePublisher{}
}

func (f *filePublisher) Publish(contentType string, content []byte, config map[string]ctypes.ConfigValue) error {
	logger := log.New()
	logger.Println("Publishing started")
	var mts []plugin.MetricType

	switch contentType {
	case plugin.SnapGOBContentType:
		dec := gob.NewDecoder(bytes.NewBuffer(content))
		if err := dec.Decode(&mts); err != nil {
			logger.Printf("Error decoding: error=%v content=%v", err, content)
			return fmt.Errorf("Error decoding %v", err)
		}
	default:
		return fmt.Errorf("Unknown content type '%s'", contentType)
	}

	logger.Printf("publishing %v metrics to %v", len(mts), config)
	file, err := os.OpenFile(config["file"].(ctypes.ConfigValueStr).Value, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		logger.Printf("Error: %v", err)
		return err
	}
	w := bufio.NewWriter(file)

	// format metrics types to metrics to be published
	metrics := formatMetricTypes(mts)
	jsonOut, err := json.Marshal(metrics)
	if err != nil {
		return fmt.Errorf("Error while marshalling metrics to JSON: %v", err)
	}

	w.Write(jsonOut)
	w.WriteString("\n")
	w.Flush()

	return nil
}

//Meta returns metadata about the plugin
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(PluginName, PluginVersion, PluginType, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

func (f *filePublisher) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	config := cpolicy.NewPolicyNode()

	r1, err := cpolicy.NewStringRule("file", true)
	handleErr(err)
	r1.Description = "Absolute path to the output file for publishing"

	config.Add(r1)
	cp.Add([]string{""}, config)
	return cp, nil
}

// formatMetricTypes returns metrics in format to be publish as a JSON based on incoming metrics types;
// i.a. namespace is formatted as a single string
func formatMetricTypes(mts []plugin.MetricType) []MetricToPublish {
	var metrics []MetricToPublish
	for _, mt := range mts {
		metrics = append(metrics, MetricToPublish{
			Timestamp:          mt.Timestamp(),
			Namespace:          mt.Namespace().String(),
			Data:               mt.Data(),
			Unit:               mt.Unit(),
			Tags:               mt.Tags(),
			Version_:           mt.Version(),
			LastAdvertisedTime: mt.LastAdvertisedTime(),
		})
	}
	return metrics
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}
