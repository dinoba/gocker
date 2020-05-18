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

check logs

```bash
2020-05-18 17:12:40 [INFO] Starting
2020-05-18 17:12:45 [INFO] Start log collecting on 192.168.1.182
2020-05-18 17:15:45 [INFO] New docker found nginx
2020-05-18 17:15:45 [INFO] read docker log file /var/lib/docker/containers/d9789310fddb3d2c75c087d3ca68bfff4ae92166e5159b2dab0777a4f85e1bf3/d9789310fddb3d2c75c087d3ca68bfff4ae92166e5159b2dab0777a4f85e1bf3-json.log
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
