package entities

type (
	AttachmentType uint8
	MessageType    uint8
	ProtocolType   uint8
	RecipientType  uint8
)

const (
	SMTPmail MessageType = iota
	IRCmsg
)
const (
	Image AttachmentType = iota
	URL
)
const (
	From RecipientType = iota
	To
	Cc
	Bcc
)

const (
	Email ProtocolType = iota
	IRC
	XMPP
)

type IRCidentity struct {
	IRCUser     string `mapstructure:"irc_user"`
	IRCNickname string `mapstructure:"irc_nickname"`
}

type ConciergeConfig struct {
	IRCserver   string      `mapstructure:"irc_server"`
	IRCRoom     string      `mapstructure:"irc_room"`
	Concierge   IRCidentity `mapstructure:"concierge"`
	User        IRCidentity `mapstructure:"user"`
	Backend     EliasConfig `mapstructure:"elias_config"`
	FrontServer FrontConfig `mapstructure:"front_server"`
}

type EliasConfig struct {
	DataPath string `mapstructure:"data_path"`
}

type IRCconfig struct {
	IRCserver string `mapstructure:"irc_server"`
	IRCRoom   string `mapstructure:"irc_room"`
}

type FrontConfig struct {
	Listen string `mapstructure:"listen_host"`
}
