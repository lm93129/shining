# shining

- 🌟「闪灵」🌟
- 一个 App 在 workflow 的管理工具 🚀

## 运行环境

- docker
- golang v1.13.5+
- npm v8.15+
- postgresql 12+
- https 证书（请自行申请免费的 dv 证书）
  > 如果缺少 https 证书，则无法安装 iOS 的安装包

## 配置

- 修改.env 文件中的配置
- 修改 src/config/index.ts 中的 prod 域名地址

## 运行方式

> Docker 运行

```bush
docker build -t shining .
docker run -d --name shining -p8099:80 shining
```

```bush
docker build -t appdlserver -f appdlserver/Dockerfile ./appdlserver
docker run -d --name appdlserver -p3000:3000 appdlserver
```

> Docker Compose 运行

```bash
docker-compose -f docker-compose.yml up -d
```

## 编译

```bush
npm run build:prod / yarn run build:prod
cd appdlserver && go build -tags=jsoniter -o appdlserver
```

## 接口文档

https://github.com/lm93129/shining/wiki/API-doc
