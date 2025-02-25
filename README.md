# ApiSixKafkaLoggerToSecScan
通过Apisix_kafkalogger 插件收集流量进行安全扫描（如配置下级代理到 Xray、Burp Suite），实现实时流量监控和漏洞扫描。   
*流程图：流量 -> APISix -> KafkaLogger 插件 -> Kafka -> ApiSixKafkaLoggerToSecScan -> 安全扫描工具*

| **步骤** | **组件**          | **描述**                                                                 |
|----------|-------------------|--------------------------------------------------------------------------|
| 1        | 流量              | 外部请求或数据流进入系统。                                               |
| 2        | APISix            | 流量进入API网关，进行初步处理和转发。                                    |
| 3        | KafkaLogger 插件  | APISix通过KafkaLogger插件将日志或数据发送到Kafka。                       |
| 4        | Kafka             | Kafka作为消息中间件，接收并存储来自APISix的日志或数据。                  |
| 5        | ApiSixKafkaLoggerToSecScan             | ApiSixKafkaLoggerToSecScan读取kafka的数据解析成http/https流量，并将流量下级代理到安全扫描工具。                  |
| 6        | 安全扫描工具      | 从Kafka读取数据，并进行安全扫描分析。                                    |
## 适用场景
| **场景**                   | **自动化扫描工具**                  | **人工测试**                    |
|----------------------------|------------------------------------|---------------------------------|
| **大规模流量实时监控**         | ✅ 高效处理，7x24 实时扫描。            | ❌ 人力不足，无法处理实时流量。       |
| **较高测试覆盖率**         | ✅ 通过kafkalogger插件收集流量，覆盖率100%。           | ❌ 人工测试依靠点击功能点抓包来测试接口，难以保证较高覆盖率。      |
| **常见漏洞检测（XSS、SQL 注入等）** | ✅ 内置规则库，快速发现常见漏洞。         | ✅ 也可发现，但效率较低。             |
| **复杂业务逻辑漏洞**            | ❌ 工具难以检测，需要人工调整。            | ✅ 可灵活应对，适合复杂测试。         |
| **定期安全审计**              | ✅ 可快速完成大范围扫描，降低重复工作量。     | ✅ 适合深入分析特定模块的高风险漏洞。 |
| **低成本长期监控**            | ✅ 工具部署后维护成本低。                 | ❌ 人工测试成本较高，不适合长期监控。  |

## 用法
1. Apisix-kafkalogger插件配置，配置方法可参考[这里](https://blog.csdn.net/weixin_45945976/article/details/139123020?spm=1001.2014.3001.5501)，*记得修改如下配置文件中的kafka地址和topic；
```
{
  "_meta": {
    "disable": true
  },
  "batch_max_size": 1,
  "brokers": [
    {
      "host": "1.2.3.4",
      "port": 9092
    }
  ],
  "buffer_duration": 60,
  "cluster_name": 1,
  "inactive_timeout": 5,
  "include_req_body": true,
  "include_resp_body": true,
  "kafka_topic": "secapisixtest",
  "max_retry_count": 0,
  "meta_format": "default",
  "name": "kafka logger",
  "producer_batch_num": 1,
  "producer_batch_size": 1048576,
  "producer_max_buffering": 50000,
  "producer_time_linger": 1,
  "producer_type": "async",
  "required_acks": 1,
  "retry_delay": 1,
  "timeout": 3
}

```
2. 修改`main.go`中`Brokers`值为kafka地址（即与kafkalogger插件中的kafka地址保持一致），修改`topic`的值为kafka的topic；
3. 修改`secscan.go`中的`proxyURL, _ := url.Parse("http://127.0.0.1:3234")`里的下级代理地址为你的安全扫描器地址（如Xray或Burp的代理地址）；
4. 编译并部署扫描（可在`tools.go`文件中的`isWithinWorkingHours`中配置扫描的时间，如只在`00:00 - 09:00` 进行扫描）.
