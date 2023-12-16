package zymfony

import (
	"github.com/pasenidis/zimfony-agent/pkg/installer"
	"github.com/pasenidis/zimfony-agent/pkg/updater"
)

func Run() {
	installer.Install("Example")
	updater.Update("Example")
}
