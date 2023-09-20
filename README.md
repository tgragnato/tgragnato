# EBMBTreeBench

EBMBTreeBench is a naive measurement tool for ebmb trees.

Run sequentially `make 100000`, `make 1000000`, `make 10000000`

The size is the number of nodes in the tree.
You are testing 100000, 1000000, 10000000 lookups for each size of the tree.

```
set xlabel 'Size'
set ylabel 'Time'
set logscale x

plot '100000.csv' using 1:2 with linespoints title 'insertion time', '100000.csv' using 1:3 with linespoints title 'listing time', '100000.csv' using 1:4 with linespoints title 'lookup time'

plot '1000000.csv' using 1:2 with linespoints title 'insertion time', '1000000.csv' using 1:3 with linespoints title 'listing time', '1000000.csv' using 1:4 with linespoints title 'lookup time'

plot '1000000.csv' using 1:2 with linespoints title 'insertion time', '1000000.csv' using 1:3 with linespoints title 'listing time', '1000000.csv' using 1:4 with linespoints title 'lookup time'
```

## 100k lookups

![100k lookups](/ebmbtreebench/100000.png)

## 1m lookups

![1m lookups](/ebmbtreebench/1000000.png)

## 10m lookups

![10m lookups](/ebmbtreebench/10000000.png)

## References

[https://wtarreau.blogspot.com/2011/12/elastic-binary-trees-ebtree.html](https://wtarreau.blogspot.com/2011/12/elastic-binary-trees-ebtree.html)

[https://github.com/wtarreau/ebtree](https://github.com/wtarreau/ebtree)

[https://github.com/haproxy/haproxy/blob/master/src/ebmbtree.c](https://github.com/haproxy/haproxy/blob/master/src/ebmbtree.c)
