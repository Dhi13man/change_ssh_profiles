#!/usr/bin/env python3

"""
Python script to activate particular SSH profile, disabling all other profiles
"""

from argparse import ArgumentParser
from os import environ
from subprocess import call, DEVNULL
import sys

HOME_DIR: str = environ["HOME"]
PROFILE_MAP: dict = {
    "work": f"{HOME_DIR}/.ssh/work_rsa",
    "personal": f"{HOME_DIR}/.ssh/id_rsa",
}


class SshConfigureArgParseWrapper:
    """
    Wrapper class for the argparse module.
    """

    def __init__(self) -> None:
        self.parser = ArgumentParser()
        self.parser.add_argument(
            "-p",
            "--profiles",
            help="Names of the profiles to activate. Valid values are: " +
            ", ".join(PROFILE_MAP.keys()) + ".",
            default=["None"],
            nargs="+"
        )

    def get_arg_profile(self) -> list:
        """
        Get the profile name from the command line arguments.
        """
        return self.parser.parse_args().profiles

    def print_help(self) -> None:
        """
        Print the help message.
        """
        self.parser.print_help()


if __name__ == "__main__":
    # Get the profile name from the command line arguments.
    argparse_wrapper: SshConfigureArgParseWrapper = SshConfigureArgParseWrapper()
    profiles: list = argparse_wrapper.get_arg_profile()
    for profile in profiles:
        if profile not in PROFILE_MAP:
            print("Invalid profile name: " + profile)
            argparse_wrapper.print_help()
            sys.exit(1)

    # Disable all other SSH profiles.
    for path in PROFILE_MAP.values():
        call(
            ["ssh-add", "-d", path],
            stderr=DEVNULL
        )

    # Enable the specified SSH profile.
    for profile in profiles:
        # Get the path to the SSH profile.
        ssh_profile_path: str = PROFILE_MAP[profile]
        call(["ssh-add", "--apple-use-keychain", ssh_profile_path])

    call(["ssh-agent", "-s"])
