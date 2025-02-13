## gocker

Gocker collects logs from running Docker containers and sends them to storage. It currently supports Elasticsearch 6.x and Kafka.

## Description
The application retrieves information about active Docker containers using the Docker API.
```go
containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
```  
For each running container (unless explicitly marked for exclusion in the configuration), a new Go routine is spawned to handle log collection.

Each routine:  
Use tail to fetch from each log.  

```go
func ReadLog(container Docker, logsChan chan storage.DockerLog) {
```  

Parse log and send via channel to storage.  
```go
logsChan <- parseLog(rawLog, container)
``` 

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

If you get errors like  
```bash
Error response from daemon: client version 1.41 is too new. Maximum supported API version is 1.40
```
Force API with following ENV variable

```bash
export DOCKER_API_VERSION='1.40'
```
## Elasticserach template

Use storage/es-template.tpl template to define ES mappings

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
