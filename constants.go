package main

// Discord permission bit flags.
// See https://discord.com/developers/docs/topics/permissions
const (
	PermViewChannel         = 0x00000400
	PermSendMessages        = 0x00000800
	PermManageMessages      = 0x00002000
	PermEmbedLinks          = 0x00004000
	PermAttachFiles         = 0x00008000
	PermReadMessageHistory  = 0x00010000
	PermConnect             = 0x00100000
	PermSpeak               = 0x00200000
	PermAdministrator       = 0x00000008
)

// Composite permission sets.
const (
	PermPlayer = PermViewChannel | PermSendMessages | PermEmbedLinks |
		PermAttachFiles | PermReadMessageHistory | PermConnect | PermSpeak

	PermDead = PermViewChannel | PermReadMessageHistory | PermConnect

	PermModerator = PermManageMessages | PermViewChannel | PermSendMessages

	PermTextAll = PermViewChannel | PermSendMessages | PermManageMessages |
		PermEmbedLinks | PermAttachFiles | PermReadMessageHistory
)
