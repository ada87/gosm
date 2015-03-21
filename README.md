# gosm
学习中的项目 

##项目说明
想做一个系统工具，没有窗口风格的GUI.但通过WebSocket和HTTP服务做一个html版的UI。简单方便跨平台。

暂时想实现：

1. 系统信息速查
2. HTTP发包工具
3. window下实现tail -f
4. 正则表达式小工具
5. 加密工具盒子
6. 压力测试发送器
7. 爬虫工具

99. SQL功能设置

## 使用到的第三方包

目前用到的有两个

* code.google.com/p/go.net/websocket
 websocket官方扩展
* github.com/mattn/go-sqlite3
 嵌入式数据库sqlite3驱动，目前sqlite3唯一支持支持db/sql的驱动