# kubecube-audit

kubecube-audit 为 KubeCube 提供审计服务，对 KubeCube 和 K8s 产生的审计日志进行记录和管理。

## 解决什么问题

它可以帮助集群管理员处理以下问题：

- 发生了什么？
- 什么时候发生的？
- 谁触发的？
- 活动发生在哪个（些）对象上？
- 在哪观察到的？
- 它从哪触发的？
- 活动的后续处理行为是什么？

## 使用说明

- kubecube-audit 服务在用户一键部署时也会同时部署，接收来自 KubeCube 和 Kubernetes 的审计日志。

- 该服务会将接收到的日志发送到配置的 Elasticsearch 地址（如果用户未配置 Elasticsearch 地址，发送到内部 Elasticsearch）。

- 审计功能默认开启，如果希望关闭审计功能，参考[操作审计使用文档](https://www.kubecube.io/docs/user-guide/administration/audit/)

### 审计查询和导出

![审计界面](./docs/审计界面.png)

#### 查询

提供审计查询功能，用户可以在页面查询产生的审计日志，由 kubecube-audit 对 Elasticsearch 中的日志进行查询。

查询功能仅限平台管理员，对其他角色用户不开放入口。

#### 导出

支持对查询到的审计结果进行导出，权限限制同上。

用户首先进行查询日志，选择导出时，默认根据刚刚查询的过滤参数重新查询，查询条数不大于为10000条。文件格式默认为 csv 格式。

## 开源协议

```
Copyright 2021 KubeCube Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```