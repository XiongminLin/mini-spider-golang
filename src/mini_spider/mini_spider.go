/* mini_spider.go */
/*
modification history
--------------------
2017/07/20, by Xiongmin LIN, create
*/
/*
DESCRIPTION
- mini spider
*/

package mini_spider

import (
	"mini_spider_config"
)

type MiniSpider struct {
	config   *mini_spider_config.MiniSpiderConf
	urlTable *UrlTable
	queue    Queue
	crawlers []*Crawler
}

// crawl task
type CrawlTask struct {
	Url    string             // url to crawl
	Depth  int                // depth of the url
	Header map[string]string  // http header
}

// create new mini-spider
func NewMiniSpider(conf *mini_spider_config.MiniSpiderConf, seeds []string) (*MiniSpider, error) {
	ms := new(MiniSpider)
	ms.config = conf

	// create url table
	ms.urlTable = NewUrlTable()

	// initialize queue
	ms.queue.Init()

	// add seeds to queue
	for _, seed := range seeds {
		task := &CrawlTask{Url: seed, Depth: 1, Header: make(map[string]string)}
		ms.queue.Add(task)
	}

	// create crawlers, thread count was defined in conf
	ms.crawlers = make([]*Crawler, 0)
	for i := 0; i < conf.Basic.ThreadCount; i++ {
		crawler := NewCrawler(ms.urlTable, ms.config, &ms.queue)
		ms.crawlers = append(ms.crawlers, crawler)
	}

	return ms, nil
}

// run mini spider
func (ms *MiniSpider) Run() {
	// start all crawlers
	for _, crawler := range ms.crawlers {
		go crawler.Run()
	}
}

// get number of unfinished task
func (ms *MiniSpider) GetUnfinished() int {

	return ms.queue.GetUnfinished()
}