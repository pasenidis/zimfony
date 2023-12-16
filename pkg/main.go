package zymfony

import (
	"zimfony-agent/pkg/installer"
	"zimfony-agent/pkg/updater"
)

func Run() {
	installer.Install("Example")
	updater.Update("Example")
}
