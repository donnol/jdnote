#!/bin/bash

# 压力测试，看情况决定是否加'-k'标志
ab -n 1000 -c 100 http://localhost:8810$@