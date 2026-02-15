package main

import (
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/discord/v2/discord"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createMembers(ctx *pulumi.Context, serverId pulumi.StringInput, roles *Roles) error {
	// Hunter — server owner, Admin role
	_, err := discord.NewMemberRoles(ctx, "hunter", &discord.MemberRolesArgs{
		ServerId: serverId,
		UserId:   pulumi.String("371034483836846090"),
		Roles: discord.MemberRolesRoleArray{
			discord.MemberRolesRoleArgs{
				RoleId:  roles.Admin.ID().ToStringOutput(),
				HasRole: pulumi.Bool(true),
			},
		},
	})
	if err != nil {
		return err
	}

	// User 1291416111916515378 — Admin role
	_, err = discord.NewMemberRoles(ctx, "admin-1291416111916515378", &discord.MemberRolesArgs{
		ServerId: serverId,
		UserId:   pulumi.String("1291416111916515378"),
		Roles: discord.MemberRolesRoleArray{
			discord.MemberRolesRoleArgs{
				RoleId:  roles.Admin.ID().ToStringOutput(),
				HasRole: pulumi.Bool(true),
			},
		},
	})
	return err
}
