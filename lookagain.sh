#!/bin/bash

find . -name "*.sh" -exec sh -c 'printf "%s\n" "${0%.*}"' {} ';' | cut -b 3-
