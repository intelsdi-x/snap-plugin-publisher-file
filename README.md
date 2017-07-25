[![Build Status](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-file.svg?branch=master)](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-file)

# Snap publisher plugin - File

This plugin supports pushing metrics into a local file

It's used in the [Snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

* [golang 1.7+](https://golang.org/dl/) (needed only for building)

### Installation

#### Download File plugin binary:
You can get the pre-built binaries for your OS and architecture at plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-publisher-file/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-publisher-file

Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-publisher-file.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation
To use this plugin you have to specify a config with the location of the file you want to write to.

```
# JSON
"config": {
    "file": "/tmp/snap_published_file.json"
}

# YAML
config: 
    file: "/tmp/snap_published_file.json"

```

The plugin will write out all metrics serialized as JSON to the specified file. An example of this output is below:

```json
[
  {
    "timestamp": "2016-07-25T11:27:59.795513984+02:00",
    "namespace": "/intel/mock/host0/baz",
    "data": 86,
    "unit": "",
    "tags": {
      "plugin_running_on": "my-machine"
    },
    "version": 0,
    "last_advertised_time": "0001-01-01T00:00:00Z"
  },
  {
    "timestamp": "2016-07-25T11:27:59.795514856+02:00",
    "namespace": "/intel/mock/host1/baz",
    "data": 70,
    "unit": "",
    "tags": {
      "plugin_running_on": "my-machine"
    },
    "version": 0,
    "last_advertised_time": "0001-01-01T00:00:00Z"
  },
  {
    "timestamp": "2016-07-25T11:27:59.795548989+02:00",
    "namespace": "/intel/mock/bar",
    "data": 82,
    "unit": "",
    "tags": {
      "plugin_running_on": "my-machine"
    },
    "version": 0,
    "last_advertised_time": "2016-07-25T11:27:21.852064032+02:00"
  },
  {
    "timestamp": "2016-07-25T11:27:59.795549268+02:00",
    "namespace": "/intel/mock/foo",
    "data": 72,
    "unit": "",
    "tags": {
      "plugin_running_on": "my-machine"
    },
    "version": 0,
    "last_advertised_time": "2016-07-25T11:27:21.852063228+02:00"
  }
]
```

### Examples

Example of running [psutil collector plugin](https://github.com/intelsdi-x/snap-plugin-collector-psutil), [movingaverage processor plugin](https://github.com/intelsdi-x/snap-plugin-processor-movingaverage), and writing data as a JSON to file:

Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

Ensure [snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `sudo snapteld -l 1 -t 0 &`


Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-processor-movingaverage/latest/linux/x86_64/snap-plugin-processor-movingaverage
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-psutil/latest/linux/x86_64/snap-plugin-collector-psutil
$ snaptel plugin load snap-plugin-publisher-file
$ snaptel plugin load snap-plugin-processor-movingaverage
$ snaptel plugin load snap-plugin-collector-psutil
```

Create a [task manifest](https://github.com/intelsdi-x/snap/blob/master/docs/TASKS.md) (see [exemplary tasks](examples/tasks/)),
for example `psutil-movingaverage-file.json` with following content:
```json
{
  "version": 1,
  "schedule": {
    "type": "simple",
    "interval": "1s"
  },
  "max-failures": 10,
  "workflow": {
    "collect": {
      "metrics": {
        "/intel/psutil/load/load1": {},
        "/intel/psutil/load/load15": {},
        "/intel/psutil/load/load5": {},
        "/intel/psutil/vm/available": {},
        "/intel/psutil/vm/free": {},
        "/intel/psutil/vm/used": {}
      },
      "process": [
        {
          "plugin_name": "movingaverage",
          "config": {
            "MovingAvgBufLength": 5
          }
        }
      ],
      "publish": [
        {
          "plugin_name": "file",
          "config": {
            "file": "/tmp/snap-psutil-movingaverage-file.log"
          }
        }
      ]
    }
  }
}

```

Create a task:
```
$ snaptel task create -t psutil-movingaverage-file.json
```

See JSON file containing the published data:
```
$ tailf /tmp/snap-psutil-movingaverage-file.json
```

To stop previously created task:
```
$ snaptel task stop <task_id>
```

### Roadmap

There isn't a current roadmap for this plugin. However, if additional output types are wanted, please open an issue or submit a pull request as mentioned below. 

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-file/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-publisher-file/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions! 

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
Original code from the [Snap](http://github.com/intelsdi-x/snap) repo.

Additional code written by:
* Author: [Taylor Thomas](https://github.com/thomastaylor312)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
