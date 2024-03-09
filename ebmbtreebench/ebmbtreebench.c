#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "ebsttree.h"

int main(int argc, char **argv) {
	int size;
	int loops;
	struct eb_root root = EB_ROOT;
	struct ebmb_node *node;
	struct ebmb_node *node_mb, *node_len, *node_len1;
	clock_t start_time;
	float insertion_time;
	float listing_time;
	float lookup_time;

	/* disable output buffering */
	setbuf(stdout, NULL);

	if (argc != 3) {
		fprintf(stderr, "Usage: %s size loops\n", argv[0]);
		exit(1);
	}

	size = atoi(argv[1]);
	loops = atoi(argv[2]);

	start_time = clock();
	for (int i = 0; i < size; i++) {
		char* str = malloc(65);
		snprintf(str, 65, "%d", i);
		node = calloc(1, sizeof(*node) + 65);
		memcpy(node->key, *argv, 65);
		ebst_insert(&root, node);
	}
	insertion_time = (float)(clock() - start_time) / CLOCKS_PER_SEC;

	start_time = clock();
	node = ebmb_first(&root);
	while (node) {
		node = ebmb_next(node);
	}
	listing_time = (float)(clock() - start_time) / CLOCKS_PER_SEC;

	start_time = clock();
	for (int i = 0; i < loops; i++) {
		char* str = malloc(65);
		snprintf(str, 65, "%d", i);
		node      = ebst_lookup(&root, str);
		node_mb   = ebmb_lookup(&root, str, 65);
		node_len  = ebst_lookup_len(&root, str, 65);
		node_len1 = ebst_lookup_len(&root, str, 64);
	}
	lookup_time = (float)(clock() - start_time) / CLOCKS_PER_SEC;

	printf("%d, %.6f, %.6f, %.6f\n", size, insertion_time, listing_time, lookup_time);

	return 0;
}
