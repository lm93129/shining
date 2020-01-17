# shining
- 🌟「闪灵」🌟
- 一个 App 在 workflwo 的管理工具🚀

### 如何启动 
```bush
前端
docker build -t shining .
docker run -d --name shining -p8099:80 shining
后端
docker build -t appdlserver -f appdlserver/Dockerfile ./appdlserver
docker run -d --name shining -p3000:3000 appdlserver
```  

### 运行环境
* golang v1.13.5+
* npm v8.15+
* nginx 1.17.0+
* postgresql 12+
* https证书（请自行申请免费的dv证书）

如果没得https证书，则无法安装iOS的安装包

### 快速启动
修改.env文件中的配置

修改src/config/index.ts中的prod域名地址

运行：
```bash
docker-compose -f docker-compose.debug.yml up -d
```

### 安装运行(需要有nodejs和golang的环境)
前端
```bush
npm i / yarn 
npm run dev / yarn run dev
```
后端
```bash
cd appdlserver && go run main.go
```

### 生产环境打包
```bush
npm run build:prod / yarn run prod:dev
```

## 接口说明文档
后端使用3000端口，需要修改可以更改main.go文件中的r.Run(":3000")

### 上传安装包

crul例子：
```bash
curl --request POST \
  --url http://127.0.0.1:3000/appFile/uploadFile \
  --header 'cache-control: no-cache' \
  --header 'content-type: multipart/form-data' \
  --form upload=@xxx.ipa \
  --form 'UpdateDescription=更新说明' \
  --form ProjectId=asoco-app
```

#### Request
- Method: **POST**
- URL:  ```/appFile/uploadFile```
- Headers：```'content-type: multipart/form-data'```
- Params: 
```
# 安装包所属项目
ProjectId=asoco-app
# 更新说明
UpdateDescription=更新说明
# 安装包文件
upload=@xxx.apk
```
#### Response
- Body
```
{
  "code": 200,
  "message": "OK"
}
```