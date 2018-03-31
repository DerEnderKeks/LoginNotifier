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

package main

import (
	"github.com/DerEnderKeks/LoginNotifier/config"
	"github.com/DerEnderKeks/LoginNotifier/alert"
	"github.com/DerEnderKeks/LoginNotifier/log"
)

func init() {
	config.Init()
	log.Info("Config loaded.")
}

func main() {
	alert.Watch()
}
