```shell
# tps 测试
#"loop", "l",  "合约执行循环数量 eg. 1000"
#"concurrency", "c", "并发数量. eg. 10"
#"name", "n",  "合约名称"
#"method", "m", "合约内的方法"
#"args", "a", "", "合约参数"
#"sdkPath", "s", "SdkConfig路径"
./main tps -l 1 -c 1 -n fact -m save - a "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}" -s ./sdk_config.yml
```

```shell
#cpts 测试

```