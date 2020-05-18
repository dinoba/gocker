package dockerlog

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/hpcloud/tail"

	log "github.com/dinoba/gocker/log"
	"github.com/dinoba/gocker/storage"
)

//Docker object
type Docker struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	Status      string `json:"status"`
	State       string `json:"state"`
	Description string `json:"description"`
}

//RawLog log from docker
type RawLog struct {
	LogText string    `json:"log"`
	Stream  string    `json:"stream"`
	Time    time.Time `json:"time"`
}

//GetAllDockers get all existing dockers
func GetAllDockers(dockersToSkip string) ([]Docker, error) {
	var dockers []Docker
	skipDockers := strings.Split(dockersToSkip, ",")

	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		monitorDocker := true
		//check if this docker needs to be skiped for monitoring
		for _, word := range skipDockers {
			if strings.Contains(strings.ToLower(container.Image), strings.ToLower(word)) {
				monitorDocker = false
				continue
			}
		}
		if !monitorDocker {
			continue
		}

		var d Docker
		d.Id = container.ID
		d.Image = container.Image
		d.Status = container.Status
		d.State = container.State
		dockers = append(dockers, d)
	}

	return dockers, nil
}

//CheckForNewDockers checks if new dockers are created
func CheckForNewDockers(logsChan chan storage.DockerLog, dockersIDs *[]string, dockersToSkip string) (err error) {
	logger := log.GetInstance()
	files, err := ioutil.ReadDir("/var/lib/docker/containers/")
	if err != nil {
		return
	}

	if len(files) > len(*dockersIDs) {
		newDocker := true
		//read docker again
		var dockers []Docker
		dockers, err = GetAllDockers(dockersToSkip)
		if err != nil {
			log.WithPrefix(logger, err.Error(), "[ERROR] CheckForNewDockers() ")
			return
		}

		for _, v := range dockers {

			//check is docker id allready colected
			for _, id := range *dockersIDs {
				if v.Id == id {
					newDocker = false
				}
			}

			if newDocker {
				log.WithPrefix(logger, "New docker found "+v.Image, "[INFO]")
				*dockersIDs = append(*dockersIDs, v.Id)
				//wg.Add(1)
				go func(docker Docker) {
					//defer wg.Done()
					ReadLog(docker, logsChan)
				}(v)
			}
		}
	}
	return
}

func parseLog(rawLog RawLog, container Docker) (l storage.DockerLog) {
	l.LogText = rawLog.LogText
	l.Image = container.Image
	l.Status = container.Status
	l.State = container.State
	l.Time = rawLog.Time.Format("2006-01-02 15:04:05")

	return
}

//ReadLog open log file, start monitoring and send to logs channel
func ReadLog(container Docker, logsChan chan storage.DockerLog) {

	logger := log.GetInstance()
	//open json file from /var/lib/docker
	logFile := "/var/lib/docker/containers/" + container.Id + "/" + container.Id + "-json.log"
	log.WithPrefix(logger, "read docker log file "+logFile, "[INFO]")
	//get lines via tail
	t, _ := tail.TailFile(logFile, tail.Config{Follow: true, MustExist: true})
	//parse line
	for line := range t.Lines {
		var rawLog RawLog
		err := json.Unmarshal([]byte(line.Text), &rawLog)
		if err != nil {
			log.WithPrefix(logger, err.Error(), "[ERROR] readLog() ")
		}
		logsChan <- parseLog(rawLog, container)
	}
}
