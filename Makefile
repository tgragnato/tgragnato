OBJS = ebtree.o eb32tree.o eb64tree.o ebmbtree.o ebsttree.o ebimtree.o ebistree.o
CFLAGS = -O3 -W -Wall -Wextra -Wundef -Wdeclaration-after-statement -Wno-address-of-packed-member
EXAMPLES = $(basename $(wildcard examples/*.c))
VALUES = 1 10 100 1000 10000 100000 1000000 10000000

all: libebtree.a

examples: ${EXAMPLES}

libebtree.a: $(OBJS)
	$(AR) rv $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -o $@ -c $^

examples/%: examples/%.c libebtree.a
	$(CC) $(CFLAGS) -I. -o $@ $< -L. -lebtree

test: test32 test64 testst

test%: test%.c libebtree.a
	$(CC) $(CFLAGS) -o $@ $< -L. -lebtree

ebmbtreebench: ebmbtreebench/ebmbtreebench.c libebtree.a
	$(CC) $(CFLAGS) -I. -o ebmbtreebench/$@ $< -L. -lebtree

100000: ebmbtreebench
	$(foreach var,$(VALUES),./ebmbtreebench/ebmbtreebench $(var) $@ >> ebmbtreebench/$@.csv;)

1000000: ebmbtreebench
	$(foreach var,$(VALUES),./ebmbtreebench/ebmbtreebench $(var) $@ >> ebmbtreebench/$@.csv;)

10000000: ebmbtreebench
	$(foreach var,$(VALUES),./ebmbtreebench/ebmbtreebench $(var) $@ >> ebmbtreebench/$@.csv;)

clean:
	-rm -fv libebtree.a $(OBJS) *~ *.o *.rej core test32 test64 testst ebmbtreebench/*.csv ebmbtreebench/ebmbtreebench ${EXAMPLES}

ifeq ($(wildcard .git),.git)
VERSION := $(shell [ -d .git/. ] && ref=`(git describe --tags --match 'v*') 2>/dev/null` && ref=$${ref%-g*} && echo "$${ref\#v}")
SUBVERS := $(shell comms=`git log --no-merges v$(VERSION).. 2>/dev/null |grep -c ^commit `; [ $$comms -gt 0 ] && echo "-$$comms" )
endif

git-tar: .git
	git archive --format=tar --prefix="ebtree-$(VERSION)/" HEAD | gzip -9 > ebtree-$(VERSION)$(SUBVERS).tar.gz

.PHONY: examples tests
