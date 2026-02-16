# Xandaris Hardcore SMP Discord Infrastructure

This repository contains the Infrastructure as Code (IaC) for the Xandaris Hardcore SMP Discord server. It uses **Pulumi** with the **Go** runtime to manage channels, roles, and permissions.

## Features

- **Automated Channel Structure**: Organizes the server into Info, General, Game, Voice, and Admin categories.
- **Hardcore Mechanics**: Implements custom permission logic for the **Dead** role, automatically silencing players across the server.
- **Role Management**: Manages staff (Admin, Moderator) and participant (Player, Dead) roles.
- **Automated Content**: Deploys and pins rules and server information embeds in `#rules` and `#server-status`.
- **CI/CD Integration**: Uses GitHub Actions to automatically deploy changes pushed to the `main` branch.

## Prerequisites

- [Pulumi CLI](https://www.pulumi.com/docs/get-started/install/)
- [Go](https://golang.org/doc/install) (v1.24.0+)
- [Discord Token](https://discord.com/developers/docs/intro) (set via `DISCORD_TOKEN` environment variable)

## Setup & Deployment

1. Install the Discord provider:
   ```bash
   pulumi package add terraform-provider lucky3028/discord
   ```

2. Download Go dependencies:
   ```bash
   go mod download
   ```

3. Deploy changes:
   ```bash
   pulumi up --stack prod
   ```

## Repository Structure

- `main.go`: Entry point for the Pulumi application.
- `roles.go`: Definition of server roles.
- `channels.go` & `categories.go`: Server structure definitions.
- `permissions.go`: Custom permission logic and overrides.
- `messages.go`: Automated message embeds.
- `state/`: Local Pulumi state storage.
