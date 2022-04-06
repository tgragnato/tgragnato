/*
 * ---------------------------------------------------------------------------
 * Simplex - No restrictions (utopia)
 * ---------------------------------------------------------------------------
 * Provides for data transmission in one direction only,
 * from sender to receiver.
 * The communication channel is assumed to be error free,
 * and the receiver is assumed to be able to process
 * all the input infinitely quickly.
 * Consequently, the sender just sits in a loop pumping data out onto
 * the line as last as il can.
 * ---------------------------------------------------------------------------
 */

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