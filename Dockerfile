FROM golang:alpine as builder

# 为镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct" \
	TZ="Asia/Shanghai"

# 定义$PROJECT_DIR工作目录，项目代码复制到这个目录
ENV PROJECT_DIR="/home/go"
WORKDIR $PROJECT_DIR

# 将代码复制到容器中
COPY . .
RUN go mod download

#安装swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
#创建swagger文件
RUN swag init

# 将代码编译成二进制可执行文件,可执行文件名为 app
RUN go mod vendor
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o app .

# 移动到用于存放生成的二进制文件的/dist目录
WORKDIR /dist

# 将二进制文件从$PROJECT_DIR目录复制到这里
RUN cp $PROJECT_DIR/app .
# 拷贝文件
RUN cp -rf $PROJECT_DIR/.env .
# 拷贝静态资源文件
RUN cp -r $PROJECT_DIR/static .
RUN cp -r $PROJECT_DIR/view .
RUN cp -r $PROJECT_DIR/docs .

# 运行阶段
FROM alpine

# 时区设置
RUN apk add --no-cache tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

# 拷贝所有文件
COPY --from=builder /dist/ .
# 声明服务端口
EXPOSE 8089

# 容器启动时运行的命令
ENTRYPOINT ["./app"]