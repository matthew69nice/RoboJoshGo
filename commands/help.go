package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/matthewlowe/Robojosh/framework"
)

func HelpCommand(context *framework.Context) {
	commandMap := *context.CmdRegistry.GetCommandMap()

	names := make([]string, len(commandMap))
	helps := make([]string, len(commandMap))
	c := 0

	for name, commandStruct := range commandMap {
		names[c] = name
		helps[c] = commandStruct.Help
		c++
	}

	fields := make([]*discordgo.MessageEmbedField, c)

	for i := 0; i < c; i++ {
		fields[i] = &discordgo.MessageEmbedField{
			Name:   names[i],
			Value:  helps[i],
			Inline: false,
		}
	}

	footer := discordgo.MessageEmbedFooter{
		Text:    "brought to you by your local electronic nugget lover",
		IconURL: context.Discord.State.User.AvatarURL("256x256"),
	}

	embed := discordgo.MessageEmbed{
		Type:        discordgo.EmbedType("rich"),
		Title:       "List of commands:",
		Description: "Current prefix: " + context.Prefix,
		//Timestamp: time.Now().Format("15:04"),
		Color:  10038562, // DARKER_RED
		Footer: &footer,
		Fields: fields,
	}

	context.Discord.ChannelMessageSendEmbed(context.TextChannel.ID, &embed)
}