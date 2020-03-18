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

package log

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/spf13/viper"
	"os"
	"time"
)

type logLevel struct {
	ID   int
	Name string
}

var (
	LevelError   = logLevel{1, "ERROR"}
	LevelWarning = logLevel{2, "WARNING"}
	LevelInfo    = logLevel{3, "INFO"}
	LevelDebug   = logLevel{4, "DEBUG"}
)

func printLog(message interface{}, level logLevel) {
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05.999Z"), "["+level.Name+"]", message)
}

func Panic(err error) {
	raven.CaptureErrorAndWait(err, nil)
	if err != nil && viper.GetInt("loglevel") >= LevelError.ID {
		printLog(err, LevelError)
		panic(err)
	}
}

func Fatal(err error) {
	raven.CaptureErrorAndWait(err, nil)
	if err != nil && viper.GetInt("loglevel") >= LevelError.ID {
		printLog(err, LevelError)
		os.Exit(1)
	}
}

func Warning(args interface{}) {
	if args != nil && viper.GetInt("loglevel") >= LevelWarning.ID {
		printLog(args, LevelWarning)
	}
}

func Info(args interface{}) {
	if args != nil && viper.GetInt("loglevel") >= LevelInfo.ID {
		printLog(args, LevelInfo)
	}
}

func Debug(args interface{}) {
	if args != nil && viper.GetInt("loglevel") >= LevelDebug.ID {
		printLog(args, LevelDebug)
	}
}
