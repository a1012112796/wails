//go:build android

package application

import "github.com/wailsapp/wails/v3/internal/browser"

// Route browser.OpenURL / browser.OpenFile through the Android host, which opens
// the URL with Intent.ACTION_VIEW. The desktop path (xdg-open) does not exist on
// Android. Registered via the package-level hook to keep internal/browser free
// of any dependency on this package.
func init() {
	browser.OpenURLFunc = func(target string) error {
		Android.OpenURL(target)
		return nil
	}
}
