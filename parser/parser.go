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

package parser

import (
	"regexp"
	"time"
	"github.com/DerEnderKeks/LoginNotifier/log"
)

type Session struct {
	user string
	ip   string
	host string
	time time.Time
}

func (s *Session) SetUser(user string) {
	s.user = user
}

func (s Session) User() string {
	return s.user
}

func (s *Session) SetIP(ip string) {
	s.ip = ip
}

func (s Session) IP() string {
	return s.ip
}

func (s *Session) SetHost(host string) {
	s.host = host
}

func (s Session) Host() string {
	return s.host
}

func (s *Session) SetTime(time time.Time) {
	s.time = time
}

func (s Session) Time() time.Time {
	return s.time
}

func (s Session) String() string {
	return s.User() + ", " + s.IP() + ", " + s.Host() + ", " + s.Time().Format(time.ANSIC)
}

var pubkeyRegex = regexp.MustCompile(`^(?P<time>([A-Z][a-z]{2}) [0-9]{2} ([0-9]{2}:){2}[0-9]{2}) (?P<host>\w+) (\w+\[[0-9]+]:) (Accepted publickey for) (?P<user>\w+) (from) (?P<ip>(\w\.?)+)+`)

func ParseLine(line string) (*Session) {
	var session Session
	match := pubkeyRegex.FindStringSubmatch(line)
	if match == nil {
		return nil
	}
	for i, name := range pubkeyRegex.SubexpNames() {
		if i != 0 && name != "" {
			switch name {
			case "time":
				sessiontime, _ := time.Parse(time.Stamp, match[i])
				session.SetTime(sessiontime)
			case "host":
				session.SetHost(match[i])
			case "user":
				session.SetUser(match[i])
			case "ip":
				session.SetIP(match[i])
			}
		}
	}
	log.Debug("New Session: " + session.String())
	return &session
}
