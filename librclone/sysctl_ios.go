//go:build ios

package main

import _ "unsafe" // for go:linkname

// internal/cpu.sysctlEnabled is defined in the Go standard library only for
// `darwin && !ios` (src/internal/cpu/cpu_darwin.go) — the maintainers gate it
// out for iOS on purpose. rclone's storj backend pulls in
// storj.io/common/internal/hmacsha512, whose cpu_darwin_arm64.go reaches that
// symbol via //go:linkname to detect hardware SHA512. The _darwin_arm64
// filename tag is also active under GOOS=ios, so the reference survives into
// the iOS archive while the target symbol does not exist there, leaving
// librclone.a with an undefined `_internal/cpu.sysctlEnabled` at the C++ link
// step.
//
// Provide the symbol here (iOS only) so the archive links. Returning false
// just makes storj fall back to the generic Go SHA512 implementation, and we
// don't use the storj backend at all — so this is behaviourally inert.
// See go.dev/issue/67401 and go.dev/issue/76221.
//
//go:linkname sysctlEnabled internal/cpu.sysctlEnabled
func sysctlEnabled(name []byte) bool { return false }
