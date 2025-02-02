package core

import (
	"log"

	"github.com/hemanth2004/doomsday-protocol/dday/debug"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Tui struct {
		Main struct {
			Theme              string `koanf:"tui.main.theme"`
			UseAlternateBuffer bool   `koanf:"tui.main.use_alternate_buffer"`
			TickInterval       int    `koanf:"tui.main.tick_interval"`
		}
		Colors struct {
			PrimaryColor   string `koanf:"tui.colors.primary"`
			SecondaryColor string `koanf:"tui.colors.secondary"`
			TertiaryColor  string `koanf:"tui.colors.tertiary"`
			Accent1Color   string `koanf:"tui.colors.accent1"`
			Accent2Color   string `koanf:"tui.colors.accent2"`
			Accent3Color   string `koanf:"tui.colors.accent3"`
			Accent4Color   string `koanf:"tui.colors.accent4"`
		}
	}
	Logs struct {
		SaveLogs      bool   `koanf:"logs.save_logs"`
		LogFilePath   string `koanf:"logs.log_file_path"`
		ShowTimestamp bool   `koanf:"logs.show_timestamp"`
		MaxLogLines   int    `koanf:"logs.max_log_lines"`
	}
}

func LoadConfig(workingDirectory string) *Config {
	var k = koanf.New(workingDirectory)

	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	debug.Log("use_alternate_buffer = " + k.String("tui.main.use_alternate_buffer"))

	return &Config{}
}
