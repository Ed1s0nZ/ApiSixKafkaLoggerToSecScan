# ApiSixKafkaLoggerToSecScan
通过Apisix_kafkalogger 插件收集流量进行安全扫描（如配置下级代理到 Xray、Burp Suite），实现实时流量监控和漏洞扫描。
## 与人工测试的对比

| **对比项**             | **自动化扫描工具（APISIX + Xray/Burp）**                                      | **人工测试**                                                                 |
|------------------------|----------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| **效率**               | 自动化工具可快速处理大量流量，7x24 不间断工作，覆盖范围广。                          | 人工测试需手动操作，速度慢，无法处理大规模流量，时间和人力成本高。                                     |
| **覆盖范围**            | 自动化工具可以扫描常见漏洞（SQL 注入、XSS、文件上传等），对已知漏洞有较高覆盖率。          | 人工测试可针对复杂逻辑漏洞和业务逻辑漏洞，发现自动化工具难以探测的问题。                                  |
| **实时性**             | 流量实时采集和扫描，可快速发现漏洞并反馈。                                       | 通常是定期进行渗透测试，无法实时监控漏洞风险。                                                   |
| **灵活性**             | 依赖于工具的规则库和能力，复杂场景（如多步骤验证、加密参数）可能无法处理。                  | 人工测试可根据具体业务场景灵活调整，适合处理复杂和动态交互的场景。                                   |
| **成本**               | 一次部署后，自动化工具的运行成本较低，适合长期使用，节省人力。                           | 需要专业安全人员，成本高，尤其是复杂系统需要更多人力投入。                                      |
| **误报率**             | 自动化工具可能会产生较多误报，需要额外筛选结果。                                      | 人工操作通常误报率较低，但容易遗漏规模化、重复性的问题。                                         |
| **学习曲线**            | 工具（如 APISIX、Xray/Burp）需要一定的配置和学习成本，但长期使用后维护较简单。                | 人工测试需要较高的安全专业技能，学习成本高，且需要持续更新知识。                                     |
| **定制化能力**          | 工具支持自定义规则，但灵活性有限，难以完全适配某些复杂业务场景。                          | 人工测试可以完全定制，适配复杂业务逻辑和动态交互场景。                                           |
## 用法
1、Apisix-kafkalogger插件配置，配置方法可参考[这里](https://blog.csdn.net/weixin_45945976/article/details/139123020?spm=1001.2014.3001.5501)，*记得修改如下配置文件中的kafka地址和topic；
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
2、修改`main.go`中`Brokers`值为kafka地址（即与kafkalogger插件中的kafka地址保持一致），修改`topic`的值为kafka的topic；
3、修改`secscan.go`中的`proxyURL, _ := url.Parse("http://127.0.0.1:3234")`里的下级代理地址为你的安全扫描器地址（如Xray或Burp的代理地址）；
4、编译并部署扫描（可在`tools.go`文件中的`isWithinWorkingHours`中配置扫描的时间，如只在`00:00 - 09:00` 进行扫描）；
