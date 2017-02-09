package slackHook

import "github.com/bluele/slack"

// Config contain slack client and channel information
type Config struct {
	Client      *slack.Slack
	ChannelName string
}

// Data makes a POST request with message and channel ID
type Data struct {
	Message   string
	ChannelID string
}

// New returns pointer to client
func New(token, channel string) *Config {
	client := slack.New(token)

	return &Config{
		Client:      client,
		ChannelName: channel,
	}
}

// GetChannel returns channelinfo
func (c Config) GetChannel() (*slack.Group, error) {
	groups, err := c.Client.GroupsList()
	if err != nil {
		return nil, err
	}

	for _, g := range groups {

		if c.ChannelName == g.Name {
			return g, nil
		}
	}

	return nil, nil
}

// PostMessage writes data to channel
func (c *Config) PostMessage(d *Data) error {
	err := c.Client.ChatPostMessage(d.ChannelID, d.Message, nil)
	if err != nil {
		return err
	}
	return nil
}
