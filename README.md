# anywhere-go
Anywhere任我行——新一代互联网辅助连接框架

开发文档：
1. client连接时server分配给每个client一个独特的ID(类似于teamviewer)
2. server储存下发的ID以供client相互连接
3. server给出调用ID的接口(输入ID返回通道)
4. client通过ID与另一个client进行连接


core.Dial 返回conn conn负责通信

Write:
'''
payload := "commander request."
nBytes, err := conn.Write([]byte(payload))
'''

Read:
'''
response := make([]byte, 1024)
nBytes, err = conn.Read(response)
'''
