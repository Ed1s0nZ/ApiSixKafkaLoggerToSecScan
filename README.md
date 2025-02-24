# ApiSixKafkaLoggerToSecScan
通过Apisix_kafkalogger 插件收集流量进行安全扫描（如配置下级代理到xray、Burp）的工具。
## 用法
Apisix-kafkalogger插件配置
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
