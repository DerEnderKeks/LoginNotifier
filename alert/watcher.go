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
	"github.com/spf13/viper"
	"github.com/hpcloud/tail"
	"github.com/DerEnderKeks/LoginNotifier/log"
	"github.com/DerEnderKeks/LoginNotifier/parser"
)

func Watch() {
	t, err := tail.TailFile(viper.GetString("source_log"), tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}})
	defer t.Cleanup()
	for line := range t.Lines {
		session := parser.ParseLine(line.Text)
		if session != nil {
			Alert(*session)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
