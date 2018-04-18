/* conf_basic.go - basic config for mini spider */
/*
modification history
--------------------
2017/07/20, by linxiongmin, create
*/
/*
DESCRIPTION
*/

package mini_spider_config

import (
	"fmt"
	"regexp"
)

type BasicConfig struct {
	UrlListFile             string
	OutputDirectory         string
	MaxDepth                int
	CrawlInterval           int
	CrawlTimeout            int
	TargetUrl               string
	ThreadCount             int
	GracefulShutdownTimeout int
}

func (conf *BasicConfig) Check() error {
	// check urlListFile
	if conf.UrlListFile == "" {
		return fmt.Errorf("UrlListFile not set\n")
	}

	// check OutputDirectory
	if conf.OutputDirectory == "" {
		return fmt.Errorf("OutputDirectory not set\n")
	}

	// check MaxDepth
	if conf.MaxDepth < 1 {
		return fmt.Errorf("MaxDepth should >= 1\n")
	}

	// check CrawlInterval
	if conf.CrawlInterval < 1 {
		return fmt.Errorf("CrawlInterval should >= 1\n")
	}

	// check CrawlTimeout
	if conf.CrawlTimeout < 1 {
		return fmt.Errorf("CrawlTimeout should >= 1\n")
	}

	// check TargetUrl
	_, err := regexp.Compile(conf.TargetUrl)
	if err != nil {
		return fmt.Errorf("regexp.Compile(TargetUrl):%s", err.Error())
	}

	// check ThreadCount
	if conf.ThreadCount < 1 {
		return fmt.Errorf("ThreadCount should >= 1\n")
	}

	// check graceful shutdown timeout
	if conf.GracefulShutdownTimeout < 1 || conf.GracefulShutdownTimeout > 60{
		return fmt.Errorf("GracefulShutdownTimeout out of range [1, 60]\n")
	}

	return nil
}
