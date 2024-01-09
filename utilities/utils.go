/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package utilities

import (
	"fmt"
	"runtime"
)

// Contains - helper function to determine if a slice contains a particular value
//
//goland:noinspection GoUnusedExportedFunction
func Contains[T comparable](slice []T, e T) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}

	return false
}

//goland:noinspection GoUnusedExportedFunction
func FuncName() string {
	pc, _, line, _ := runtime.Caller(1)
	return fmt.Sprintf("(%s:L%d)", runtime.FuncForPC(pc).Name(), line)
}

//goland:noinspection GoUnusedExportedFunction
func ToPtr[T any](p T) *T {
	return &p
}
