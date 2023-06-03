# SSH Profile Activator

This repository contains two implementations of an SSH profile activator script: one in Python and another in Go. These scripts allow you to activate a specific SSH profile while disabling all other profiles, providing a convenient way to manage and switch between different SSH configurations based on your needs.

## Prerequisites

Before using either of the scripts, ensure that you have the following:

- Python 3 installed on your system (for the Python script)
- Go installed on your system (for the Go script)
- SSH profiles set up with their respective private key files

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

## Important Notes

- Make sure you have the necessary permissions to access the SSH key files.

- The scripts assume that the SSH profile key files are located in the `.ssh` directory in the user's home directory. If your key files are located elsewhere, you will need to modify the scripts accordingly.

- For the Python script, ensure that Python 3 is installed on your system and that the necessary dependencies are met.

- For the Golang script, ensure that Go is installed on your system and that the Go environment is properly set up.

Feel free to modify and extend these scripts to suit your specific needs.
