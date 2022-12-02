/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package logging

import (
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

//goland:noinspection GoUnusedConst
const (
	// LogPrefixDiscord - log output format for logs sent from this package
	LogPrefixDiscord = "[DISCORD  ]"
)

var (
	log = logrus.Logger{
		Out: os.Stderr,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %msg%\n",
		},
		ReportCaller: true,
	}
)

// LogLevel - The level of logging to allow in the console
var LogLevel int

func init() {
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		log.SetLevel(logrus.DebugLevel)

	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

// Traceln - Logs the event at the Trace level
//
//goland:noinspection SpellCheckingInspection
func Traceln(args ...interface{}) {
	if LogLevel == 0 {
		log.Traceln(args...)
	}
}

// Tracef - Logs the event at the Traceln level with formatting
//
//goland:noinspection GoUnusedExportedFunction
func Tracef(format string, args ...interface{}) {
	if LogLevel == 0 {
		log.Tracef(format, args...)
	}
}

// Debugln - Logs the event at the Debug level
//
//goland:noinspection GoUnusedExportedFunction
func Debugln(args ...interface{}) {
	if LogLevel <= 1 {
		log.Debugln(args...)
	}
}

// Debugf - Logs the event at the Debug level
//
//goland:noinspection GoUnusedExportedFunction
func Debugf(format string, args ...interface{}) {
	if LogLevel <= 1 {
		log.Debugf(format, args...)
	}
}

// Infoln - Logs the event at the Infoln level
func Infoln(args ...interface{}) {
	if LogLevel <= 2 {
		log.Infoln(args...)
	}
}

// Warnln - Logs the event at the Warning level
func Warnln(args ...interface{}) {
	if LogLevel <= 3 {
		log.Warningln(args...)
	}
}

// Warnf - Logs the event at the Warning level with formatting
//
//goland:noinspection GoUnusedExportedFunction
func Warnf(format string, args ...interface{}) {
	if LogLevel == 0 {
		log.Warnf(format, args...)
	}
}

// Errorln - Logs the event at the Error level
func Errorln(args ...interface{}) {
	if LogLevel <= 4 {
		log.Errorln(args...)
	}
}

// Fatalln - Logs the event at the Fatal level
//
//goland:noinspection GoUnusedExportedFunction
func Fatalln(args ...interface{}) {
	if LogLevel <= 5 {
		log.Fatalln(args...)
	}
}
