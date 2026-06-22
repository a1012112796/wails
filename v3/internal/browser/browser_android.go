//go:build android

package browser

import "errors"

// OpenURLFunc routes URL/file opens to the Android host (Intent.ACTION_VIEW).
// Android has no xdg-open, so the desktop implementation cannot be used; the
// application package registers this on Android builds to forward through the
// Java WailsBridge. It is a package-level hook to avoid an import cycle
// (internal/browser must not import pkg/application).
var OpenURLFunc func(target string) error

func open(target string) error {
	if OpenURLFunc == nil {
		return errors.New("browser: Android open handler not registered")
	}
	return OpenURLFunc(target)
}
