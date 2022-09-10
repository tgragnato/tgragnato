#!/usr/bin/env python3

import os
import sys

input = 'pixel_madness.txt'
output = 'pixel_madness.bin'

if not os.path.isfile(input):
    print('{} non esiste'.format(input))
    sys.exit()

if os.path.isfile(output):
    os.remove(output)
fo = open(output, 'w')

with open(input) as fi:
    for linea in fi:
        dupla = linea.split('+')
        for item in dupla:
            x = item.split('x')
            for _ in range(int(x[1])):
                if int(x[0]) == 0:
                    fo.write('X')
                else:
                    fo.write(' ')
        fo.write('\n')

fo.close()