/*
 * Copyright (C) 2020 DerEnderKeks
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

package util

import (
	"github.com/DerEnderKeks/LoginNotifier/log"
	"github.com/pkg/errors"
	"net"
	"regexp"
)

func TestAllRegexs(s []string) {
	for _, el := range s {
		_, err := regexp.Compile(el)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to compile regex '"+el+"'"))
		}
	}
}

func AnyRegexMatch(patternList []string, target string) (bool, error) {
	for _, el := range patternList {
		re, err := regexp.Compile(el)
		if err != nil {
			return false, err
		}
		if re.Match([]byte(target)) {
			return true, nil
		}
	}
	return false, nil
}

func TestAllNets(netList []string) {
	for _, el := range netList {
		_, _, err := net.ParseCIDR(el)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to parse net '"+el+"'"))
		}
	}
}

func AnyIPMatch(netList []string, ipString string) (bool, error) {
	targetIp := net.ParseIP(ipString)
	if targetIp == nil {
		return false, errors.New("failed to parse ip")
	}
	for _, el := range netList {
		_, n, err := net.ParseCIDR(el)
		if err != nil {
			return false, err
		}
		if n.Contains(targetIp) {
			return true, nil
		}
	}
	return false, nil
}
