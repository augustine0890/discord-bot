package commands

type CmdPing struct{}

func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p", "pong"}
}

func (c *CmdPing) Description() string {
	return "Pong!"
}

func (c *CmdPing) AdminRequired() bool {
	return false
}

func (c *CmdPing) Execute(ctx *Context) (err error) {
	// Ignore all message created by the bot itself
	if ctx.Message.Author.ID == ctx.Session.State.User.ID {
		return
	}

	switch ctx.Message.Content {
	case "!ping":
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!")
	case "!pong":
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Ping!")
	}

	return
}
