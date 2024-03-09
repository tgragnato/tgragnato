---
title: EBMBTreeBench
description: a naive measurement tool for ebmb trees
layout: default
lang: en
---

An elastic binary tree is a kind of binary search tree, it's used by haproxy. This is a simplistic and unprecise benchmark against the reference implementation.

The main goal of the code below is to measure the time taken to insert elements into the tree, list them, and search for them. The output is the number of inserted elements and the insertion, listing, and lookup times.

```c
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
```

## 100k lookups

> set xlabel 'Size' / set ylabel 'Time' / set logscale x

> plot '100000.csv' using 1:2 with linespoints title 'insertion time', '100000.csv' using 1:3 with linespoints title 'listing time', '100000.csv' using 1:4 with linespoints title 'lookup time'

![100k lookups](/images/2023-09-20-100000.png){:loading="lazy"}

```
1, 0.000004, 0.000001, 0.009432
10, 0.000005, 0.000001, 0.009847
100, 0.000014, 0.000000, 0.009955
1000, 0.000150, 0.000001, 0.009988
10000, 0.001525, 0.000001, 0.009946
100000, 0.016220, 0.000001, 0.010079
1000000, 0.166881, 0.000000, 0.010158
10000000, 1.784307, 0.000001, 0.010158
```

## 1m lookups

> set xlabel 'Size' / set ylabel 'Time' / set logscale x

> plot '1000000.csv' using 1:2 with linespoints title 'insertion time', '1000000.csv' using 1:3 with linespoints title 'listing time', '1000000.csv' using 1:4 with linespoints title 'lookup time'

![1m lookups](/images/2023-09-20-1000000.png){:loading="lazy"}

```
1, 0.000004, 0.000000, 0.096662
10, 0.000004, 0.000000, 0.101871
100, 0.000016, 0.000000, 0.101853
1000, 0.000151, 0.000001, 0.101973
10000, 0.001510, 0.000001, 0.102159
100000, 0.016231, 0.000000, 0.101562
1000000, 0.167010, 0.000001, 0.101705
10000000, 1.788271, 0.000000, 0.101617
```

## 10m lookups

> set xlabel 'Size' / set ylabel 'Time' / set logscale x

> plot '1000000.csv' using 1:2 with linespoints title 'insertion time', '1000000.csv' using 1:3 with linespoints title 'listing time', '1000000.csv' using 1:4 with linespoints title 'lookup time'

![10m lookups](/images/2023-09-20-10000000.png){:loading="lazy"}

```
1, 0.000005, 0.000000, 0.999802
10, 0.000004, 0.000001, 1.054061
100, 0.000017, 0.000001, 1.059000
1000, 0.000149, 0.000001, 1.053368
10000, 0.001530, 0.000000, 1.053350
100000, 0.016166, 0.000000, 1.054473
1000000, 0.166407, 0.000000, 1.059266
10000000, 1.783292, 0.000000, 1.032653
```

## References

[https://wtarreau.blogspot.com/2011/12/elastic-binary-trees-ebtree.html](https://wtarreau.blogspot.com/2011/12/elastic-binary-trees-ebtree.html)

[https://github.com/wtarreau/ebtree](https://github.com/wtarreau/ebtree)

[https://github.com/haproxy/haproxy/blob/master/src/ebmbtree.c](https://github.com/haproxy/haproxy/blob/master/src/ebmbtree.c)
