CC = gcc
CFLAGS = -Wall -Werror -fPIC -shared
DIRECTIVE = /usr/local/src/rootfs_prefix

# rootfs_prefix.so creates a shared object file for rootfs_prefix.c
# which is preloaded with LD_PRELOAD
rootfs_prefix.so: rootfs_prefix.c
	$(CC) -I$(DIRECTIVE) $(CFLAGS) -o rootfs_prefix.so rootfs_prefix.c -ldl
