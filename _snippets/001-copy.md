---
title: Copy
---

```c
#include <sys/types.h>
#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[]);

#define BUF_SIZE 4096
#define OUTPUT_MODE 0700

int main(int argc, char *argv[]) {

  int in_fd, out_fd, byte_letti, byte_scritti;
  char buffer[BUF_SIZE];

  if (argc != 3) exit(1);
  in_fd = open(argv[1], O_RDONLY);
  if(in_fd < 0) exit(2);
  out_fd = creat(argv[2], OUTPUT_MODE);
  if(out_fd < 0) exit(3);

  while(1) {
    byte_letti = read(in_fd, buffer, BUF_SIZE);
    if(byte_letti <= 0) break;
    byte_scritti = write(out_fd, buffer, byte_letti);
    if (byte_scritti <= 0) exit(4);
  }

  close(in_fd); close(out_fd);
  if (byte_letti == 0) exit(0);
  else exit(5);
}
```