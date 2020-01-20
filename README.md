# shining
- 🌟「闪灵」🌟
- 一个 App 在 workflow 的管理工具🚀

## 运行环境
* docker
* golang v1.13.5+
* npm v8.15+
* postgresql 12+
* https证书（请自行申请免费的dv证书）
> 如果缺少https证书，则无法安装iOS的安装包  

## 运行方式
1. Docker 运行  
```bush
docker build -t shining .
docker run -d --name shining -p8099:80 shining
```  
```bush
docker build -t appdlserver -f appdlserver/Dockerfile ./appdlserver
docker run -d --name appdlserver -p3000:3000 appdlserver
```

2. Docker Compose 启动  
修改.env文件中的配置
修改src/config/index.ts中的prod域名地址
```bash
docker-compose -f docker-compose.yml up -d
```

## 编译
```bush
前端
npm run build:prod / yarn run build:prod
后端
cd appdlserver && go build -tags=jsoniter -o appdlserver
```

## 接口说明文档
https://github.com/lm93129/shining/wiki/API-doc