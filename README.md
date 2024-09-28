# CoD2 URL Protocol Handler

This Go application sets up a URL protocol handler for "Call of Duty 2" (CoD2) on Windows. It retrieves the executable path from the Windows Registry and creates a custom URL protocol to enable launching CoD2 with specific commands.

## Features

- Retrieves the CoD2 executable path from the Windows Registry.
- Creates a custom URL protocol (`cod2://`) that can be used to open CoD2 with specific URLs.
- Sets the default icon for the URL protocol.
- Configures the command to launch CoD2 with the provided URL.

## Prerequisites

- Go (version 1.11 or later) installed on your machine.
- Windows operating system.
