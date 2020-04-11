//sha-net-mwg.ubisoft.org:3128

/*
	_auther = sun hai lang
	_time	= 2020-04-11
*/

package main

import (
	"fmt"
	"time"
)

func init(){

}

func main() {

	fmt.Println("Hello, 世界")
	// time包格式化日期，必须得是这个日期与时间
	fmt.Println("Time: " + time.Now().Format("2006-01-02 15:04:05"))
}

