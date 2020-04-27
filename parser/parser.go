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

package parser

import (
	"github.com/DerEnderKeks/LoginNotifier/log"
	"net"
	"regexp"
	"time"
)

type Session struct {
	user        string
	ip          string
	reverseHost string
	host        string
	time        time.Time
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

func (s *Session) SetReverseHost(reverseHost string) {
	s.reverseHost = reverseHost
}

func (s Session) ReverseHost() string {
	return s.reverseHost
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
	return s.User() + ", " + s.IP() + ", " + s.Host() + ", " + s.Time().Format(time.RFC3339)
}

var pubkeyRegex = regexp.MustCompile(`^(?P<time>(([A-Z][a-z]{2})\s+[0-9]{1,2}\s+([0-9]{2}:?){3})|\d{4}-(\d{2}-?){2}T(\d{2}:?){3}.\d+[+-Z]((\d{2}:?){2})?)\s+(?P<host>[\w.-]+)\s+(\w+\[[0-9]+]:)\s+Accepted publickey for (?P<user>[\w.-]+) from (?P<ip>((:?[A-Fa-f0-9]{0,4})+|([0-9]{1,3}\.?)+))\s`)

func ParseLine(line string) *Session {
	var session Session
	match := pubkeyRegex.FindStringSubmatch(line)
	if match == nil {
		return nil
	}
	for i, name := range pubkeyRegex.SubexpNames() {
		if i != 0 && name != "" {
			switch name {
			case "time":
				sessionTime, err := time.Parse(time.Stamp, match[i])
				if err != nil {
					sessionTime, _ = time.Parse(time.RFC3339, match[i])
				}
				session.SetTime(sessionTime)
			case "host":
				session.SetHost(match[i])
			case "user":
				session.SetUser(match[i])
			case "ip":
				session.SetIP(match[i])
				reverseHosts, err := net.LookupAddr(session.IP())
				if err != nil {
					log.Warning(err)
				} else if len(reverseHosts) > 0 {
					re := regexp.MustCompile(`\.$`) // remove trailing dot
					session.SetReverseHost(re.ReplaceAllString(reverseHosts[0], ""))
				}
			}
		}
	}
	log.Debug("New Session: " + session.String())
	return &session
}
