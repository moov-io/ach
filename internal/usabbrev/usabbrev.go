// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package usabbrev

import (
	"strings"
)

var (
	// usStateAndTerritoryAbbreviations is a mapping of all commonly accepted
	// two-letter abbreviations.
	// Source: https://www.hud.gov/sites/documents/22211X2CHCH.PDF
	usStateAndTerritoryAbbreviations = map[string]bool{
		"AK": true, // Alaska
		"AL": true, // Alabama
		"AR": true, // Arkansas
		"AS": true, // American Samoa
		"AZ": true, // Arizona
		"CA": true, // California
		"CO": true, // Colorado
		"CT": true, // Connecticut
		"DC": true, // District of Columbia
		"DE": true, // Delaware
		"FL": true, // Florida
		// "FT": true, // Foreign (for CTS use)
		"GA": true, // Georgia
		"GU": true, // Guam
		"HI": true, // Hawaii
		"IA": true, // Iowa
		"ID": true, // Idaho
		"IL": true, // Illinois
		"IN": true, // Indiana
		"KS": true, // Kansas
		"KY": true, // Kentucky
		"LA": true, // Louisiana
		"MA": true, // Massachusetts
		"MD": true, // Maryland
		"ME": true, // Maine
		"MI": true, // Michigan
		"MN": true, // Minnesota
		"MO": true, // Missouri
		"MP": true, // Northern Marianas Islands
		"MS": true, // Mississippi
		"MT": true, // Montana
		"NC": true, // North Carolina
		"ND": true, // North Dakota
		"NE": true, // Nebraska
		"NH": true, // New Hampshire
		"NJ": true, // New Jersey
		"NM": true, // New Mexico
		"NV": true, // Nevada
		"NY": true, // New York
		"OH": true, // Ohio
		"OK": true, // Oklahoma
		"OR": true, // Oregon
		"PA": true, // Pennsylvania
		"PR": true, // Puerto Rico
		"RI": true, // Rhode Island
		"SC": true, // South Carolina
		"SD": true, // South Dakota
		"TN": true, // Tennessee
		"TX": true, // Texas
		"UT": true, // Utah
		"VA": true, // Virginia
		"VI": true, // Virgin Islands
		"VT": true, // Vermont
		"WA": true, // Washington
		"WI": true, // Wisconsin
		"WV": true, // West Virginia
		"WY": true, // Wyoming
	}
)

// Valid returns true if code is a valid abbreviation for a United States territory or state.
// Example: UT (Utah) or GU (Guam)
func Valid(code string) bool {
	_, ok := usStateAndTerritoryAbbreviations[strings.ToUpper(code)]
	return ok
}
