# 网站状态监控
监控指定网站是否可正常访问

## 使用方法
1. 将需要检测的网址以一个一行的方式放入 data/urls.txt 文件中；
2. 配置 config/config.json 文件，确定输出报表的格式；
3. 启动程序
```bash
./server
```
你可以将程序加入定时任务来按时检测，每一次检测都将在 data/reports 目录下生成一个以时间为文件名的报表文件。

## 配置文件说明
debug: 是否开启调试模式，可选值为 true|false
reportFormat: 报表输出格式，可选值为 txt|csv