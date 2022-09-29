package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// [prefix][invoke/alias] [1st arg] [2nd arg] [3rd arg]
type CommandHandler struct {
	prefix string

	cmdInstances []Command
	cmdMap       map[string]Command

	OnError func(err error, ctx *Context)
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:       prefix,
		cmdInstances: make([]Command, 0),
		cmdMap:       make(map[string]Command),
		OnError:      func(err error, ctx *Context) {},
	}
}

func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.cmdInstances = append(c.cmdInstances, cmd)
	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}

func (c *CommandHandler) HandleMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.ID == s.State.User.ID || e.Author.Bot || !strings.HasPrefix(e.Content, c.prefix) {
		return
	}

	split := strings.Split(e.Content[len(c.prefix):], " ")
	if len(split) < 1 {
		return
	}

	invoke := strings.ToLower(split[0])
	args := split[1:]

	cmd, ok := c.cmdMap[invoke]
	if !ok || cmd == nil {
		return
	}

	ctx := &Context{
		Session: s,
		Message: e.Message,
		Args:    args,
		Handler: c,
	}

	if err := cmd.Execute(ctx); err != nil {
		c.OnError(err, ctx)
	}
}
