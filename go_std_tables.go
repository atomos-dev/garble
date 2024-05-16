// Code generated by scripts/gen-go-std-tables.sh; DO NOT EDIT.

// Generated from Go version go1.22.0.

package main

var runtimeAndDeps = map[string]bool{
	"internal/goarch":          true,
	"unsafe":                   true,
	"internal/abi":             true,
	"internal/cpu":             true,
	"internal/bytealg":         true,
	"internal/chacha8rand":     true,
	"internal/coverage/rtcov":  true,
	"internal/godebugs":        true,
	"internal/goexperiment":    true,
	"internal/goos":            true,
	"runtime/internal/atomic":  true,
	"runtime/internal/math":    true,
	"runtime/internal/sys":     true,
	"runtime/internal/syscall": true,
	"runtime":                  true,
}

var runtimeLinknamed = []string{
	"arena",
	"crypto/internal/boring",
	"crypto/internal/boring/bcache",
	"crypto/internal/boring/fipstls",
	"crypto/x509/internal/macos",
	"internal/godebug",
	"internal/poll",
	"internal/reflectlite",
	"internal/syscall/unix",
	"internal/syscall/windows",
	"maps",
	"os",
	"os/signal",
	"plugin",
	"reflect",
	"runtime/coverage",
	"runtime/debug",
	"runtime/metrics",
	"runtime/pprof",
	"runtime/trace",
	"sync",
	"sync/atomic",
	"syscall",
	"syscall/js",
	"time",
	// The net package linknames to the runtime, not the other way around.
	// TODO: support this automatically via our script.
	"net",
}

var compilerIntrinsicsPkgs = map[string]bool{
	"math":                    true,
	"math/big":                true,
	"math/bits":               true,
	"runtime":                 true,
	"runtime/internal/atomic": true,
	"runtime/internal/math":   true,
	"runtime/internal/sys":    true,
	"sync":                    true,
	"sync/atomic":             true,
}

