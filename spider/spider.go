package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/yanyiwu/igo"
	"sync"
	"time"
)

const (
	SleepSeconds = 60
)

var wg sync.WaitGroup

func spiderRunner(url string) {
	defer wg.Done()
	for {
		content := igo.HttpGet(url)
		content = convert(content)
		msgs := Parse(content)
		if msgs == nil {
			glog.Error("Parse failed")
		} else {
			for _, item := range msgs {
				//glog.Info(item.GetTitle())
				//glog.Info(item.GetUrl())
				err := Insert("feeds", item.GetTitle(), item.GetContent(), item.GetUrl())
				if err == nil {
					glog.Info(item.GetTitle(), item.GetUrl())
				} else {
					glog.V(2).Info(err)
				}
			}
		}
		glog.V(3).Info("time.Sleep ", SleepSeconds, " seconds")
		time.Sleep(SleepSeconds * time.Second)
	}
}

func main() {
	flag.Parse()
	kwDisp := NewKeywordDispatcher()
	kwDisp.Insert("实习")
	Connect("127.0.0.1", "inforyoumation")
	iter, err := igo.NewLineIterator("urls")
	if err != nil {
		glog.Fatal(err)
	}
	for iter.HasNext() {
		url := iter.Next()
		wg.Add(1)
		go spiderRunner(url)
		glog.Info(url)
	}
	wg.Wait()
}