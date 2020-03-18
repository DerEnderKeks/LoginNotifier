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

package alert

import (
	"github.com/DerEnderKeks/LoginNotifier/parser"
	"github.com/DerEnderKeks/LoginNotifier/util"
	"github.com/noonien/discohook"
	"github.com/spf13/viper"
	"time"
)

func discordAlert(session parser.Session) {
	message := discohook.Message{}
	message.Username = viper.GetString("alerts.discord.webhook.username")
	message.AvatarURL = viper.GetString("alerts.discord.webhook.avatar_url")
	message.TTS = viper.GetBool("alerts.discord.webhook.tts")
	userField := discohook.EmbedField{Name: "User", Value: session.User(), Inline: true}
	ipField := discohook.EmbedField{Name: "IP", Value: "[" + session.IP() + "](https://whois.domaintools.com/" + session.IP() + ") (" + session.ReverseHost() + ")", Inline: true}
	hostField := discohook.EmbedField{Name: "Host", Value: session.Host(), Inline: true}
	timeField := discohook.EmbedField{Name: "Time", Value: session.Time().Format(time.Stamp), Inline: true}
	color := util.HexToRGB(viper.GetString("alerts.discord.webhook.color"))
	if color == nil {
		color = &util.Color{}
	}
	message.Embeds = []discohook.Embed{{
		Description: "`" + session.User() + "` logged in on `" + session.Host() + "`" + " from " + session.IP(),
		Type:        "rich",
		Color:       &discohook.Color{R: color.R, G: color.G, B: color.B},
		Fields:      []discohook.EmbedField{userField, hostField, ipField, timeField}}}
	discohook.Send(viper.GetString("alerts.discord.webhook.url"), &message, false)
}
