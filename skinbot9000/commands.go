package main

import (
	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name: "basic-command",
		// All commands and options must have a description
		// Commands/options without description will fail the registration
		// of the command.
		Description: "Basic command",
	},
	{
		Name:        "options",
		Description: "Command for demonstrating options",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "String option",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "integer-option",
				Description: "Integer option",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "bool-option",
				Description: "Boolean option",
				Required:    true,
			},

			// Required options must be listed first since optional parameters
			// always come after when they're used.
			// The same concept applies to Discord's Slash-commands API

			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel-option",
				Description: "Channel option",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user-option",
				Description: "User option",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionRole,
				Name:        "role-option",
				Description: "Role option",
				Required:    false,
			},
		},
	},
	{
		Name:        "subcommands",
		Description: "Subcommands and command groups example",
		Options: []*discordgo.ApplicationCommandOption{
			// When a command has subcommands/subcommand groups
			// It must not have top-level options, they aren't accesible in the UI
			// in this case (at least not yet), so if a command has
			// subcommands/subcommand any groups registering top-level options
			// will cause the registration of the command to fail

			{
				Name:        "scmd-grp",
				Description: "Subcommands group",
				Options: []*discordgo.ApplicationCommandOption{
					// Also, subcommand groups aren't capable of
					// containing options, by the name of them, you can see
					// they can only contain subcommands
					{
						Name:        "nst-subcmd",
						Description: "Nested subcommand",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
				},
				Type: discordgo.ApplicationCommandOptionSubCommandGroup,
			},
			// Also, you can create both subcommand groups and subcommands
			// in the command at the same time. But, there's some limits to
			// nesting, count of subcommands (top level and nested) and options.
			// Read the intro of slash-commands docs on Discord dev portal
			// to get more information
			{
				Name:        "subcmd",
				Description: "Top-level subcommand",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	},
	{
		Name:        "responses",
		Description: "Interaction responses testing initiative",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "resp-type",
				Description: "Response type",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Channel message with source",
						Value: 4,
					},
					{
						Name:  "Deferred response With Source",
						Value: 5,
					},
				},
				Required: true,
			},
		},
	},
	{
		Name:        "followups",
		Description: "Followup messages",
	},
}
