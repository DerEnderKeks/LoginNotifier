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

package alert

import (
	"github.com/DerEnderKeks/LoginNotifier/parser"
	"github.com/DerEnderKeks/LoginNotifier/log"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/spf13/viper"
	"time"
)

func SlackAlert(session parser.Session) {
	attachment := slack.Attachment{}
	attachment.AddField(
		slack.Field{Title: "User", Value: session.User(), Short: true}).AddField(
		slack.Field{Title: "IP", Value: session.IP(), Short: true}).AddField(
		slack.Field{Title: "Host", Value: session.Host(), Short: true}).AddField(
		slack.Field{Title: "Time", Value: session.Time().Format(time.Stamp), Short: true})
	color := "#ffcc00"
	attachment.Color = &color
	payload := slack.Payload{
		Text:        "`" + session.User() + "` logged in on `" + session.Host() + "`" + " from <https://whois.domaintools.com/" + session.IP() + "|" + session.IP() + ">",
		Username:    viper.GetString("alerts.slack.webhook.username"),
		Channel:     viper.GetString("alerts.slack.webhook.channel"),
		IconEmoji:   viper.GetString("alerts.slack.webhook.icon_emoji"),
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(viper.GetString("alerts.slack.webhook.url"), "", payload)
	if len(err) > 0 {
		log.Warning(err)
	}
}
