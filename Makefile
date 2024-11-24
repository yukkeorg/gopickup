PROGNAME := pickup
CMD_DIR := ./cmd/${PROGNAME}
BUILD_DIR := ./build

.PHONY: all

all: $(BUILD_DIR)/$(PROGNAME)

$(BUILD_DIR)/$(PROGNAME):
	go build -o $@ $(CMD_DIR)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
