# 设置输出目录
RELEASE_DIR := release

# 设置可执行文件名
BINARY_NAME := gostonc

# 创建输出目录
$(RELEASE_DIR):
	mkdir $(RELEASE_DIR)

# 编译目标
build: $(RELEASE_DIR)
	go build -o $(RELEASE_DIR)/$(BINARY_NAME).exe

# 清理目标
clean:
	rm -rf $(RELEASE_DIR)

# 默认目标
.PHONY: build clean
all: build