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

package util

import (
	"regexp"
	"strconv"
)

type Color struct {
	R, G, B uint8
}

func HexToRGB(color string) (*Color) {
	hexRegex := regexp.MustCompile(`(?i)^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$`)
	result := hexRegex.FindStringSubmatch(color)
	if len(result) < 4 {
		return nil
	}
	r, err := strconv.ParseUint(result[1], 16, 8)
	g, err := strconv.ParseUint(result[2], 16, 8)
	b, err := strconv.ParseUint(result[3], 16, 8)
	if err != nil {
		return nil
	}
	return &Color{R: uint8(r), G: uint8(g), B: uint8(b)}
}
