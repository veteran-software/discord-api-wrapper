/*
 * Copyright (c) 2023. Veteran Software
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

package gateway

// Gateway Close Event Codes
//
// In order to prevent broken reconnect loops, you should consider some close codes as a signal to stop reconnecting.
// This can be because your token expired, or your identification is invalid.
// This table explains what the application defined close codes for the gateway are, and which close codes you should not attempt to reconnect.

//goland:noinspection GoUnusedExportedFunction
func GetCloseCode(c int) (code int, reconnect bool, description string) {
	switch c {
	case 4000:
		return 4000, true, "Unknown Error"
	case 4001:
		return 4001, true, "Unknown opcode"
	case 4002:
		return 4002, true, "Decode Error"
	case 4003:
		return 4003, true, "Not Authenticated"
	case 4004:
		return 4004, false, "Authentication Failed"
	case 4005:
		return 4005, true, "Already Authenticated"
	case 4007:
		return 4007, true, "Invalid Sequence Number"
	case 4008:
		return 4008, true, "Rate Limited"
	case 4009:
		return 4009, true, "Session Timed Out"
	case 4010:
		return 4010, false, "Invalid Shard"
	case 4011:
		return 4011, false, "Sharding Required"
	case 4012:
		return 4012, false, "Invalid API Version"
	case 4013:
		return 4013, false, "Invalid Intent(s)"
	case 4014:
		return 4014, false, "Disallowed Intent(s)"
	default:
		return 0, false, "Unknown Close Code"
	}
}
