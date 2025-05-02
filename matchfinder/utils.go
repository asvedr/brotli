package matchfinder

import "runtime"

func CheckArch() (is32 bool, is64 bool) {
	switch runtime.GOARCH {
	case "386", "amd64p32", "arm", "armbe", "mips",
		"mips64p32", "mips64p32le", "mipsle", "ppc",
		"riscv", "s390", "sparc":
		return true, false
	case "amd64", "arm64", "arm64be", "loong64", "mips64",
		"mips64le", "ppc64", "ppc64le", "riscv64", "s390x",
		"sparc64", "wasm":
		return false, true
	default:
		return false, false
	}
}
