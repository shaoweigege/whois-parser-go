/*
 * Copyright 2014-2020 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Go module for domain whois info parse
 * https://www.likexian.com/
 */

package whoisparser

var (
	// keyRule is parser key rule mapper
	keyRule = map[string]string{
		"id":                                     "domain_id",
		"roid":                                   "domain_id",
		"domain id":                              "domain_id",
		"domain":                                 "domain_name",
		"domain name":                            "domain_name",
		"status":                                 "domain_status",
		"state":                                  "domain_status",
		"domain status":                          "domain_status",
		"registration status":                    "domain_status",
		"query status":                           "domain_status",
		"dnssec":                                 "domain_dnssec",
		"domain dnssec":                          "domain_dnssec",
		"registrar dnssec":                       "domain_dnssec",
		"signing key":                            "domain_dnssec",
		"domain signed":                          "domain_dnssec",
		"whois":                                  "whois_server",
		"whois server":                           "whois_server",
		"registrar whois server":                 "whois_server",
		"nserver":                                "name_servers",
		"name server":                            "name_servers",
		"name servers":                           "name_servers",
		"nameserver":                             "name_servers",
		"nameservers":                            "name_servers",
		"name servers information":               "name_servers",
		"host name":                              "name_servers",
		"domain nameservers":                     "name_servers",
		"domain name servers":                    "name_servers",
		"domain servers in listed order":         "name_servers",
		"created":                                "created_date",
		"registered":                             "created_date",
		"created on":                             "created_date",
		"create date":                            "created_date",
		"created date":                           "created_date",
		"creation date":                          "created_date",
		"domain registration date":               "created_date",
		"registration date":                      "created_date",
		"domain create date":                     "created_date",
		"domain name commencement date":          "created_date",
		"registered date":                        "created_date",
		"registered on":                          "created_date",
		"registration time":                      "created_date",
		"first registration date":                "created_date",
		"domain record activated":                "created_date",
		"record created on":                      "created_date",
		"domain registered":                      "created_date",
		"modified":                               "updated_date",
		"changed":                                "updated_date",
		"update date":                            "updated_date",
		"updated date":                           "updated_date",
		"updated on":                             "updated_date",
		"last update":                            "updated_date",
		"last updated":                           "updated_date",
		"last updated on":                        "updated_date",
		"last modified":                          "updated_date",
		"last updated date":                      "updated_date",
		"domain last updated date":               "updated_date",
		"domain record last updated":             "updated_date",
		"domain datelastmodified":                "updated_date",
		"expire":                                 "expired_date",
		"expires":                                "expired_date",
		"expires on":                             "expired_date",
		"paid till":                              "expired_date",
		"expire date":                            "expired_date",
		"expired date":                           "expired_date",
		"expiration date":                        "expired_date",
		"expiration on":                          "expired_date",
		"registrar registration expiration date": "expired_date",
		"domain expiration date":                 "expired_date",
		"expiry date":                            "expired_date",
		"expiration time":                        "expired_date",
		"domain expires":                         "expired_date",
		"record expires on":                      "expired_date",
		"record will expire on":                  "expired_date",
		"registrar www":                          "referral_url",
		"referral url":                           "referral_url",
		"registrar url":                          "referral_url",
		"registrar web":                          "referral_url",
		"registrar website":                      "referral_url",
		"registration service url":               "referral_url",
		"registrant c":                           "registrant_id",
		"registrant id":                          "registrant_id",
		"registrant iana id":                     "registrant_id",
		"registrant contact id":                  "registrant_id",
		"registrant name":                        "registrant_name",
		"registrant person":                      "registrant_name",
		"registrant contact":                     "registrant_name",
		"registrant contact name":                "registrant_name",
		"registrant given name":                  "registrant_name",
		"registrant holder name":                 "registrant_name",
		"registrant holder english name":         "registrant_name",
		"registrant service provider":            "registrant_name",
		"registrant org":                         "registrant_organization",
		"registrant organization":                "registrant_organization",
		"registrant contact organization":        "registrant_organization",
		"registrant organisation":                "registrant_organization",
		"registrant contact organisation":        "registrant_organization",
		"registrant company name":                "registrant_organization",
		"registrant company english name":        "registrant_organization",
		"registrant address":                     "registrant_street",
		"registrant address1":                    "registrant_street",
		"registrant street":                      "registrant_street",
		"registrant street1":                     "registrant_street",
		"registrant contact address":             "registrant_street",
		"registrant contact address1":            "registrant_street",
		"registrant contact street":              "registrant_street",
		"registrant contact street1":             "registrant_street",
		"registrant s address":                   "registrant_street",
		"registrant s address1":                  "registrant_street",
		"registrant postal address":              "registrant_street",
		"registrant postal address1":             "registrant_street",
		"registrant city":                        "registrant_city",
		"registrant contact city":                "registrant_city",
		"registrant state province":              "registrant_state_province",
		"registrant contact state province":      "registrant_state_province",
		"registrant zipcode":                     "registrant_postal_code",
		"registrant zip code":                    "registrant_postal_code",
		"registrant postal code":                 "registrant_postal_code",
		"registrant contact postal code":         "registrant_postal_code",
		"registrant country":                     "registrant_country",
		"registrant country economy":             "registrant_country",
		"registrant contact country":             "registrant_country",
		"registrant phone":                       "registrant_phone",
		"registrant phone number":                "registrant_phone",
		"registrant contact phone":               "registrant_phone",
		"registrant contact phone number":        "registrant_phone",
		"registrant abuse contact phone":         "registrant_phone",
		"registrant phone ext":                   "registrant_phone_ext",
		"registrant contact phone ext":           "registrant_phone_ext",
		"registrant fax":                         "registrant_fax",
		"registrant fax no":                      "registrant_fax",
		"registrant fax number":                  "registrant_fax",
		"registrant facsimile":                   "registrant_fax",
		"registrant facsimile number":            "registrant_fax",
		"registrant contact fax":                 "registrant_fax",
		"registrant contact fax number":          "registrant_fax",
		"registrant contact facsimile":           "registrant_fax",
		"registrant contact facsimile number":    "registrant_fax",
		"registrant fax ext":                     "registrant_fax_ext",
		"registrant contact fax ext":             "registrant_fax_ext",
		"registrant mail":                        "registrant_email",
		"registrant email":                       "registrant_email",
		"registrant e mail":                      "registrant_email",
		"registrant contact mail":                "registrant_email",
		"registrant contact email":               "registrant_email",
		"registrant contact e mail":              "registrant_email",
		"registrant abuse contact email":         "registrant_email",
	}
)
