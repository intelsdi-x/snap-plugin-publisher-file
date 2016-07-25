[![Build Status](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-file.svg?branch=master)](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-file)

# snap publisher plugin - File

This plugin supports pushing metrics into a local file

It's used in the [snap framework](http://github.com/intelsdi-x/snap).

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
To use this plugin you have to specify a config with the location of the file you want to write to.

```
# JSON
"config": {
    "file": "/tmp/snap_published_file.json"
}

# YAML
config: 
    file: "/tmp/snap_published_mock_file.json"

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
See full task manifests using the mock and psutil plugins which are available in [examples](examples)  


Example of running mock collector plugin, passthru processor plugin, and writing data as a JSON to file:

Make sure that your `$SNAP_PATH` is set, if not:
```
$ export SNAP_PATH=<snapDirectoryPath>/build
```
In one terminal window, open the snap daemon (in this case with logging set to 1 and trust disabled):
```
$ $SNAP_PATH/bin/snapd -l 1 -t 0
```
In another terminal window:  
Load snap-plugin-publisher-file plugin:
```
$ $SNAP_PATH/bin/snapctl plugin load build/rootfs/snap-plugin-publisher-file

Plugin loaded
Name: file
Version: 2
Type: publisher
Signed: false
Loaded Time: Mon, 25 Jul 2016 12:30:25 CEST
```
Load snap-plugin-processor-passthru plugin:
```
$ $SNAP_PATH/bin/snapctl plugin load $SNAP_PATH/plugin/snap-plugin-processor-passthru

Plugin loaded
Name: passthru
Version: 1
Type: processor
Signed: false
Loaded Time: Mon, 25 Jul 2016 12:33:00 CEST
```

Load snap-plugin-collector-mock2 plugin:
```
$ $SNAP_PATH/bin/snapctl plugin load $SNAP_PATH/plugin/snap-plugin-collector-mock2

Plugin loaded
Name: mock
Version: 2
Type: collector
Signed: false
Loaded Time: Mon, 25 Jul 2016 12:32:09 CEST
```

See available metrics for your system:
```
$ $SNAP_PATH/bin/snapctl metric list
```

Create a task manifest to use File publisher plugin (see [exemplary task manifest](examples/tasks/mock.json)):
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
                "/intel/mock/foo": {},
                "/intel/mock/bar": {},
                "/intel/mock/*/baz": {}
            },
            "config": {
                "/intel/mock": {
                    "name": "root",
                    "password": "secret"
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",                    
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "file",                            
                            "config": {
                                "file": "/tmp/snap_published_mock_file.json"
                            }
                        }
                    ]
                }
            ]
        }
    }
}
```

Create a task:
```
$ $SNAP_PATH/bin/snapctl task create -t examples/tasks/mock.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See JSON file containing the published data:
```
tail -f /tmp/snap_published_mock_file.json
```

To stop previously created task:
```
$ $SNAP_PATH/bin/snapctl task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
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
