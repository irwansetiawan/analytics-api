main_package_path = .
binary_name = analytics

.PHONY: build
build:
	go build -o=${binary_name} ${main_package_path}
