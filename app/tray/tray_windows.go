package tray

import (
	"github.com/richardjennings/ollama/app/tray/commontray"
	"github.com/richardjennings/ollama/app/tray/wintray"
)

func InitPlatformTray(icon, updateIcon []byte) (commontray.OllamaTray, error) {
	return wintray.InitTray(icon, updateIcon)
}
