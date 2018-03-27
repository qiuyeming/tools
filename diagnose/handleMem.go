package tools

import (
	"os"
	"log"
	"runtime/pprof"
	"fmt"
	"runtime"
)

var logDir = "memlogs"

func init() {
	os.Mkdir(logDir, 0666)
}

//SamplingMem 内存采样
func SamplingMem(fileName string) {
	f, err := os.Create(fmt.Sprintf(`%s/%s`, logDir, fileName))
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	f.Close()
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)
	if mem.Sys > (2 << 30) {
		os.Exit(0)
	}
}
