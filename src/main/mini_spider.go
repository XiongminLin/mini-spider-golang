/* mini_spider.go - program entry point */
/*
modification history
--------------------
2017/07/20, by Xiongmin LIN, create
*/
/*
DESCRIPTION
mini spider
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


import (
	"code.google.com/p/log4go"
	"www.baidu.com/golang-lib/log"
)

import (
	"mini_spider_config"
	"mini_spider"
)

var (
	confPath    *string = flag.String("c", "../conf/spider.conf", "mini_spider configure path")
	help        *bool   = flag.Bool("h", false, "show help")
	logPath     *string = flag.String("l", "../log", "dir path of log")
	showVer     *bool   = flag.Bool("v", false, "show version")
	stdOut      *bool   = flag.Bool("s", false, "to show log in stdout")
	debugLog    *bool   = flag.Bool("d", false, "to show debug log (otherwise >= info)")
)

func Exit(code int) {
	log.Logger.Close()
	/* to overcome bug in log, sleep for a while    */
	time.Sleep(1 * time.Second)
	os.Exit(code)
}

/* the main function */
func main() {
	var logSwitch string

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}
	if *showVer {
		fmt.Printf("version is: 1.0.0\n")
		return
	}

	// debug switch
	if *debugLog {
		logSwitch = "DEBUG"
	} else {
		logSwitch = "INFO"
	}
	fmt.Printf("mini_spider starts...\n")

	/* initialize log   */
	/* set log buffer size  */
	log4go.SetLogBufferLength(10000)
	/* if blocking, log will be dropped */
	log4go.SetLogWithBlocking(false)
	/* we want to get state of log4go   */
	log4go.SetWithModuleState(true)

	err := log.Init("mini_spider", logSwitch, *logPath, *stdOut, "midnight", 5)
	if err != nil {
		fmt.Printf("main(): err in log.Init():%s\n", err.Error())
		os.Exit(-1)
	}

	// load config
	config, err := mini_spider_config.LoadConfig(*confPath)
	if err != nil {
		log.Logger.Error("main():err in ConfigLoad():%s", err.Error())
		Exit(-1)
	}

	// load seeds
	seeds, err := mini_spider.LoadSeedFile(config.Basic.UrlListFile)
	if err != nil {
		log.Logger.Error("main():err in loadSeedFile(%s):%s", config.Basic.UrlListFile, err.Error())
		Exit(1)
	}

	// create mini-spider
	miniSpider, err:= mini_spider.NewMiniSpider(&config, seeds)
	if err != nil {
		log.Logger.Error("main():err in NewMiniSpider():%s", err.Error())
		Exit(1)
	}

	// run mini-spider
	miniSpider.Run()

	// waiting for all tasks to finish. 
	go func() {
		for {
			if miniSpider.GetUnfinished() == 0 {
				log.Logger.Info("All task finished, quit")
				Exit(0)
			}

			log.Logger.Debug("Waiting for %d tasks to finish\n", miniSpider.GetUnfinished())

			// sleep for a while
			time.Sleep(5 * time.Second)
		}
	} ()
	
	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// ensure that all logs are export and normal exit
	Exit(0)

}
