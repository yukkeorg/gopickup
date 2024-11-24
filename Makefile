PROGNAME := pickup
CMD_DIR := ./cmd/${PROGNAME}
BUILD_DIR := ./build

.PHONY: all

all: build

build: $(BUILD_DIR)/$(PROGNAME)
	go build -o $(BUILD_DIR)/$(PROGNAME) $(CMD_DIR)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
