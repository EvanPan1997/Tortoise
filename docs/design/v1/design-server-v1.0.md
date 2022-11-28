###服务端-Server
1. 接收请求-Receive Request From Client
2. 请求报文解析-Analysis Content Of Request
3. 解析SQL-Analysis SQL Syntax
4. 执行SQL-Execute SQL
5. 存储引擎&索引引擎-Storage Engine&IndexEngine

---
###架构设计-Framework Design
1. 请求解析器-Request Analyzer
2. SQL解析器(标准SQL语法)-SQL Analyzer(Standard SQL Syntax)
3. 执行器-Executor
4. 长连接控制器-Long Socket Connection Controller
5. 数据存储结构-Data Storage Structure