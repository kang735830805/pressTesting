```shell
# tps 测试
#"loop", "l",  "合约执行循环数量 eg. 1000"
#"concurrency", "c", "并发数量. eg. 10"
#"name", "n",  "合约名称"
#"method", "m", "合约内的方法"
#"args", "a", "", "合约参数"
#"sdkPath", "s", "SdkConfig路径"
./main tps -l 1 -c 1 -t 1 -n fact -m save -a "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}" -s ./sdk_config.yml
./main tps -l 1000  -t 500 -n fact -m save -a "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}" -s "./sdk_config.yml,./sdk_config2.yml"
./main tps -l 1  -t 1 -n fact -m save -a "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"6543234\"}" -s "./sdk_config.yml,./sdk_config2.yml"
```

```shell
#cpts 测试

```


```shell
#qps 测试
./main qps -l 1 -t 1 -i '1705f26839db6648cad8cce77b12d6c3a36a55ef57b944d7becb5b5b98470350' -s "./sdk_config.yml"
./main qps -l 1000 -t 1000 -i '1705f26839db6648cad8cce77b12d6c3a36a55ef57b944d7becb5b5b98470350' -s "./sdk_config.yml,./sdk_config2.yml"
```
