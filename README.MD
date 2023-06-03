# SSH Profile Activator

[![License](https://img.shields.io/github/license/dhi13man/change_ssh_profiles)](https://github.com/Dhi13man/change_ssh_profiles/blob/main/LICENSE)
[![Contributors](https://img.shields.io/github/contributors-anon/dhi13man/change_ssh_profiles?style=flat)](https://github.com/Dhi13man/change_ssh_profiles/graphs/contributors)
[![GitHub forks](https://img.shields.io/github/forks/dhi13man/change_ssh_profiles?style=social)](https://github.com/Dhi13man/change_ssh_profiles/network/members)
[![GitHub Repo stars](https://img.shields.io/github/stars/dhi13man/change_ssh_profiles?style=social)](https://github.com/Dhi13man/change_ssh_profiles/stargazers)
[![Last Commit](https://img.shields.io/github/last-commit/dhi13man/change_ssh_profiles)](https://github.com/Dhi13man/change_ssh_profiles/commits/main)
[![Build, Format, Test | Golang](https://github.com/dhi13man/change_ssh_profiles/actions/workflows/go.yml/badge.svg)](https://github.com/Dhi13man/change_ssh_profiles/actions)
[![Build, Format, Test | Python](https://github.com/dhi13man/change_ssh_profiles/actions/workflows/python-package.yml/badge.svg)](https://github.com/Dhi13man/change_ssh_profiles/actions)

[![Language | Golang](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://golang.org/)
[![Language | Python](http://ForTheBadge.com/images/badges/made-with-python.svg)](https://www.python.org/)

[!["Buy Me A Coffee"](https://img.buymeacoffee.com/button-api/?text=Buy%20me%20an%20Ego%20boost&emoji=%F0%9F%98%B3&slug=dhi13man&button_colour=FF5F5F&font_colour=ffffff&font_family=Lato&outline_colour=000000&coffee_colour=FFDD00****)](https://www.buymeacoffee.com/dhi13man)

This repository contains two implementations of an SSH profile activator script: one in Python and another in Go. These scripts allow you to activate a specific SSH profile while disabling all other profiles, providing a convenient way to manage and switch between different SSH configurations based on your needs.

## Prerequisites

Before using either of the scripts, ensure that you have the following:

- Python 3 installed on your system (for the Python script)
- Go installed on your system (for the Go script)
- SSH profiles set up with their respective private key files

## Golang Implementation

### Golang Usage

To activate a specific SSH profile using the Go script, use the following command:

```bash
go run ssh_profile_activator.go <profile_name>
```

Replace `<profile_name>` with the name of the profile you want to activate. Valid profile names include:

- work
- personal

You can specify multiple profile names to activate multiple profiles at once:

```bash
go run ssh_profile_activator.go <profile1> <profile2> ...
```

### Golang Usage Example

To activate the "work" SSH profile and disable all other profiles using the Go script, use the following command:

```bash
go run ssh_profile_activator.go work
```

## Python Implementation

### Python Usage

To activate a specific SSH profile using the Python script, use the following command:

```bash
python ssh_profile_activator.py -p <profile_name>
```

Replace `<profile_name>` with the name of the profile you want to activate. Valid profile names include:

- work: Assumed to be a profil with the private key located in "{HOME_DIR}/.ssh/work_rsa"
- personal: Assumed to be a profile with the private key located in "{HOME_DIR}/.ssh/id_rsa"

You can add more profiles or change the paths to existing profiles by modifying the `PROFILE_MAP` dictionary in the script.

You can specify multiple profile names to activate multiple profiles at once:

```bash
python ssh_profile_activator.py -p <profile1> <profile2> ...
```

### Python Usage Example

To activate the "work" SSH profile and disable all other profiles using the Python script, use the following command:

```bash
python ssh_profile_activator.py -p work
```

## Important Notes

- Make sure you have the necessary permissions to access the SSH key files.

- The scripts assume that the SSH profile key files are located in the `.ssh` directory in the user's home directory. If your key files are located elsewhere, you will need to modify the scripts accordingly.

- For the Python script, ensure that Python 3 is installed on your system and that the necessary dependencies are met.

- For the Golang script, ensure that Go is installed on your system and that the Go environment is properly set up.

Feel free to modify and extend these scripts to suit your specific needs.
