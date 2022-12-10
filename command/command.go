package command

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	help             bool
	gpt              bool
	version, Version bool
	test, Test       bool
	quit             *bool

	s                string
	prefix           string
	configuration    string
	globalDirectives string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.BoolVar(&version, "v", false, "show version and exit")
	flag.BoolVar(&version, "V", false, "show version and configure options then exit")

	flag.BoolVar(&test, "t", false, "test configuration and exit")
	flag.BoolVar(&Test, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	quit = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&prefix, "p", "/usr/local/ctl/", "set `prefix` path")
	flag.StringVar(&configuration, "c", "conf/ctl.conf", "set configuration `file`")
	flag.StringVar(&globalDirectives, "g", "conf/ctl.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
	}
	if Version {
		flag.Usage()
	}
	if version {
		flag.Usage = versionFunc
	}
}
func versionFunc() {
	fmt.Fprintf(os.Stderr, `ctl version: ctl/1.10.0`)
}

func usage() {
	fmt.Fprintf(os.Stderr, `ctl version: ctl/1.10.0
Usage: ctl [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
