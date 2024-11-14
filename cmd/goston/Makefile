# 默认目标
all: linux_build

# 构建 Linux 版本
linux_build:
	set GOOS=linux && set GOARCH=amd64 && go build -o myprogram

# 清理构建产物
clean:
	del myprogram

# 查看帮助信息
help:
	@echo "make linux_build - 构建 Linux 版本的程序"
	@echo "make clean      - 清理构建产物"
	@echo "make help       - 显示帮助信息"

.PHONY: all linux_build clean help