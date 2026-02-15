package main

import (
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/discord/v2/discord"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createMessages(ctx *pulumi.Context, textChannels *TextChannels) error {
	_, err := discord.NewMessage(ctx, "rules-message", &discord.MessageArgs{
		ChannelId: textChannels.Rules.ChannelId,
		Pinned:    pulumi.Bool(true),
		Embed: discord.MessageEmbedArgs{
			Title:       pulumi.String("Xandaris Hardcore SMP — Rules"),
			Description: pulumi.String("Welcome to the XHC. Pure vanilla, borderline anarchy. Everything goes — except hacks."),
			Color:       pulumi.Float64(0xFF0000),
			Fields: discord.MessageEmbedFieldArray{
				discord.MessageEmbedFieldArgs{
					Name:  pulumi.String("1. No Cheating"),
					Value: pulumi.String("No hacked clients, x-ray, duping, or exploits. Play legit or don't play. This is the only real rule."),
				},
				discord.MessageEmbedFieldArgs{
					Name:  pulumi.String("2. Everything Else Is Fair Game"),
					Value: pulumi.String("Griefing, raiding, PvP, betrayal, theft — all allowed. This is hardcore. Trust no one."),
				},
				discord.MessageEmbedFieldArgs{
					Name:  pulumi.String("3. Death is Permanent"),
					Value: pulumi.String("When you die, you get the **Dead** role. You can still read and listen, but you cannot send messages or speak. No appeals, no second chances."),
				},
				discord.MessageEmbedFieldArgs{
					Name:  pulumi.String("4. Don't Be Bigoted"),
					Value: pulumi.String("Trash talk and toxicity are part of the game. Slurs and targeted harassment are not. Know the difference."),
				},
				discord.MessageEmbedFieldArgs{
					Name:  pulumi.String("5. Admin Decisions Are Final"),
					Value: pulumi.String("If an admin makes a call, respect it. Disputes go to #admin, not public channels."),
				},
			},
			Footer: discord.MessageEmbedFooterArgs{
				Text: pulumi.String("Last updated by Pulumi IaC — modify in code, not in Discord"),
			},
		},
	})
	return err
}
