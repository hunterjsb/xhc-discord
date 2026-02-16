package main

import (
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/discord/v2/discord"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createPermissions(
	ctx *pulumi.Context,
	serverId pulumi.StringInput,
	roles *Roles,
	cats *Categories,
	textChannels *TextChannels,
	voiceChannels *VoiceChannels,
) error {
	// --- Admin category: deny @everyone, allow Admin + Moderator ---

	// @everyone role ID == server ID in Discord
	_, err := discord.NewChannelPermission(ctx, "admin-cat-deny-everyone", &discord.ChannelPermissionArgs{
		ChannelId:   cats.Admin.ChannelId,
		Type:        pulumi.String("role"),
		OverwriteId: serverId,
		Allow:       pulumi.Float64(0),
		Deny:        pulumi.Float64(PermTextAll),
	})
	if err != nil {
		return err
	}

	_, err = discord.NewChannelPermission(ctx, "admin-cat-allow-admin", &discord.ChannelPermissionArgs{
		ChannelId:   cats.Admin.ChannelId,
		Type:        pulumi.String("role"),
		OverwriteId: roles.Admin.ID().ToStringOutput(),
		Allow:       pulumi.Float64(PermTextAll),
		Deny:        pulumi.Float64(0),
	})
	if err != nil {
		return err
	}

	_, err = discord.NewChannelPermission(ctx, "admin-cat-allow-mod", &discord.ChannelPermissionArgs{
		ChannelId:   cats.Admin.ChannelId,
		Type:        pulumi.String("role"),
		OverwriteId: roles.Moderator.ID().ToStringOutput(),
		Allow:       pulumi.Float64(PermTextAll),
		Deny:        pulumi.Float64(0),
	})
	if err != nil {
		return err
	}

	// --- Dev Server Console: Deny Moderator (Ensure Admin-only) ---

	_, err = discord.NewChannelPermission(ctx, "dev-server-console-deny-mod", &discord.ChannelPermissionArgs{
		ChannelId:   textChannels.DevServerConsole.ChannelId,
		Type:        pulumi.String("role"),
		OverwriteId: roles.Moderator.ID().ToStringOutput(),
		Allow:       pulumi.Float64(0),
		Deny:        pulumi.Float64(PermTextAll),
	})
	if err != nil {
		return err
	}

	// --- Dead role: deny SEND_MESSAGES and SPEAK on all public text channels ---

	deadDenyBits := float64(PermSendMessages | PermSpeak)

	textChannelIDs := map[string]pulumi.StringOutput{
		"rules":         textChannels.Rules.ChannelId,
		"announcements": textChannels.Announcements.ChannelId,
		"server-status": textChannels.ServerStatus.ChannelId,
		"general":       textChannels.General.ChannelId,
		"media":         textChannels.Media.ChannelId,
		"bot-commands":  textChannels.BotCommands.ChannelId,
		"coordinates":   textChannels.Coordinates.ChannelId,
		"builds":        textChannels.Builds.ChannelId,
		"deaths":        textChannels.Deaths.ChannelId,
		"trading":       textChannels.Trading.ChannelId,
	}

	for chName, chID := range textChannelIDs {
		_, err = discord.NewChannelPermission(ctx, "dead-deny-"+chName, &discord.ChannelPermissionArgs{
			ChannelId:   chID,
			Type:        pulumi.String("role"),
			OverwriteId: roles.Dead.ID().ToStringOutput(),
			Allow:       pulumi.Float64(0),
			Deny:        pulumi.Float64(deadDenyBits),
		})
		if err != nil {
			return err
		}
	}

	// Dead role: deny SPEAK on voice channels
	voiceChannelIDs := map[string]pulumi.StringOutput{
		"general-vc": voiceChannels.General.ChannelId,
		"gaming-vc":  voiceChannels.Gaming.ChannelId,
	}

	for vcName, vcID := range voiceChannelIDs {
		_, err = discord.NewChannelPermission(ctx, "dead-deny-"+vcName, &discord.ChannelPermissionArgs{
			ChannelId:   vcID,
			Type:        pulumi.String("role"),
			OverwriteId: roles.Dead.ID().ToStringOutput(),
			Allow:       pulumi.Float64(0),
			Deny:        pulumi.Float64(PermSpeak),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
