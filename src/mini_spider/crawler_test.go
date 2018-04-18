/* crawler_test.go: test for crawler */
/*
modification history
--------------------
2017/07/21, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/

package mini_spider

import (
    "sync"
    "testing"
)

import (
    "mini_spider_config"
)

func TestCrawler(t *testing.T) {
    urlTable := NewUrlTable()
    conf, _ := mini_spider_config.LoadConfig("../mini_spider_config/test_data/spider.conf")
    
    var queue Queue
    queue.Init()
    queue.Add(&CrawlTask{"http://pycm.baidu.com:8081", 2, nil})

    c := NewCrawler(urlTable, &conf, &queue)

    // wait and check result
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        c.Run()
        wg.Done()
    }()
    wg.Wait()

    // check visit result
    verifyLinks := []string{
        "http://pycm.baidu.com:8081/page1.html",
        "http://pycm.baidu.com:8081/page2.html",
        "http://pycm.baidu.com:8081/page3.html",
        "http://pycm.baidu.com:8081/mirror/index.html",
        "http://pycm.baidu.com:8081/page4.html",
    }

    for _, link := range verifyLinks {
        if !c.urlTable.Exist(link) {
            t.Errorf("%s not visited", link)
        }
    }
}
