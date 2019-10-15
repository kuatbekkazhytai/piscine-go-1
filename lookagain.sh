#!/bin/bash

#find -d -name "*.sh" | cut -b 2- | cut -d '.' -f1
find -type f -name "*.sh" -exec basename {} .sh \;
