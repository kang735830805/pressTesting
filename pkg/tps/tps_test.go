package tps

import (
	"fmt"
	"testing"
)
//tps -l 1 -c 1 -n fact -m save - a "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}" -s
//	flags.IntVarP(&loop, "loop", "l", 1, "合约执行循环数量 eg. 1000")
//	flags.IntVarP(&concurrency, "concurrency", "c", 1, "并发数量. eg. 10")
//	flags.StringVarP(&name, "name", "n", "", "合约名称")
//	flags.StringVarP(&method, "method", "m", "", "合约内的方法")
//	flags.StringVarP(&args, "args", "a", "", "合约参数")
//	flags.StringVarP(&sdkPath, "sdkPath", "s", "", "SdkConfig路径")


//var loop int
//var concurrency int
//var name string
//var method string
//var args string
//var sdkPath string


func TestRunTps(t *testing.T) {
	loop, concurrency, threadNum, name, method, parameter, sdkPath = 1, 5,1, "fact", "save", "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}", "../../sdk_config2.yml"
	//loop, concurrency, threadNum, name, method, parameter, sdkPath = 1000, 1,1000, "fact", "save", "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}", "../../sdk_config2.yml"
	//
	//fmt.Println(loop)
	//fmt.Println(concurrency)
	//fmt.Println(name)
	//fmt.Println(method)
	//fmt.Println(args)
	//fmt.Println(sdkPath)
	//dir, _ := os.ReadDir("../../")
	//for _, file := range dir {
	//	fmt.Println(file.Name())
	//}

	e := RunTps()
	fmt.Println(e)
	//hex.DecodeString()


}

//func StringToBytes(s string) []byte {
//	return *(*[]byte)(unsafe.Pointer(
//		&struct {
//			string
//			Cap int
//		}{s, len(s)},
//	))
//}