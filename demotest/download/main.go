package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

const durl = "http://10.63.72.13:9333/"

func main() {
	//s := readFile()
	//for _, fid := range s {
	//
	//	var u = durl + strings.Trim(fid, "\n")
	//	fmt.Println(u)
	//	woker("http://10.63.72.13:9333/252,de9036b6c64f24")
	//}
	woker("http://10.63.72.13:9333/252,de9036b6c64f24")
}

func readFile() []string {
	f, err := os.Open("肿瘤肺部片子.txt")
	var s []string
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		s = append(s, line)
	}
	return s
}

func woker(u string) {
	uri, err := url.ParseRequestURI(u)
	if err != nil {
		panic("网址错误")
	}

	filename := path.Base(uri.Path)
	log.Println("[*] Filename " + filename)

	client := http.DefaultClient
	client.Timeout = time.Second * 60 //设置超时时间
	resp, err := client.Get(u)
	if err != nil {
		panic(err)
	}
	if resp.ContentLength <= 0 {
		log.Println("[*] Destination server does not support breakpoint download.")
	}
	raw := resp.Body
	defer raw.Close()
	reader := bufio.NewReaderSize(raw, 1024*32)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	buff := make([]byte, 32*1024)
	written := 0
	go func() {
		for {
			nr, er := reader.Read(buff)
			if nr > 0 {
				nw, ew := writer.Write(buff[0:nr])
				if nw > 0 {
					written += nw
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortWrite
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		if err != nil {
			panic(err)
		}
	}()

	//spaceTime := time.Second * 1
	//ticker := time.NewTicker(spaceTime)
	//lastWtn := 0
	//stop := false
	//
	//for {
	//	select {
	//	case <-ticker.C:
	//		speed := written - lastWtn
	//		log.Printf("[*] Speed %s / %s \n", bytesToSize(speed), spaceTime.String())
	//		if written-lastWtn == 0 {
	//			ticker.Stop()
	//			stop = true
	//			break
	//		}
	//		lastWtn = written
	//	}
	//	if stop {
	//		break
	//	}
	//}
}

func bytesToSize(length int) string {
	var k = 1024 // or 1024
	var sizes = []string{"Bytes", "KB", "MB", "GB", "TB"}
	if length == 0 {
		return "0 Bytes"
	}
	i := math.Floor(math.Log(float64(length)) / math.Log(float64(k)))
	r := float64(length) / math.Pow(float64(k), i)
	return strconv.FormatFloat(r, 'f', 3, 64) + " " + sizes[int(i)]
}
