
## 如何使用 How to use it
### 安装运行
```bush
npm i / yarn 
npm run dev / yarn run dev
```

### 测试环境打包
```bush
npm run build:dev / yarn run build:dev
```

### 生产环境打包
```bush
npm run build:prod / yarn run prod:dev
```

## 运行Docker版本
```bush
docker build -t yukino .
docker run -d --name yukino -p8099:80 yukino
```