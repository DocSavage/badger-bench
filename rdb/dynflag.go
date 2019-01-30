// +build !embed

package rdb

// #cgo CXXFLAGS: -std=c++11 -O2
// #cgo LDFLAGS: -L/usr/local/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy
// #include <stdint.h>
// #include <stdlib.h>
// #include "rdbc.h"
import "C"
