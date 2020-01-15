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
	"github.com/DerEnderKeks/LoginNotifier/log"
	"github.com/DerEnderKeks/LoginNotifier/parser"
	"github.com/hpcloud/tail"
	"github.com/spf13/viper"
	"strings"
)

func Watch() {
	log.Info("Watching file '" + viper.GetString("source_log") + "'...")
	tailLogger := tail.DiscardingLogger
	if viper.GetInt("loglevel") > 3 {
		tailLogger = tail.DefaultLogger
	}
	t, err := tail.TailFile(viper.GetString("source_log"), tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}, Logger: tailLogger})
	if err != nil {
		log.Fatal(err)
	}
	defer t.Cleanup()
	for line := range t.Lines {
		log.Debug("New line: '" + strings.TrimRight(line.Text, "\r\n") + "'")
		session := parser.ParseLine(line.Text)
		if session != nil {
			Alert(*session)
		}
	}
}
