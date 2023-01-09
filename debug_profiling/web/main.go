package main

import (
	"bytes"
	"encoding/json"
	"github.com/xuegl/go-in-action/util"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var shanghaiTZ *time.Location

func main() {
	go func() {
		for {
			loadTimezone()
			processJson([]byte(`{"a": 1, "b": 2, "c": 3}`))
			processString([]byte("hello world"))
		}
	}()
	log.Println("server is listening on :8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
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
