package vscode

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/loft-sh/devpod/pkg/command"
	"github.com/loft-sh/log"
	"github.com/skratchdot/open-golang/open"
)

func Open(ctx context.Context, workspace, folder string, newWindow bool, releaseChannel ReleaseChannel, log log.Logger) error {
	log.Infof("Starting VSCode...")
	err := openViaCLI(ctx, workspace, folder, newWindow, releaseChannel, log)
	if err != nil {
		log.Debugf("Error opening vscode via cli: %v", err)
	} else {
		return nil
	}

	return openViaBrowser(workspace, folder, newWindow, releaseChannel, log)
}

func openViaBrowser(workspace, folder string, newWindow bool, releaseChannel ReleaseChannel, log log.Logger) error {
	protocol := `vscode://`
	if releaseChannel == ReleaseChannelInsiders {
		protocol = `vscode-insiders://`
	}
	openURL := protocol + `vscode-remote/ssh-remote+` + workspace + `.devpod/` + folder
	if newWindow {
		openURL += "?windowId=_blank"
	}

	err := open.Run(openURL)
	if err != nil {
		log.Debugf("Starting VSCode caused error: %v", err)
		log.Errorf("Seems like you don't have Visual Studio Code installed on your computer locally. Please install VSCode via https://code.visualstudio.com/")
		return err
	}

	return nil
}

func openViaCLI(ctx context.Context, workspace, folder string, newWindow bool, releaseChannel ReleaseChannel, log log.Logger) error {
	// try to find code cli
	codePath := findCLI(releaseChannel)
	if codePath == "" {
		return fmt.Errorf("couldn't find the code binary")
	}

	// make sure ms-vscode-remote.remote-ssh is installed
	out, err := exec.Command(codePath, "--list-extensions").Output()
	if err != nil {
		return command.WrapCommandError(out, err)
	}
	splitted := strings.Split(string(out), "\n")
	found := false
	foundContainers := false
	for _, str := range splitted {
		if strings.TrimSpace(str) == "ms-vscode-remote.remote-ssh" {
			found = true
		} else if strings.TrimSpace(str) == "ms-vscode-remote.remote-containers" {
			foundContainers = true
		}
	}

	// install remote-ssh extension
	if !found {
		args := []string{"--install-extension", "ms-vscode-remote.remote-ssh"}
		log.Debugf("Run vscode command %s %s", codePath, strings.Join(args, " "))
		out, err := exec.CommandContext(ctx, codePath, args...).Output()
		if err != nil {
			return fmt.Errorf("install ssh extension: %w", command.WrapCommandError(out, err))
		}
	}

	// open vscode via cli
	args := make([]string, 0, 9)
	if foundContainers {
		args = append(args, "--disable-extension", "ms-vscode-remote.remote-containers")
	}

	// if current environment has USE_SERVE_WEV, use serve-web
	if _, ok := os.LookupEnv("USE_SERVE_WEB"); ok {
		args = append(args, "serve-web", "--host", "0.0.0.0", "--without-connection-token", "--accept-server-license-terms")
		args = append(args, "--server-base-path", folder)
		log.Debugf("Run vscode command %s %s", codePath, strings.Join(args, " "))
		out, err = exec.CommandContext(ctx, codePath, args...).CombinedOutput()
		if err != nil {
			return command.WrapCommandError(out, err)
		}
	} else {
		if newWindow {
			args = append(args, "--new-window")
		} else {
			args = append(args, "--reuse-window")
		}
		// Needs to be separated by `=` because of windows
		folderUriArg := fmt.Sprintf("--folder-uri=vscode-remote://ssh-remote+%s.devpod/%s", workspace, folder)
		args = append(args, folderUriArg)
		log.Debugf("Run vscode command %s %s", codePath, strings.Join(args, " "))
		out, err = exec.CommandContext(ctx, codePath, args...).CombinedOutput()
		if err != nil {
			return command.WrapCommandError(out, err)
		}
	}

	return nil
}

func findCLI(releaseChannel ReleaseChannel) string {
	if releaseChannel == ReleaseChannelStable {
		if command.Exists("code") {
			return "code"
		} else if runtime.GOOS == "darwin" && command.Exists("/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code") {
			return "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code"
		}

		return ""
	}

	if releaseChannel == ReleaseChannelInsiders {
		if command.Exists("code-insiders") {
			return "code-insiders"
		} else if runtime.GOOS == "darwin" && command.Exists("/Applications/Visual Studio Code - Insiders.app/Contents/Resources/app/bin/code") {
			return "/Applications/Visual Studio Code - Insiders.app/Contents/Resources/app/bin/code"
		}

		return ""
	}

	return ""
}
