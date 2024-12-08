# 定义要执行的命令
YAEGI_CMD := yaegi extract
SOURCE_DIR := goravel/app/models
TARGET_DIR := ./symbols

# 目标
all: check_yaegi extract_models move_file
# 检查是否已经安装 yaegi，如果没有则安装
check_yaegi:
	@if ! command -v yaegi &> /dev/null; then \
		echo "yaegi not found, installing..."; \
		go install github.com/traefik/yaegi/cmd/yaegi@v0.15.0; \
	else \
		echo "yaegi is already installed"; \
	fi

# 执行 yaegi extract 仅针对指定目录
extract_models:
	$(YAEGI_CMD) $(SOURCE_DIR)

# 移动生成的文件到目标目录，并修改 package 声明
move_file:
	@find . -maxdepth 1 -name 'goravel-app-models.go' -exec sed -i  's/package tinker/package symbols/' {} \;
	@mv goravel-app-models.go $(TARGET_DIR)/
	@echo "Moved and updated package name in goravel-app-models.go to 'symbols'"
.PHONY: all install_yaegi extract_models move_file
