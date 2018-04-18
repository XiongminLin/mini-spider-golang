/* crawler.go - crawler thread*/
/*
modification history
--------------------
2017/07/18, by Xiongmin LIN, create
2017/08/01, by Xiongmin LIN, modify
 - add crawlChild function
*/
/*
DESCRIPTION
*/
package mini_spider

import (
	"fmt"
	"regexp"
	"time"
)

import (
	"www.baidu.com/golang-lib/http_util"
	"www.baidu.com/golang-lib/log"
)

import (
	"mini_spider_config"
	"web_package"
)

type Crawler struct {
	urlTable   *UrlTable
	config     *mini_spider_config.BasicConfig
	queue      *Queue
	urlPattern *regexp.Regexp
	stop       bool
}

// create new crawler
func NewCrawler(urlTable *UrlTable, config *mini_spider_config.MiniSpiderConf, queue *Queue) *Crawler {
	c := new(Crawler)
	c.urlTable = urlTable
	c.config = &config.Basic
	c.queue = queue

	// TargetUrl has been checked in conf load
	c.urlPattern, _ = regexp.Compile(c.config.TargetUrl)

	c.stop = false

	return c
}

// start crawler
func (c *Crawler) Run() {
	for !c.stop {
		// get new task from queue
		task := c.queue.Pop()
		log.Logger.Debug("from queue: url=%s, depth=%d", task.Url, task.Depth)

		// read data from given task
		data, err := http_util.Read(task.Url, c.config.CrawlTimeout, task.Header)
		if err != nil {
			log.Logger.Error("http_util.Read(%s):%s", task.Url, err.Error())
			c.queue.FinishOneTask()
			continue
		}

		// save data to file
		if c.urlPattern.MatchString(task.Url) {
			err = web_package.SaveWebPage(c.config.OutputDirectory, task.Url, data)
			if err != nil {
				log.Logger.Error("web_package.SaveWebPage(%s):%s", task.Url, err.Error())
			} else {
				log.Logger.Debug("save to file: %s", task.Url)
			}
		}

		// add to url table
		c.urlTable.Add(task.Url)

		// continue crawling until max depth
		if task.Depth < c.config.MaxDepth {
			err = c.crawlChild(data, task)
			if(err != nil){
				log.Logger.Error("crawlChild(%s):%s in depth of %d", task.Url, err.Error(), task.Depth)
			}
		}

		// confirm to remove task from queue
		c.queue.FinishOneTask()

		// sleep for a while
	    time.Sleep(time.Duration(c.config.CrawlInterval) * time.Second)
	}
}

// stop crawler
func (c *Crawler) Stop() {
	c.stop = true
}

// crawl child url
func (c *Crawler) crawlChild(data []byte, task *CrawlTask ) error {
	// parse url from web page
	links, err := web_package.ParseWebPage(data, task.Url)
	if err != nil {
		return fmt.Errorf("web_package.ParseWebPage():%s", err.Error())
	}

	// add child task to queue
	for _, link := range links {
		// check whether url match the pattern, or url exists already
		if c.urlTable.Exist(link) {
			continue
		}

		taskNew := &CrawlTask{Url: link, Depth: task.Depth + 1, Header: make(map[string]string)}
		log.Logger.Debug("add to queue: url=%s, depth=%d", taskNew.Url, taskNew.Depth)
		c.queue.Add(taskNew)
	}

	return nil
}
