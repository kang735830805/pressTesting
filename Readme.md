
##TPS
|  longname   |  shortname   | description  | examples  |
|  ----  | ----  | ----  | ----  |
| loop  | l | 交易执行总数量 | 1000 |
| threadNum  | t | 进程数量 | 1000 |
| concurrency  | c | 单一进程内处理交易数量 | 100 |
| name  | n | 合约名称 | test |
| method  | m | 合约内方法 | save |
| parameter  | p | 合约参数 | "{\"file_name\":\"name007\"}" |
| sdkPath  | s | SdkConfig路径 | ./sdk_config.yml |

####tps命令行Examples：
```shell
./main tps -l 10000 -c 70 -t 3000 -n fact -m save -p "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"65432ç34\"}" -s ./sdk_config.yml
```


##QPS
|  longname   |  shortname   | description  | examples  |
|  ----  | ----  | ----  | ----  |
| loop  | l | 交易执行总数量 | 1000 |
| threadNum  | t | 进程数量 | 1000 |
| concurrency  | c | 单一进程内处理交易数量 | 100 |
| txId  | t | 长安链内的交易txId | 1705f56583e6cd78ca18426f40000b933b8842529c9346b49f4bbcae4b57a57e |
| sdkPath  | s | SdkConfig路径 | ./sdk_config.yml |

####qps命令行Examples：
```shell
#qps 测试
./main qps -l 150000 -c 1000 -t 50000 -i '1705f56583e6cd78ca18426f40000b933b8842529c9346b49f4bbcae4b57a57e' -s "./sdk_config.yml"
```
