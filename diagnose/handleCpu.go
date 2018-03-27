package tools

import (
	"os"
	"log"
	"fmt"
	"time"
	"runtime/pprof"
)

var logDirCPU = "cpulogs"

func init() {
	os.Mkdir(logDirCPU, 0666)
}

//SamplingCpu CPU采样
func SamplingCpu(fileName string, duration time.Duration) {
	f, err := os.Create(fmt.Sprintf(`%s/%s`, logDirCPU, fileName))
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	time.Sleep(duration)
}
