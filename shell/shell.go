package shell

import (
	"os"
)

// Shell configurations
type Shell struct {
	name        string
	rcPath      string
	historyPath string
}

var (
	// Sh defaul configs
	Sh = Shell{"/bin/sh", "/etc/profile", "~/.sh_history"}

	// Bash default configs
	Bash = Shell{"/bin/bash", "~/.bashrc", "~/.bash_history"}

	//Zsh default configs
	Zsh = Shell{"/bin/zsh", "~/.zshrc", "~/.zsh_history"}
)

// GetShell returns a struct representing the shells configuration
func GetShell() Shell {
	shell := os.Getenv("SHELL")

	switch shell {
	case Bash.name:
		return Bash
	case Zsh.name:
		return Zsh
	default:
		return Sh
	}

}
