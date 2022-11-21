package main

import (
	"log"
	"net/http"
)

/*
问题描述：
	if中的代码块重新定义了名为client的变量，即最外层的client和if里面的client是不同的两个变量，当你
	在后面使用client的时候，最外层的client仍为nil
*/

func listing1() error {
	var client *http.Client
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	_ = client
	return nil
}

/*
解决方案1：
	可以通过将代码块中的变量重新命名为和外层的client不同名避免上述nil的问题
*/

func listing2() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		client = c
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	_ = client
	return nil
}

/*
解决方案2：
	提前将变量都声明好，缺点就是需要提前声明
*/

func listing3() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

/*
解决方案3：
	将err的处理从嵌套语句中提取出来，减少代码的嵌入，可以让代码的阅读更简单
*/
func listing4() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
