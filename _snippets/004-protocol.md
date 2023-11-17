---
title: Simplex protocol
---

```c
#define MAX_PKT 1024

#define MAX_SEQ 4096

typedef unsigned int seq_nr;

#define inc(k) if(k<MAX_SEQ) k++; else k=0;

typedef struct {
  unsigned char data[MAX_PKT];
} packet;

typedef enum {
  data, ack, nack
} frame_kind;

typedef struct {
  frame_kind kind;
  seq_nr seq;
  seq_nr ack;
  packet data;
} frame;

typedef enum {
  frame_arrival
} event_type;

void wait_for_event(event_type *event);

void from_network_layer(packet *p);

void to_network_layer(packet *p);

void from_physical_layer(frame *r);

void to_physical_layer(frame *s);

void start_timer(seq_nr k);

void stop_timer(seq_nr k);

void start_ack_timer(void);

void stop_ack_timer(void);

void enable_network_layer(void);

void disable_network_layer(void);
```

```c
#include <stdbool.h>
#include "protocol.h"

void sender1(void) {
  frame s;
  packet buffer;

  while (true) {
    from_network_layer(&buffer);
    s.data = buffer;
    to_physical_layer(&s);
  }
}

void receiver1(void) {
  frame r;
  event_type event;

  while (true) {
    wait_for_event(&event);
    from_physical_layer(&r);
    to_network_layer(&r.data);
  }
}
```

```c
#include <stdbool.h>
#include "protocol.h"

void sender2(void) {
  frame s;
  packet buffer;
  event_type event;

  while(true) {
    from_network_layer(&buffer);
    s.data = buffer;
    to_physical_layer(&s);
    wait_for_event(&event);
  }
}

void receiver2(void) {
  frame r,s;
  event_type event;

  while(true) {
    wait_for_event(&event);
    from_physical_layer(&r);
    to_network_layer(&r.data);
    to_physical_layer(&s);
  }
}
```
