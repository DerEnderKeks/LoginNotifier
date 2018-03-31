/*
 * Copyright (C) 2018 DerEnderKeks
 *
 * This file is part of LoginNotifier.
 *
 * LoginNotifier is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * LoginNotifier is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
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
	"github.com/DerEnderKeks/LoginNotifier/util"
	"path/filepath"
)

func Init() {
	configPath := pflag.String("config", "/etc/loginnotifier/config.json", "Path to the config file")
	generateConfig := pflag.Bool("generate-config", false, "Generate the config file and exit")
	pflag.Parse()

	viper.SetConfigType("json")
	viper.SetConfigFile(*configPath)

	setDefaults()

	exists, err := util.Exists(*configPath)
	if err != nil {
		log.Warning(err)
		err = nil
	}

	if !exists {
		err = os.MkdirAll(filepath.Dir(*configPath), 0755)
		if err != nil {
			log.Warning(err)
			err = nil
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		if !(os.IsNotExist(err) && *generateConfig) {
			log.Warning(err)
		}
	}

	viper.AutomaticEnv()
	viper.WriteConfig()
	if *generateConfig {
		os.Exit(0)
	}
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
	if viper.GetBool("alerts.discord.enabled") && len(viper.GetString("alerts.discord.webhook.url")) == 0 {
		log.Fatal(errors.New("discord is enabled but the url not set"))
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
