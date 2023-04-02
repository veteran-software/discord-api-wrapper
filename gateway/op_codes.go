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

// OpCode
//
// All gateway events in Discord are tagged with an opcode that denotes the payload type.
// Your connection to our gateway may also sometimes close.
// When it does, you will receive a close code that tells you what happened.
type OpCode int

//goland:noinspection GoUnusedConst
const (
	Dispatch OpCode = iota
	Heartbeat
	Identify
	PresenceUpdate
	VoiceStateUpdate
	Resume OpCode = iota + 1
	Reconnect
	RequestGuildMembers
	InvalidSession
	Hello
	HeartbeatAck
)
