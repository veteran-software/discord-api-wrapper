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

package api

/*
ApplicationRoleConnectionMetadata

A representation of role connection metadata for an Application.

When a Guild has added a bot and that bot has configured its RoleConnectionsVerificationURL (in the developer portal), the application will render as a potential verification method in the guild's role verification configuration.

If an application has configured role connection metadata, its metadata will appear in the role verification configuration when the application has been added as a verification method to the role.

When a user connects their account using the bots RoleConnectionsVerificationURL, the bot will update a user's role connection with metadata using the OAuth2 scopes.RoleConnectionsWrite scope.
*/
type ApplicationRoleConnectionMetadata struct {
	Type                      ApplicationRoleConnectionMetadataType // type of metadata value
	Key                       string                                // dictionary key for the metadata field (must be a-z, 0-9, or _ characters; 1-50 characters)
	Name                      string                                // name of the metadata field (1-100 characters)
	NameLocalizations         LocalizationDict                      // translations of the name
	Description               string                                //	description of the metadata field (1-200 characters)
	DescriptionsLocalizations LocalizationDict                      // translations of the description
}

// ApplicationRoleConnectionMetadataType - type of metadata value
type ApplicationRoleConnectionMetadataType int

//goland:noinspection GoUnusedConst
const (
	// IntegerLessThanOrEqual - the metadata value (integer) is less than or equal to the guild's configured value (integer)
	IntegerLessThanOrEqual ApplicationRoleConnectionMetadataType = iota + 1

	// IntegerGreaterThanOrEqual - the metadata value (integer) is greater than or equal to the guild's configured value (integer)
	IntegerGreaterThanOrEqual

	// IntegerEqual - the metadata value (integer) is equal to the guild's configured value (integer)
	IntegerEqual

	// IntegerNotEqual - the metadata value (integer) is not equal to the guild's configured value (integer)
	IntegerNotEqual

	// DateTimeLessThanOrEqual - the metadata value (ISO8601 string) is less than or equal to the guild's configured value (integer; days before current date)
	DateTimeLessThanOrEqual

	// DateTimeGreaterThanOrEqual - the metadata value (ISO8601 string) is greater than or equal to the guild's configured value (integer; days before current date)
	DateTimeGreaterThanOrEqual

	// BooleanEqual - the metadata value (integer) is equal to the guild's configured value (integer; 1)
	BooleanEqual

	// BooleanNotEqual - the metadata value (integer) is not equal to the guild's configured value (integer; 1)
	BooleanNotEqual
)
