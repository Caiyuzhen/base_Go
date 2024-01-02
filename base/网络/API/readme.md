## 热重载库
- 【第一步】首先需要将项目给```模块化!!```
  - ```go mod init <modulename>``````
    - 例如:```go mod init API```

- 【第二步】安装对应的库
  - ```go install github.com/cosmtrek/air@latest```

- 【第三步】build
  - 记得需要有入口文件 main 才能 build!!!
    - go build
- 【第四步】运行 air
  - 初始化
    - ```air init```
  - 运行热更新
    - - ```air```
  - 如果没有权限运行的话
    - chmod +x /Users/aic/base_Go/base/网络/API/tmp/main