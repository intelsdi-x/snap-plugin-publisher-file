# Example tasks

[This](tasks/mock-passthru-file.json) example task will publish metrics to **file** 
from the mock plugin.  

## Running the example

### Requirements
 * `docker` and `docker-compose` are **installed** and **configured** 

Running the sample is as *easy* as running the script `./run-mock-passthru-file.sh`. 

## Files

- [run-mock-passthru-file.sh](run-mock-passthru-file.sh) 
    - The example is launched with this script     
- [tasks/mock-passthru-file.json](tasks/mock-passthru-file.json)
    - Snap task definition
- [docker-compose.yml](docker-compose.yml)
    - A docker compose file which defines two linked containers
        - "runner" is the container where snapd is run from.  You will be dumped 
        into a shell in this container after running 
        [run-mock-passthru-file.sh](run-mock-passthru-file.sh).  Exiting the shell will 
        trigger cleaning up the containers used in the example.
- [mock-passthru-file.sh](mock-passthru-file.sh)
    - Downloads `snapd`, `snapctl`, `snap-plugin-publisher-file`,
    `snap-plugin-processor-passthru`, `snap-plugin-collector-mock1` and starts the task 
    [tasks/mock-passthru-file.json](tasks/mock-passthru-file.json).
- [.setup.sh](.setup.sh)
    - Verifies dependencies and starts the containers.  It's called 
    by [run-mock-passthru-file.sh](run-mock-passthru-file.sh).