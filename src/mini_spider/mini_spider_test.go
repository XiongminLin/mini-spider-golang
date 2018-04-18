/* mini_spider_test.go: test for mini_spider.go */
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
    "testing"
)

import (
    "mini_spider_config"
)

func TestNewMiniSpider(t *testing.T) {
    conf, _ := mini_spider_config.LoadConfig("../mini_spider_config/test_data/spider.conf")
    seeds, _ := LoadSeedFile(conf.Basic.UrlListFile)
    _, err := NewMiniSpider(&conf, seeds)
    if err != nil {
        t.Errorf("err happen in NewMiniSpider:%s", err.Error())
    }
}
