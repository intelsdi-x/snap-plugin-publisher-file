[![Build Status](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-file.svg?branch=master)](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-file)

# snap publisher plugin - File

This plugin supports pushing metrics into a local file

It's used in the [snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

* [golang 1.5+](https://golang.org/dl/) (needed only for building)

### Installation

#### Download File plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

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
This builds the plugin in `./build/rootfs/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

## Documentation
To use this plugin you have to specify a config with the location of the file you want to write to:

```
# JSON
"config": {
    "file": "/tmp/snap_published_file.log"
}

# YAML
config: 
    file: "/tmp/snap_published_mock_file.log"

```

The plugin will write out all metrics serialized as JSON to the specified file. An example of this output is below:

```
[
    {
        "namespace": [
            {
                "Value": "intel",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "mock",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "foo",
                "Description": "",
                "Name": ""
            }
        ],
        "last_advertised_time": "2016-07-13T10:39:47.922391602-07:00",
        "version": 0,
        "config": {
            "name": "root",
            "password": "secret"
        },
        "data": 88,
        "tags": {
            "plugin_running_on": "testhost.local"
        },
        "Unit_": "",
        "description": "",
        "timestamp": "2016-07-13T10:40:27.183090674-07:00"
    },
    {
        "namespace": [
            {
                "Value": "intel",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "mock",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "host0",
                "Description": "name of the host",
                "Name": "host"
            },
            {
                "Value": "baz",
                "Description": "",
                "Name": ""
            }
        ],
        "last_advertised_time": "0001-01-01T00:00:00Z",
        "version": 0,
        "config": null,
        "data": 66,
        "tags": {
            "plugin_running_on": "testhost.local"
        },
        "Unit_": "",
        "description": "",
        "timestamp": "2016-07-13T10:40:27.183093127-07:00"
    },
    {
        "namespace": [
            {
                "Value": "intel",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "mock",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "host1",
                "Description": "name of the host",
                "Name": "host"
            },
            {
                "Value": "baz",
                "Description": "",
                "Name": ""
            }
        ],
        "last_advertised_time": "0001-01-01T00:00:00Z",
        "version": 0,
        "config": null,
        "data": 69,
        "tags": {
            "plugin_running_on": "testhost.local"
        },
        "Unit_": "",
        "description": "",
        "timestamp": "2016-07-13T10:40:27.183094592-07:00"
    },
    {
        "namespace": [
            {
                "Value": "intel",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "mock",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "host2",
                "Description": "name of the host",
                "Name": "host"
            },
            {
                "Value": "baz",
                "Description": "",
                "Name": ""
            }
        ],
        "last_advertised_time": "0001-01-01T00:00:00Z",
        "version": 0,
        "config": null,
        "data": 80,
        "tags": {
            "plugin_running_on": "testhost.local"
        },
        "Unit_": "",
        "description": "",
        "timestamp": "2016-07-13T10:40:27.183096041-07:00"
    },
    {
        "namespace": [
            {
                "Value": "intel",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "mock",
                "Description": "",
                "Name": ""
            },
            {
                "Value": "host3",
                "Description": "name of the host",
                "Name": "host"
            },
            {
                "Value": "baz",
                "Description": "",
                "Name": ""
            }
        ],
        "last_advertised_time": "0001-01-01T00:00:00Z",
        "version": 0,
        "config": null,
        "data": 67,
        "tags": {
            "plugin_running_on": "testhost.local"
        },
        "Unit_": "",
        "description": "",
        "timestamp": "2016-07-13T10:40:27.183096841-07:00"
    }
]
```

### Examples
Full task configs using the mock and psutil plugins are available in [examples](examples)

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