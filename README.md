# shining
- 🌟「闪灵」🌟
- 一个 App 在 workflow 的管理工具🚀

### 运行环境
* docker
* golang v1.13.5+
* npm v8.15+
* postgresql 12+
* https证书（请自行申请免费的dv证书）
> 如果缺少https证书，则无法安装iOS的安装包  

### Docker 启动
```bush
docker build -t shining .
docker run -d --name shining -p8099:80 shining
```  
```bush
docker build -t appdlserver -f appdlserver/Dockerfile ./appdlserver
docker run -d --name appdlserver -p3000:3000 appdlserver
```

### Docker Compose 启动
修改.env文件中的配置

修改src/config/index.ts中的prod域名地址

运行：
```bash
docker-compose -f docker-compose.yml up -d
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

### 手动打包
```bush
前端
npm run build:prod / yarn run build:prod
后端
cd appdlserver && go build -tags=jsoniter -o appdlserver
```

## 接口说明文档
后端使用3000端口，需要修改可以更改main.go文件中的r.Run(":3000")
* swagger路径/appFile/swagger/index.html
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
ProjectId=app
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