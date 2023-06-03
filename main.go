package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	// The name of the environment variable that stores the path to the user's home directory.
	homeDirEnvVariable string = os.Getenv("HOME")

	// A map of SSH profile names to their absolute paths.
	profileMap map[string]string = map[string]string{
		"work":     homeDirEnvVariable + "/.ssh/work_rsa",
		"personal": homeDirEnvVariable + "/.ssh/id_rsa",
	}

	// The command to add an SSH profile to the ssh-agent.
	sshAddCmd string = "ssh-add"

	// The command to restart the ssh-agent.
	sshAgentCmd string = "ssh-agent"

	// The argument to use the macOS keychain with ssh-add.
	sshAddUseKeychainArg string = "--apple-use-keychain"
)

func main() {
	// Get the profile names from the command line arguments.
	profiles := os.Args[1:]
	for _, profile := range profiles {
		if _, ok := profileMap[profile]; !ok {
			fmt.Println("Invalid profile name: " + profile)
			fmt.Println("Valid profiles are: " + fmt.Sprintf("%v", profileMap))
			os.Exit(1)
		}
	}

	// Disable all other SSH profiles. Ignore exit code 1, which means the
	// profile is not currently enabled.
	for _, path := range profileMap {
		sshAddCmdOutput := exec.Command(sshAddCmd, "-d", path)
		sshAddCmdOutput.Run()
	}

	// Enable the specified SSH profile.
	for _, profile := range profiles {
		// Get the path to the SSH profile.
		sshProfilePath := profileMap[profile]
		sshAddArgs := []string{sshAddUseKeychainArg, sshProfilePath}
		sshAddCmdOutput := exec.Command(sshAddCmd, sshAddArgs...)
		if err := sshAddCmdOutput.Run(); err != nil {
			fmt.Println("Error while enabling profile: " + err.Error())
			os.Exit(1)
		}
	}

	// Restart the ssh-agent.
	sshAgentCmd := exec.Command(sshAgentCmd, "-s")
	if err := sshAgentCmd.Run(); err != nil {
		fmt.Println("Error while restarting ssh-agent: " + err.Error())
		os.Exit(1)
	}
	fmt.Printf("SSH profile %s activated successfully.", profiles)
}
