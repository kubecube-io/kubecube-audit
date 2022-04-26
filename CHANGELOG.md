# v1.2.0
2022-4-26
##Enhance:
- replace kubecube version v1.1.0 to v1.2.0
- replace jwt version v3.2.0 to v3.2.1
## Bugfix:
- 修复查询审计日志的分页问题

# v1.1.0
2021-12-16
##Feature: 
- add apis which receive generic and webconsole audit logs

# v1.0.1
2021-9-3
## Bugfix:
- 修改开关逻辑：hotplug中审计功能为开且用户配置或内置ES时，审计功能为开，否则为关

# v1.0.0
## Features:
- 接收KubeCube和K8s的审计信息，并发送到ES
- 查询和下载审计信息