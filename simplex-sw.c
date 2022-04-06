/*
 * ---------------------------------------------------------------------------
 * Simplex - Stop and Wait
 * ---------------------------------------------------------------------------
 * Also provides for a one-directional flow of data from sender to receiver.
 * The communication channel is once again assumed to be assumed error free.
 * However, the receiver has only a finite buffer capacity and a finite
 * processing speed, so the protocol must expliclly prevent the sender from
 * flooding the receiver wiLh data faster than it can be handled.
 * ---------------------------------------------------------------------------
 */

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