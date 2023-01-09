package main

import (
	"bytes"
	"encoding/json"
	"github.com/xuegl/go-in-action/util"
	"os"
	"runtime/pprof"
	"time"
)

var shanghaiTZ *time.Location

func main() {
	file, err := os.Create("./go-in-action.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(file)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	for i := 0; i < 100; i++ {
		loadTimezone()
		processJson([]byte(`{"a": 1, "b": 2, "c": 3}`))
		processString([]byte("hello world"))
	}

}

func processJson(s []byte) {
	var m map[string]any
	err := json.Unmarshal(s, &m)
	if err != nil {
		panic(err)
	}
}

func processString(b []byte) {
	sa := make([]string, 0, 100)
	buf := make([]byte, 0, 100*len(b))
	writer := bytes.NewBuffer(buf)
	for i := 0; i < 100; i++ {
		sa = append(sa, util.ByteSliceToString(b))
		writer.Write(b)
	}
}

func loadTimezone() *time.Location {
	if shanghaiTZ != nil {
		return shanghaiTZ
	}

	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil
	}
	shanghaiTZ = tz
	return shanghaiTZ
}
