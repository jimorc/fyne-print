//go:build !windows

package print

//#cgo LDFLAGS: -lcups
import "C"
