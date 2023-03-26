
// determines packet size in bytes
#define MAX_PKT 1024

// sequence of n° ACK
#define MAX_SEQ 4096
typedef unsigned int seq_nr;
#define inc(k) if(k<MAX_SEQ) k++; else k=0;

// packet
typedef struct {
  unsigned char data[MAX_PKT];
} packet;

// frame kind
typedef enum {
  data, ack, nack
} frame_kind;

/*
 *             FRAME
 *
 * #------#-----#-----#----------#
 * # kind # seq # ack #   data   #
 * #------#-----#-----#----------#
 *
*/
typedef struct {
  frame_kind kind;
  seq_nr seq;
  seq_nr ack;
  packet data;
} frame;

// event type
typedef enum {
  frame_arrival
} event_type;

// wait for an event to happen
// return its type in event
void wait_for_event(event_type *event);

// fetch a packet from the network layer
void from_network_layer(packet *p);

// deliver information to the network layer
void to_network_layer(packet *p);

// go get an inbound frame from the phisical layer
void from_physical_layer(frame *r);

// pass the frame to the physical layer
void to_physical_layer(frame *s);

// start the clock running
// enable timeout event
void start_timer(seq_nr k);

// strop the clock
// disable timeout event
void stop_timer(seq_nr k);

// start an auxiliary timer
// enable the ack_timeout event
void start_ack_timer(void);

// stop the auxiliary timer
// disable the ack_timeout event
void stop_ack_timer(void);

// allow a network_layer_ready event
void enable_network_layer(void);

// forbid a network_layer_ready event
void disable_network_layer(void);