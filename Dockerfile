FROM golang:alpine as builder

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct" \
	TZ="Asia/Shanghai"
	
# 定义$PROJECT_DIR工作目录，项目代码复制到这个目录
ENV PROJECT_DIR="/home/go/project"
WORKDIR $PROJECT_DIR

# 将代码复制到容器中
COPY . .
RUN go mod download
# 将代码编译成二进制可执行文件,可执行文件名为 app
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o app .

# 移动到用于存放生成的二进制文件的/dist目录
WORKDIR /dist

# 将二进制文件从$PROJECT_DIR目录复制到这里
RUN cp $PROJECT_DIR/app .
# 在容器目录 /dist 创建一个目录 为src
#RUN mkdir src .

# 拷贝静态资源文件
RUN cp -r $PROJECT_DIR/static ./
# 拷贝env文件
RUN cp -r $PROJECT_DIR/.env ./

# 设置Go程序权限
RUN chmod +x /dist/app

# 运行阶段
FROM scratch
#COPY --from=builder /dist/app /dist/app

# 声明服务端口
EXPOSE 8088

# 容器启动时运行的命令
ENTRYPOINT ["/dist/app"]