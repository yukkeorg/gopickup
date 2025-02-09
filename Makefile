PROGNAME := pickup
CMD_DIR := ./cmd/${PROGNAME}
DIST_DIR := ./dist

.PHONY: all

all: $(DIST_DIR)/$(PROGNAME)

$(DIST_DIR)/$(PROGNAME):
	go build -o $@ $(CMD_DIR)

.PHONY: clean
clean:
	rm -rf $(DIST_DIR)