var compilerIntrinsicsFuncs = map[string]bool{
	"math.Abs":                                true,
	"math/big.mulWW":                          true,
	"math/bits.Add":                           true,
	"math/bits.Add64":                         true,
	"math/bits.Div":                           true,
	"math/bits.Div64":                         true,
	"math/bits.Len":                           true,
	"math/bits.Len16":                         true,
	"math/bits.Len32":                         true,
	"math/bits.Len64":                         true,
	"math/bits.Len8":                          true,
	"math/bits.Mul":                           true,
	"math/bits.Mul64":                         true,
	"math/bits.OnesCount":                     true,
	"math/bits.OnesCount16":                   true,
	"math/bits.OnesCount32":                   true,
	"math/bits.OnesCount64":                   true,
	"math/bits.OnesCount8":                    true,
	"math/bits.Reverse":                       true,
	"math/bits.Reverse16":                     true,
	"math/bits.Reverse32":                     true,
	"math/bits.Reverse64":                     true,
	"math/bits.Reverse8":                      true,
	"math/bits.ReverseBytes16":                true,
	"math/bits.ReverseBytes32":                true,
	"math/bits.ReverseBytes64":                true,
	"math/bits.RotateLeft":                    true,
	"math/bits.RotateLeft16":                  true,
	"math/bits.RotateLeft32":                  true,
	"math/bits.RotateLeft64":                  true,
	"math/bits.RotateLeft8":                   true,
	"math/bits.Sub":                           true,
	"math/bits.Sub64":                         true,
	"math/bits.TrailingZeros16":               true,
	"math/bits.TrailingZeros32":               true,
	"math/bits.TrailingZeros64":               true,
	"math/bits.TrailingZeros8":                true,
	"math.Ceil":                               true,
	"math.Copysign":                           true,
	"math.Floor":                              true,
	"math.FMA":                                true,
	"math.Round":                              true,
	"math.RoundToEven":                        true,
	"math.sqrt":                               true,
	"math.Trunc":                              true,
	"runtime/internal/atomic.And":             true,
	"runtime/internal/atomic.And8":            true,
	"runtime/internal/atomic.Cas":             true,
	"runtime/internal/atomic.Cas64":           true,
	"runtime/internal/atomic.Casint32":        true,
	"runtime/internal/atomic.Casint64":        true,
	"runtime/internal/atomic.Casp1":           true,
	"runtime/internal/atomic.CasRel":          true,
	"runtime/internal/atomic.Casuintptr":      true,
	"runtime/internal/atomic.Load":            true,
	"runtime/internal/atomic.Load64":          true,
	"runtime/internal/atomic.Load8":           true,
	"runtime/internal/atomic.LoadAcq":         true,
	"runtime/internal/atomic.LoadAcq64":       true,
	"runtime/internal/atomic.LoadAcquintptr":  true,
	"runtime/internal/atomic.Loadint32":       true,
	"runtime/internal/atomic.Loadint64":       true,
	"runtime/internal/atomic.Loadp":           true,
	"runtime/internal/atomic.Loaduint":        true,
	"runtime/internal/atomic.Loaduintptr":     true,
	"runtime/internal/atomic.Or":              true,
	"runtime/internal/atomic.Or8":             true,
	"runtime/internal/atomic.Store":           true,
	"runtime/internal/atomic.Store64":         true,
	"runtime/internal/atomic.Store8":          true,
	"runtime/internal/atomic.Storeint32":      true,
	"runtime/internal/atomic.Storeint64":      true,
	"runtime/internal/atomic.StorepNoWB":      true,
	"runtime/internal/atomic.StoreRel":        true,
	"runtime/internal/atomic.StoreRel64":      true,
	"runtime/internal/atomic.StoreReluintptr": true,
	"runtime/internal/atomic.Storeuintptr":    true,
	"runtime/internal/atomic.Xadd":            true,
	"runtime/internal/atomic.Xadd64":          true,
	"runtime/internal/atomic.Xaddint32":       true,
	"runtime/internal/atomic.Xaddint64":       true,
	"runtime/internal/atomic.Xadduintptr":     true,
	"runtime/internal/atomic.Xchg":            true,
	"runtime/internal/atomic.Xchg64":          true,
	"runtime/internal/atomic.Xchgint32":       true,
	"runtime/internal/atomic.Xchgint64":       true,
	"runtime/internal/atomic.Xchguintptr":     true,
	"runtime/internal/math.Add64":             true,
	"runtime/internal/math.Mul64":             true,
	"runtime/internal/math.MulUintptr":        true,
	"runtime/internal/sys.Bswap32":            true,
	"runtime/internal/sys.Bswap64":            true,
	"runtime/internal/sys.Len64":              true,
	"runtime/internal/sys.Len8":               true,
	"runtime/internal/sys.OnesCount64":        true,
	"runtime/internal/sys.Prefetch":           true,
	"runtime/internal/sys.PrefetchStreamed":   true,
	"runtime/internal/sys.TrailingZeros32":    true,
	"runtime/internal/sys.TrailingZeros64":    true,
	"runtime/internal/sys.TrailingZeros8":     true,
	"runtime.publicationBarrier":              true,
	"sync/atomic.AddInt32":                    true,
	"sync/atomic.AddInt64":                    true,
	"sync/atomic.AddUint32":                   true,
	"sync/atomic.AddUint64":                   true,
	"sync/atomic.AddUintptr":                  true,
	"sync/atomic.CompareAndSwapInt32":         true,
	"sync/atomic.CompareAndSwapInt64":         true,
	"sync/atomic.CompareAndSwapUint32":        true,
	"sync/atomic.CompareAndSwapUint64":        true,
	"sync/atomic.CompareAndSwapUintptr":       true,
	"sync/atomic.LoadInt32":                   true,
	"sync/atomic.LoadInt64":                   true,
	"sync/atomic.LoadPointer":                 true,
	"sync/atomic.LoadUint32":                  true,
	"sync/atomic.LoadUint64":                  true,
	"sync/atomic.LoadUintptr":                 true,
	"sync/atomic.StoreInt32":                  true,
	"sync/atomic.StoreInt64":                  true,
	"sync/atomic.StoreUint32":                 true,
	"sync/atomic.StoreUint64":                 true,
	"sync/atomic.StoreUintptr":                true,
	"sync/atomic.SwapInt32":                   true,
	"sync/atomic.SwapInt64":                   true,
	"sync/atomic.SwapUint32":                  true,
	"sync/atomic.SwapUint64":                  true,
	"sync/atomic.SwapUintptr":                 true,
	"sync.runtime_LoadAcquintptr":             true,
	"sync.runtime_StoreReluintptr":            true,
}

var reflectSkipPkg = map[string]bool{
	"fmt": true,
}
