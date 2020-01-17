
## 如何使用 How to use it

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
yarn install && yarn run build:prod
docker-compose -f docker-compose.debug.yml  up -d
```

### 安装运行
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

## 运行Docker版本
```bush
前端
docker build -t shining .
docker run -d --name shining -p8099:80 shining
后端
docker build -t appdlserver -f appdlserver/Dockerfile ./appdlserver
docker run -d --name shining -p3000:3000 appdlserver
```