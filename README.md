# gocker



Collects logs from running docker containers and send it to storage (Elasticsearch 6.x && Kafka supported).

App starts docker client and reads all running docker instances.
For each docker (if is not marked for skiping in config) new routine in started.
Routine find log for each running container (found in /var/lib/docker/containers + dockerid)  
Routine use tail to fetch from each log. Docker log is parsed and sent via channel to storage.

