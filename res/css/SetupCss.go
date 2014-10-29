package css

// #cgo pkg-config: gtk+-3.0
// #include "setup_css.h"
import "C"

func SetupCss() {
	C.setup_css()
}
