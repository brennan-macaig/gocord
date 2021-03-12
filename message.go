package gocord

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (o *Database) ReadMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore input from ourselves
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Let's see if it starts with the command we need
	if m.Content[:len(o.Conf.Command)] != o.Conf.Command {
		// This is a command, rather than a message to log. Do not log this command.
		return
	}

	// Check if it's from an ignored user
	if o.Conf.userIgnored(m.Author.ID) {
		return
	}

	// Check if it's from an ignored channel
	if o.Conf.channelIgnored(m.ChannelID) {
		return
	}

	// Log this message

}

func (o *Database) SaveMessage(m *discordgo.MessageCreate) error {
	if val, ok := o.Msgs[m.Author.ID]; ok {
		t, err := calculateTimeFromSnowflake(m.ID)
		if err != nil {
			return err
		}
		val = append(val, t.String())
		return nil
	}
	t, err := calculateTimeFromSnowflake(m.ID)
	if err != nil {
		return err
	}
	var sli []string
	sli = append(sli, t.String())
	o.Msgs[m.Author.ID] = sli
	return nil
}

func calculateTimeFromSnowflake(snowflake string) (time.Time, error) {
	i, err := strconv.ParseInt(snowflake, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not convert int - %w", err)
	}
	millisec := (i >> 22) + 1420070400000 // magic from discord docs, https://discord.com/developers/docs/reference#snowflakes
	return time.Unix(0, millisec*int64(time.Millisecond)), nil
}
