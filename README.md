# Zerops Lazy Stop Example

A demonstration of graceful shutdown handling in Go applications deployed on Zerops, showcasing the lazy stop feature with extended shutdown timeouts.

## Overview

This application demonstrates how to properly handle graceful shutdowns in Go applications with extended timeout periods. It runs a simple counter that responds to termination signals (SIGTERM, SIGINT, SIGQUIT) and then continues executing shutdown logic.

## Features

- Signal-based context cancellation
- Graceful shutdown handling with extended timeout (150 seconds)
- Simple counter demonstration during runtime and shutdown phases

## How It Works

The application:
1. Runs a counter that prints every second
2. Listens for termination signals (SIGTERM, SIGINT, SIGQUIT)
3. Stops the main loop when a signal is received
4. Continues running shutdown logic for an extended period

## Configuration

The project uses a custom `startCommandKillTimeout` of 150 seconds in `zerops.yml`, which gives the application ample time to complete shutdown procedures before being forcefully terminated.

### Systemd Tweaks

The application includes systemd configuration adjustments to support extended shutdown times:

- Creates a systemd service override directory at `/etc/systemd/system/zerops@zerops.service.d`
- Sets `TimeoutStopSec=150` to allow up to 150 seconds for service shutdown
- Reloads systemd daemon to apply the configuration changes

**Note:** Systemd has a hardcoded 90-second timeout for stop operations that cannot be extended. The configuration above attempts to set the stop timeout, but the actual behavior may still be limited by systemd's internal constraints. For graceful shutdowns exceeding 90 seconds, the application must handle the shutdown logic after the main process receives the termination signal.

## Deployment

Deploy to Zerops using the provided configuration:

