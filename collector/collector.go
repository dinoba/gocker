package collector

import (
	"time"

	"github.com/dinoba/gocker/dockerlog"
	"github.com/dinoba/gocker/log"
	"github.com/dinoba/gocker/storage"
	"github.com/dinoba/gocker/tools"
)

var _hostIP, _hostname string
var _dockers []dockerlog.Docker
var _dockersIDs []string
var logsChanDocker chan storage.DockerLog

//Collector storage def and log paths
type Collector struct {
	storageHandler storage.Handler
	dockersToSkip  string
}

//NewCollector creates collector
func NewCollector(st storage.Handler, dockersToSkip string) *Collector {
	return &Collector{
		storageHandler: st,
		dockersToSkip:  dockersToSkip,
	}
}

//CollectStats desc
func (st *Collector) CollectStats() (err error) {
	logsChanDocker = make(chan storage.DockerLog)

	logger := log.GetInstance()
	_hostIP = tools.GetIP()
	_hostname = tools.GetHostname()

	//listen docker channel
	go func() {
		for {
			select {
			case logItem := <-logsChanDocker:
				err := sendToStorage(logItem, st)
				if err != nil {
					log.WithPrefix(logger, "Can't store docker log "+err.Error(), "[ERROR]")
				}
			default:
				time.Sleep(time.Millisecond)
			}
		}
	}()

	log.WithPrefix(logger, "Start log collecting on "+_hostIP, "[INFO]")

	err = startDockerMonitoring(st.dockersToSkip)

	if err != nil {
		log.WithPrefix(logger, "startDockerMonitoring() "+err.Error(), "[ERROR]")
	}

	return
}

func sendToStorage(item storage.DockerLog, st *Collector) (err error) {
	item.HostIP = _hostIP
	item.HostName = _hostname

	err = st.storageHandler.Store(item)

	return
}

func startDockerMonitoring(dockersToSkip string) (err error) {
	_dockers, err = dockerlog.GetAllDockers(dockersToSkip)
	if err != nil {
		return
	}

	for _, v := range _dockers {
		_dockersIDs = append(_dockersIDs, v.Id)
		go func(docker dockerlog.Docker) {
			dockerlog.ReadLog(docker, logsChanDocker)
		}(v)

	}

	ticker := time.NewTicker(time.Second * 90)
	go func() {
		for _ = range ticker.C {
			dockerlog.CheckForNewDockers(logsChanDocker, &_dockersIDs, dockersToSkip)
		}
	}()

	return
}
