#!/bin/bash
gocr -i "${1}" -C 0-9 -a 60 -s 100  2> /dev/null | build/run
