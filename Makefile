BUILD_DIR := build
IMAGE_DIR := images

BOOT_SRC := src/boot/boot.asm
BOOT_BIN := $(BUILD_DIR)/boot.bin
BOOT_IMG := $(IMAGE_DIR)/os.img

KERNEL_SRC := src/kernel/kernel.c
KERNEL_OBJ := $(BUILD_DIR)/kernel.o

NASM ?= nasm
CC ?= gcc
QEMU ?= qemu-system-i386

CFLAGS := -ffreestanding -m32 -fno-pic -fno-stack-protector -nostdlib -Wall -Wextra

.PHONY: all boot kernel run clean

all: boot kernel

boot: $(BOOT_IMG)

kernel: $(KERNEL_OBJ)

$(BUILD_DIR) $(IMAGE_DIR):
	mkdir -p $@

$(BOOT_BIN): $(BOOT_SRC) | $(BUILD_DIR)
	$(NASM) $< -f bin -o $@

$(BOOT_IMG): $(BOOT_BIN) | $(IMAGE_DIR)
	cp $< $@

$(KERNEL_OBJ): $(KERNEL_SRC) | $(BUILD_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

run: $(BOOT_IMG)
	$(QEMU) -fda $(BOOT_IMG)

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(BOOT_IMG)

