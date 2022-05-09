# 阿里云WAF证书自动上传
该工具用于自动上传服务器申请的Let's Encrypt或其他证书。

推荐使用ACME脚本的hook功能，acme更新域名同时执行该脚本上传证书。
## 使用方式
### 下载二进制文件
在release里手动下载对应的二进制文件，或者下载源码编译
### 手动执行
```bash
# 参数说明
## 1. 阿里云accessKeyId，如果使用子用户请根据官方文档配置
## 2. 阿里云accessKeySecret
## 3. 域名
## 4. 证书绝对地址
## 5. 证书密钥绝对地址
## e.g.: ./AliyunLetsEncrypt LTAI5t*********TBXQtxPk SdanEjbj2**********u1sHH www.qcgzxw.cn www.qcgzxw.cn.cer www.qcgzxw.cn.key
./AliyunLetsEncrypt [accessKeyId] [accessKeySecret] [domain] [certificatePath] [privateKeyPath]
```
### 配置acme
```bash
# 方便脚本直接调用
cp AliyunLetsEncrypt /usr/local/bin/
# 编辑acme域名配置文件
vim ~/.acme.sh/[main domain]/[main domain].conf
# 更改Le_RenewHook字段为以下脚本，只需配置：阿里云accessKeyId、阿里云accessKeySecret和域名
AliyunLetsEncrypt LTAI5t*********TBXQtxPk SdanEjbj2**********u1sHH www.qcgzxw.cn $CERT_FULLCHAIN_PATH $CERT_KEY_PATH
```
### 子用户创建
![](https://cdn.jsdelivr.net/gh/image-backup/qcgzxw-images@master/image/16521030671371652103066697.png)
### 子用户添加权限
![](https://cdn.jsdelivr.net/gh/image-backup/qcgzxw-images@master/image/16521029601341652102959382.png)
![](https://cdn.jsdelivr.net/gh/image-backup/qcgzxw-images@master/image/16521030311341652103030317.png)

## 更新日志
### ver 0.1
- waf证书自动上传

## todo
- [x] 阿里云WAF
- [ ] 阿里云CDN 
- [ ] 打印日志
- [ ] 优化命令行模式
- [ ] 配置文件模式

## 参考
https://github.com/idawnlight/qcloud-lets-encrypt