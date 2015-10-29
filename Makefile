BUILD_DIR = build
GOPATH = $(CURDIR)

all: $(BUILD_DIR)/run

$(BUILD_DIR)/run: src/*.go pkg/linux_amd64/hitori.a
	go build -o $(BUILD_DIR)/run src/*.go && chmod +x $(BUILD_DIR)/run

run: $(BUILD_DIR)/run
	$(BUILD_DIR)/run puzzles/8.hit

pkg/linux_amd64/hitori.a: src/hitori/*.go
	go install hitori
