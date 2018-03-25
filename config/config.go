/*
 * Copyright (C) 2018 DerEnderKeks
 *
 * This file is part of LoginNotifier.
 *
 * LoginNotifier is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * LoginNotifier is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with LoginNotifier. If not, see <http://www.gnu.org/licenses/>.
 */

package config

import (
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"github.com/DerEnderKeks/LoginNotifier/log"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	pflag.String("config", "", "Path to the config file")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/loginnotifier")

	if viper.GetString("config") != "" {
		viper.SetConfigFile(viper.GetString("config"))
	}

	setDefaults()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warning(err)
	}

	viper.AutomaticEnv()
	viper.WriteConfig()
	checkConfig()
	registerSIGHUP()
}

func setDefaults() {
	viper.SetDefault("loglevel", log.LevelInfo.ID)
	viper.SetDefault("source_log", "/var/log/auth.log")
	// Slack
	viper.SetDefault("alerts.slack.enabled", "false")
	viper.SetDefault("alerts.slack.webhook.url", "")
	viper.SetDefault("alerts.slack.webhook.channel", "#general")
	viper.SetDefault("alerts.slack.webhook.username", "LoginNotifier")
	viper.SetDefault("alerts.slack.webhook.icon_emoji", ":key:")
	viper.SetDefault("alerts.slack.webhook.color", "#6100ff")
	// Discord
	viper.SetDefault("alerts.discord.enabled", "false")
	viper.SetDefault("alerts.discord.webhook.url", "")
	viper.SetDefault("alerts.discord.webhook.username", "LoginNotifier")
	viper.SetDefault("alerts.discord.webhook.avatar_url", "https://cdnjs.cloudflare.com/ajax/libs/emojione/2.2.7/assets/png/1f511.png")
	viper.SetDefault("alerts.discord.webhook.tts", false)
	viper.SetDefault("alerts.discord.webhook.color", "#6100ff")
}

func checkConfig() {
	if viper.GetBool("alerts.slack.enabled") && len(viper.GetString("alerts.slack.webhook.url")) == 0 {
		log.Fatal(errors.New("slack is enabled but the url not set"))
	}
}

func registerSIGHUP() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for range c {
			log.Info("Received SIGHUP. Reloading config file...")
			viper.ReadInConfig()
			checkConfig()
		}
	}()
}
