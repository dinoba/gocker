gocker

Collects logs from running docker containers and send it to storage (Elasticsearch 6.x && Kafka supported).

## Description
App starts docker client and reads all running docker instances.

For each docker (if is not marked for skiping in config) new routine in started.

Each routine:
Find log for assigned container (found in /var/lib/docker/containers + dockerid)  
Use tail to fetch from each log. 
Parse log and send via channel to storage.

## Usage
```bash
go build
./gocker #user must have docker priviledges 
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
