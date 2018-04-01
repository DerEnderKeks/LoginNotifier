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
	"github.com/spf13/viper"
	"github.com/DerEnderKeks/LoginNotifier/parser"
	"github.com/DerEnderKeks/LoginNotifier/log"
	"github.com/DerEnderKeks/LoginNotifier/util"
)

func Alert(session parser.Session) {
	if !viper.GetBool("blacklist.whitelist") == util.Contains(viper.GetStringSlice("blacklist.users"), session.User()) {
		log.Info("User '" + session.User() + "' is " + util.Ternary(viper.GetBool("blacklist.whitelist"), "not white", "black").(string) + "listed. Skipping alert.")
		return
	}
	log.Info("Sending alerts for user '" + session.User() + "'...")

	if viper.GetBool("alerts.slack.enabled") {
		log.Debug("Sending Slack alert...")
		slackAlert(session)
	}
	if viper.GetBool("alerts.discord.enabled") {
		log.Debug("Sending Discord alert...")
		discordAlert(session)
	}
}
