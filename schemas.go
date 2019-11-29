package spaceapivalidator

// CommitHash contains the hash of the commit the Validate function validates against
var CommitHash = "68da57382ab211e0f0fa8e9c13c22de7811cac83"

// SpaceAPISchemas load from the repository as a map
var SpaceAPISchemas = map[string]string{
	"11": `{
  "$id": "https://schema.spaceapi.io/11.json",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "description": "SpaceAPI 0.11",
  "type": "object",
  "properties": {
    "api": {
      "description": "The version of SpaceAPI your endpoint uses",
      "type": "string",
      "enum": [
        "0.11"
      ]
    },
    "space": {
      "description": "The name of your space",
      "type": "string"
    },
    "logo": {
      "description": "The space logo",
      "type": "string"
    },
    "url": {
      "description": "The main website",
      "type": "string"
    },
    "address": {
      "description": "The postal address of your space (street, block, housenumber, zip code, city, whatever you usually need in your country, and the country itself)",
      "type": "string"
    },
    "lat": {
      "description": "Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.",
      "type": "number"
    },
    "lon": {
      "description": "Longitude of your space location, in degree with decimal places. Use positive values for locations west of Greenwich, and negative values for locations east of Greenwich.",
      "type": "number"
    },
    "cam": {
      "description": "URL(s) of webcams in your space",
      "type": "array",
      "items": {
        "type": "string"
      },
      "minItems": 1
    },
    "stream": {
      "description": "A mapping of stream types to stream URLs. Example: <samp>{'mp4':'http//example.org/stream.mpg', 'mjpeg':'http://example.org/stream.mjpeg'}</samp>",
      "type": "object"
    },
    "open": {
      "description": "A boolean which indicates if the space is currently open",
      "type": "boolean"
    },
    "status": {
      "description": "An additional free-form string, could be something like <samp>'open for public'</samp>, <samp>'members only'</samp> or whatever you want it to be",
      "type": "string"
    },
    "lastchange": {
      "description": "The Unix timestamp when the space status changed most recently",
      "type": "number"
    },
    "events": {
      "description": "Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'",
      "type": "array",
      "items": {
        "required": [
          "name",
          "type",
          "t"
        ],
        "type": "object",
        "properties": {
          "name": {
            "description": "Name or other identity of the subject (e.g. <samp>J. Random Hacker</samp>, <samp>fridge</samp>, <samp>3D printer</samp>, \u2026)",
            "type": "string"
          },
          "type": {
            "description": "Action (e.g. <samp>check-in</samp>, <samp>check-out</samp>, <samp>finish-print</samp>, \u2026). Define your own actions and use them consistently, canonical actions are not (yet) specified",
            "type": "string"
          },
          "t": {
            "description": "Unix timestamp when the event occurred",
            "type": "number"
          },
          "extra": {
            "description": "A custom text field to give more information about the event",
            "type": "string"
          }
        }
      }
    },
    "contact": {
      "description": "Contact information about your space",
      "type": "object",
      "properties": {
        "phone": {
          "description": "Phone number, including country code with a leading plus sign. Example: <samp>+1 800 555 4567</samp>",
          "type": "string"
        },
        "sip": {
          "description": "URI for Voice-over-IP via SIP. Example: <samp>sip:yourspace@sip.example.org</samp>",
          "type": "string"
        },
        "keymaster": {
          "description": "Phone numbers of people who carry a key and are able to open the space upon request. Example: <samp>['+1 800 555 4567','+1 800 555 4544']</samp>",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "irc": {
          "description": "URL of the IRC channel, in the form <samp>irc://example.org/#channelname</samp>",
          "type": "string"
        },
        "twitter": {
          "description": "Twitter handle, with leading @",
          "type": "string"
        },
        "email": {
          "description": "E-mail address for contacting your space",
          "type": "string"
        },
        "ml": {
          "description": "The e-mail address of your mailing list. If you use Google Groups then the e-mail looks like <samp>your-group@googlegroups.com</samp>.",
          "type": "string"
        },
        "jabber": {
          "description": "A public Jabber/XMPP multi-user chatroom in the form <samp>chatroom@conference.example.net</samp>",
          "type": "string"
        }
      }
    },
    "icon": {
      "description": "Icons that show the status graphically",
      "type": "object",
      "properties": {
        "open": {
          "description": "The URL to your customized space logo showing an open space",
          "type": "string"
        },
        "closed": {
          "description": "The URL to your customized space logo showing a closed space",
          "type": "string"
        }
      },
      "required": [
        "open",
        "closed"
      ]
    }
  },
  "required": [
    "api",
    "space",
    "logo",
    "url",
    "open",
    "icon"
  ]
}
`,
	"12": `{
  "$id": "https://schema.spaceapi.io/12.json",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "description": "SpaceAPI 0.12",
  "type": "object",
  "properties": {
    "api": {
      "description": "The version of SpaceAPI your endpoint uses",
      "type": "string",
      "enum": [
        "0.12"
      ]
    },
    "space": {
      "description": "The name of your space",
      "type": "string"
    },
    "logo": {
      "description": "The space logo",
      "type": "string"
    },
    "url": {
      "description": "The main website",
      "type": "string"
    },
    "address": {
      "description": "The postal address of your space (street, block, housenumber, zip code, city, whatever you usually need in your country, and the country itself)",
      "type": "string"
    },
    "lat": {
      "description": "Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.",
      "type": "number"
    },
    "lon": {
      "description": "Longitude of your space location, in degree with decimal places. Use positive values for locations west of Greenwich, and negative values for locations east of Greenwich.",
      "type": "number"
    },
    "cam": {
      "description": "URL(s) of webcams in your space",
      "type": "array",
      "items": {
        "type": "string"
      },
      "minItems": 1
    },
    "stream": {
      "description": "A mapping of stream types to stream URLs. Example: <samp>{'mp4':'http//example.org/stream.mpg', 'mjpeg':'http://example.org/stream.mjpeg'}</samp>",
      "type": "object"
    },
    "open": {
      "description": "A boolean which indicates if the space is currently open",
      "type": "boolean"
    },
    "status": {
      "description": "An additional free-form string, could be something like <samp>'open for public'</samp>, <samp>'members only'</samp> or whatever you want it to be",
      "type": "string"
    },
    "lastchange": {
      "description": "The Unix timestamp when the space status changed most recently",
      "type": "number"
    },
    "events": {
      "description": "Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'",
      "type": "array",
      "items": {
        "required": [
          "name",
          "type",
          "t"
        ],
        "type": "object",
        "properties": {
          "name": {
            "description": "Name or other identity of the subject (e.g. <samp>J. Random Hacker</samp>, <samp>fridge</samp>, <samp>3D printer</samp>, \u2026)",
            "type": "string"
          },
          "type": {
            "description": "Action (e.g. <samp>check-in</samp>, <samp>check-out</samp>, <samp>finish-print</samp>, \u2026). Define your own actions and use them consistently, canonical actions are not (yet) specified",
            "type": "string"
          },
          "t": {
            "description": "Unix timestamp when the event occurred",
            "type": "number"
          },
          "extra": {
            "description": "A custom text field to give more information about the event",
            "type": "string"
          }
        }
      }
    },
    "contact": {
      "description": "Contact information about your space",
      "type": "object",
      "properties": {
        "phone": {
          "description": "Phone number, including country code with a leading plus sign. Example: <samp>+1 800 555 4567</samp>",
          "type": "string"
        },
        "sip": {
          "description": "URI for Voice-over-IP via SIP. Example: <samp>sip:yourspace@sip.example.org</samp>",
          "type": "string"
        },
        "keymaster": {
          "description": "Phone numbers of people who carry a key and are able to open the space upon request. Example: <samp>['+1 800 555 4567','+1 800 555 4544']</samp>",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "irc": {
          "description": "URL of the IRC channel, in the form <samp>irc://example.org/#channelname</samp>",
          "type": "string"
        },
        "twitter": {
          "description": "Twitter handle, with leading @",
          "type": "string"
        },
        "email": {
          "description": "E-mail address for contacting your space",
          "type": "string"
        },
        "ml": {
          "description": "The e-mail address of your mailing list. If you use Google Groups then the e-mail looks like <samp>your-group@googlegroups.com</samp>.",
          "type": "string"
        },
        "jabber": {
          "description": "A public Jabber/XMPP multi-user chatroom in the form <samp>chatroom@conference.example.net</samp>",
          "type": "string"
        }
      }
    },
    "icon": {
      "description": "Icons that show the status graphically",
      "type": "object",
      "properties": {
        "open": {
          "description": "The URL to your customized space logo showing an open space",
          "type": "string"
        },
        "closed": {
          "description": "The URL to your customized space logo showing a closed space",
          "type": "string"
        }
      },
      "required": [
        "open",
        "closed"
      ]
    },
    "sensors": {
      "description": "Data of various sensors in your space (e.g. temperature, humidity, amount of Club-Mate left, \u2026). The only canonical property is the <em>temp</em> property, additional sensor types may be defined by you. In this case, you are requested to share your definition for inclusion in this specification.",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "temp": {
            "description": "A mapping of measuring locations to temperature values. The values should match the basic regular expression <code>^[+-]?[0-9]+(\\.[0-9]+)?[FCK]$</code>. Example: <samp>{'kitchen':'48F', 'storage':'-273.1K'}</samp>",
            "type": "object"
          }
        }
      }
    },
    "feeds": {
      "description": "Feeds where users can get updates of your space",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": {
            "description": "A mnemonic identifier, like <samp>wiki</samp>, <samp>blog</samp>, etc.",
            "type": "string"
          },
          "type": {
            "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
            "type": "string"
          },
          "url": {
            "description": "Feed URL",
            "type": "string"
          }
        },
        "required": [
          "name",
          "url"
        ]
      }
    }
  },
  "required": [
    "api",
    "space",
    "logo",
    "url",
    "open",
    "icon"
  ]
}
`,
	"13": `{
  "$id": "https://schema.spaceapi.io/13.json",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "description": "SpaceAPI 0.13",
  "type": "object",
  "properties": {
    "api": {
      "description": "The version of SpaceAPI your endpoint uses",
      "type": "string",
      "enum": [
        "0.13"
      ]
    },
    "space": {
      "description": "The name of your space",
      "type": "string"
    },
    "logo": {
      "description": "URL to your space logo",
      "type": "string"
    },
    "url": {
      "description": "URL to your space website",
      "type": "string"
    },
    "location": {
      "description": "Position data such as a postal address or geographic coordinates",
      "type": "object",
      "properties": {
        "address": {
          "description": "The postal address of your space (street, block, housenumber, zip code, city, whatever you usually need in your country, and the country itself).<br>Examples: <ul><li>Netzladen e.V., Breite Stra\u00dfe 74, 53111 Bonn, Germany</li></ul>",
          "type": "string"
        },
        "lat": {
          "description": "Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.",
          "type": "number"
        },
        "lon": {
          "description": "Longitude of your space location, in degree with decimal places. Use positive values for locations east of Greenwich, and negative values for locations west of Greenwich.",
          "type": "number"
        }
      },
      "required": [
        "lat",
        "lon"
      ]
    },
    "spacefed": {
      "description": "A flag indicating if the hackerspace uses SpaceFED, a federated login scheme so that visiting hackers can use the space WiFi with their home space credentials.",
      "type": "object",
      "properties": {
        "spacenet": {
          "description": "See the <a target=\"_blank\" href=\"https://spacefed.net/wiki/index.php/Category:Howto/Spacenet\">wiki</a>.",
          "type": "boolean"
        },
        "spacesaml": {
          "description": "See the <a target=\"_blank\" href=\"https://spacefed.net/wiki/index.php/Category:Howto/Spacesaml\">wiki</a>.",
          "type": "boolean"
        },
        "spacephone": {
          "description": "See the <a target=\"_blank\" href=\"https://spacefed.net/wiki/index.php/Category:Howto/Spacephone\">wiki</a>.",
          "type": "boolean"
        }
      },
      "required": [
        "spacenet",
        "spacesaml",
        "spacephone"
      ]
    },
    "cam": {
      "description": "URL(s) of webcams in your space",
      "type": "array",
      "items": {
        "type": "string"
      },
      "minItems": 1
    },
    "stream": {
      "description": "A mapping of stream types to stream URLs. If you use other stream types make a <a href=\"https://github.com/spaceapi/schema/pulls\" target=\"_blank\">pull request</a> or prefix yours with <samp>ext_</samp>.",
      "type": "object",
      "properties": {
        "m4": {
          "description": "Your mpg stream URL. Example: <samp>{\"mp4\": \"http://example.org/stream.mpg\"}</samp>",
          "type": "string"
        },
        "mjpeg": {
          "description": "Your mjpeg stream URL. Example: <samp>{\"mjpeg\": \"http://example.org/stream.mjpeg\"}</samp>",
          "type": "string"
        },
        "ustream": {
          "description": "Your ustream stream URL. Example: <samp>{\"ustream\": \"http://www.ustream.tv/channel/hackspsps\"}</samp>",
          "type": "string"
        }
      }
    },
    "state": {
      "description": "A collection of status-related data: actual open/closed status, icons, last change timestamp etc.",
      "type": "object",
      "properties": {
        "open": {
          "description": "A flag which indicates if the space is currently open or closed. The state 'undefined' can be achieved by assigning this field the value 'null' (without the quotes). In most (all?) programming languages this is evaluated to false so that no app should break",
          "type": [
            "boolean",
            "null"
          ]
        },
        "lastchange": {
          "description": "The Unix timestamp when the space status changed most recently",
          "type": "number"
        },
        "trigger_person": {
          "description": "The person who lastly changed the state e.g. opened or closed the space.",
          "type": "string"
        },
        "message": {
          "description": "An additional free-form string, could be something like <samp>'open for public'</samp>, <samp>'members only'</samp> or whatever you want it to be",
          "type": "string"
        },
        "icon": {
          "description": "Icons that show the status graphically",
          "type": "object",
          "properties": {
            "open": {
              "description": "The URL to your customized space logo showing an open space",
              "type": "string"
            },
            "closed": {
              "description": "The URL to your customized space logo showing a closed space",
              "type": "string"
            }
          },
          "required": [
            "open",
            "closed"
          ]
        }
      },
      "required": [
        "open"
      ]
    },
    "events": {
      "description": "Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'",
      "type": "array",
      "items": {
        "required": [
          "name",
          "type",
          "timestamp"
        ],
        "type": "object",
        "properties": {
          "name": {
            "description": "Name or other identity of the subject (e.g. <samp>J. Random Hacker</samp>, <samp>fridge</samp>, <samp>3D printer</samp>, \u2026)",
            "type": "string"
          },
          "type": {
            "description": "Action (e.g. <samp>check-in</samp>, <samp>check-out</samp>, <samp>finish-print</samp>, \u2026). Define your own actions and use them consistently, canonical actions are not (yet) specified",
            "type": "string"
          },
          "timestamp": {
            "description": "Unix timestamp when the event occurred",
            "type": "number"
          },
          "extra": {
            "description": "A custom text field to give more information about the event",
            "type": "string"
          }
        }
      }
    },
    "contact": {
      "description": "Contact information about your space. You must define at least one which is in the list of allowed values of the issue_report_channels field.",
      "type": "object",
      "properties": {
        "phone": {
          "description": "Phone number, including country code with a leading plus sign. Example: <samp>+1 800 555 4567</samp>",
          "type": "string"
        },
        "sip": {
          "description": "URI for Voice-over-IP via SIP. Example: <samp>sip:yourspace@sip.example.org</samp>",
          "type": "string"
        },
        "keymasters": {
          "description": "Persons who carry a key and are able to open the space upon request. One of the fields irc_nick, phone, email or twitter must be specified.",
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "properties": {
              "name": {
                "description": "Real name",
                "type": "string"
              },
              "irc_nick": {
                "description": "Contact the person with this nickname directly in irc if available. The irc channel to be used is defined in the contact/irc field.",
                "type": "string"
              },
              "phone": {
                "description": "Example: <samp>['+1 800 555 4567','+1 800 555 4544']</samp>",
                "type": "string"
              },
              "email": {
                "description": "Email address which can be base64 encoded.",
                "type": "string"
              },
              "twitter": {
                "description": "Twitter username with leading <samp>@</samp>.",
                "type": "string"
              }
            }
          }
        },
        "irc": {
          "description": "URL of the IRC channel, in the form <samp>irc://example.org/#channelname</samp>",
          "type": "string"
        },
        "twitter": {
          "description": "Twitter handle, with leading @",
          "type": "string"
        },
        "facebook": {
          "description": "Facebook account URL.",
          "type": "string"
        },
        "google": {
          "description": "Google services.",
          "type": "object",
          "properties": {
            "plus": {
              "description": "Google plus URL.",
              "type": "string"
            }
          }
        },
        "identica": {
          "description": "Identi.ca or StatusNet account, in the form <samp>yourspace@example.org</samp>",
          "type": "string"
        },
        "foursquare": {
          "description": "Foursquare ID, in the form <samp>4d8a9114d85f3704eab301dc</samp>.",
          "type": "string"
        },
        "email": {
          "description": "E-mail address for contacting your space. If this is a mailing list consider to use the contact/ml field.",
          "type": "string"
        },
        "ml": {
          "description": "The e-mail address of your mailing list. If you use Google Groups then the e-mail looks like <samp>your-group@googlegroups.com</samp>.",
          "type": "string"
        },
        "jabber": {
          "description": "A public Jabber/XMPP multi-user chatroom in the form <samp>chatroom@conference.example.net</samp>",
          "type": "string"
        },
        "issue_mail": {
          "description": "A separate email address for issue reports (see the <em>issue_report_channels</em> field). This value can be Base64-encoded.",
          "type": "string"
        }
      }
    },
    "issue_report_channels": {
      "description": "This array defines all communication channels where you want to get automated issue reports about your SpaceAPI endpoint from the revalidator. This field is meant for internal usage only and it should never be consumed by any app. At least one channel must be defined. Please consider that when using <samp>ml</samp> the mailing list moderator has to moderate incoming emails or add the sender email to the subscribers. If you don't break your SpaceAPI implementation you won't get any notifications ;-)",
      "type": "array",
      "items": {
        "type": "string",
        "enum": [
          "email",
          "issue_mail",
          "twitter",
          "ml"
        ]
      },
      "minItems": 1
    },
    "sensors": {
      "description": "Data of various sensors in your space (e.g. temperature, humidity, amount of Club-Mate left, \u2026). The only canonical property is the <em>temp</em> property, additional sensor types may be defined by you. In this case, you are requested to share your definition for inclusion in this specification.",
      "type": "object",
      "properties": {
        "temperature": {
          "description": "Temperature sensor. To convert from one unit of temperature to another consider <a href=\"http://en.wikipedia.org/wiki/Temperature_conversion_formulas\" target=\"_blank\">Wikipedia</a>.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value.",
                "type": "string",
                "enum": [
                  "\u00b0C",
                  "\u00b0F",
                  "K",
                  "\u00b0De",
                  "\u00b0N",
                  "\u00b0R",
                  "\u00b0R\u00e9",
                  "\u00b0R\u00f8"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "door_locked": {
          "description": "Sensor type to indicate if a certain door is locked.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "boolean"
              },
              "location": {
                "description": "The location of your sensor such as <samp>front door</samp>, <samp>chill room</samp> or <samp>lab</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "location"
            ]
          }
        },
        "barometer": {
          "description": "Barometer sensor",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                "type": "string",
                "enum": [
                  "hPA"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "radiation": {
          "description": "Compound radiation sensor. Check this <a rel=\"nofollow\" href=\"https://sites.google.com/site/diygeigercounter/gm-tubes-supported\" target=\"_blank\">resource</a>.",
          "type": "object",
          "properties": {
            "alpha": {
              "description": "An alpha sensor",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            },
            "beta": {
              "description": "A beta sensor",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            },
            "gamma": {
              "description": "A gamma sensor",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            },
            "beta_gamma": {
              "description": "A sensor which cannot filter beta and gamma radiation separately.",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            }
          }
        },
        "humidity": {
          "description": "Humidity sensor",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                "type": "string",
                "enum": [
                  "%"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "beverage_supply": {
          "description": "How much Mate and beer is in your fridge?",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit, either <samp>btl</samp> for bottles or <samp>crt</samp> for crates.",
                "type": "string",
                "enum": [
                  "btl",
                  "crt"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Room 1</samp> or <samp>Room 2</samp> or <samp>Room 3</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit"
            ]
          }
        },
        "power_consumption": {
          "description": "The power consumption of a specific device or of your whole space.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                "type": "string",
                "enum": [
                  "mW",
                  "W",
                  "VA"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "wind": {
          "description": "Your wind sensor.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "properties": {
                "description": "",
                "type": "object",
                "properties": {
                  "speed": {
                    "description": "",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "m/s",
                          "km/h",
                          "kn"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  },
                  "gust": {
                    "description": "",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "m/s",
                          "km/h",
                          "kn"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  },
                  "direction": {
                    "description": "The wind direction in degrees.",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "\u00b0"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  },
                  "elevation": {
                    "description": "Height above mean sea level.",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "m"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  }
                },
                "required": [
                  "speed",
                  "gust",
                  "direction",
                  "elevation"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "properties",
              "location"
            ]
          }
        },
        "network_connections": {
          "description": "This sensor type is to specify the currently active  ethernet or wireless network devices. You can create different instances for each network type.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "type": {
                "description": "This field is optional but you can use it to the network type such as <samp>wifi</samp> or <samp>cable</samp>. You can even expose the number of <a href=\"https://spacefed.net/wiki/index.php/Spacenet\" target=\"_blank\">spacenet</a>-authenticated connections.",
                "type": "string",
                "enum": [
                  "wifi",
                  "cable",
                  "spacenet"
                ]
              },
              "value": {
                "description": "The amount of network connections.",
                "type": "number"
              },
              "machines": {
                "description": "The machines that are currently connected with the network.",
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "name": {
                      "description": "The machine name.",
                      "type": "string"
                    },
                    "mac": {
                      "description": "The machine's MAC address of the format <samp>D3:3A:DB:EE:FF:00</samp>.",
                      "type": "string"
                    }
                  },
                  "required": [
                    "mac"
                  ]
                }
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value"
            ]
          }
        },
        "account_balance": {
          "description": "How rich is your hackerspace?",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "How much?",
                "type": "number"
              },
              "unit": {
                "description": "What's the currency? Please use the ones provided, in the next version you can use currency definitions according to <a href=\"https://en.wikipedia.org/wiki/ISO_4217\" target=\"_blank\">ISO 4217</a>",
                "type": "string",
                "enum": [
                  "BTC",
                  "EUR",
                  "USD",
                  "GBP"
                ]
              },
              "location": {
                "description": "If you have more than one account you can use this field to specify where it is.",
                "type": "string"
              },
              "name": {
                "description": "Give your sensor instance a name.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit"
            ]
          }
        },
        "total_member_count": {
          "description": "Specify the number of space members.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The amount of your space members.",
                "type": "number"
              },
              "location": {
                "description": "Specify the location if your hackerspace has different departments (for whatever reason). This field is for one department. Every department should have its own sensor instance.",
                "type": "string"
              },
              "name": {
                "description": "You can use this field to specify if this sensor instance counts active or inactive members.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value"
            ]
          }
        },
        "people_now_present": {
          "description": "Specify the number of people that are currently in your space. Optionally you can define a list of names.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The amount of present people.",
                "type": "number"
              },
              "location": {
                "description": "If you use multiple sensor instances for different rooms, use this field to indicate the location.",
                "type": "string"
              },
              "name": {
                "description": "Give this sensor a name if necessary at all. Use the location field for the rooms. This field is not intended to be used for names of hackerspace members. Use the field 'names' instead.",
                "type": "string"
              },
              "names": {
                "description": "List of hackerspace members that are currently occupying the space.",
                "type": "array",
                "items": {
                  "type": "string"
                },
                "minItems": 1
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value"
            ]
          }
        }
      }
    },
    "feeds": {
      "description": "Feeds where users can get updates of your space",
      "type": "object",
      "properties": {
        "blog": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        },
        "wiki": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        },
        "calendar": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        },
        "flickr": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        }
      }
    },
    "cache": {
      "description": "Specifies options about caching of your SpaceAPI endpoint. Use this if you want to avoid hundreds/thousands of application instances crawling your status.",
      "type": "object",
      "properties": {
        "schedule": {
          "description": "Cache update cycle. This field must match the basic regular expression <code>^[mhd]\\.[0-9]{2}$</code>, where the first field specifies a unit of time (<code>m</code> for 1 minute, <code>h</code> for 1 hour, <code>d</code> for 1 day), and the second field specifies how many of this unit should be skipped between updates. For example, <samp>m.10</samp> means one updates every 10 minutes, <samp>h.03</samp> means one update every 3 hours, and <samp>d.01</samp> means one update every day.",
          "type": "string",
          "enum": [
            "m.02",
            "m.05",
            "m.10",
            "m.15",
            "m.30",
            "h.01",
            "h.02",
            "h.04",
            "h.08",
            "h.12",
            "d.01"
          ]
        }
      },
      "required": [
        "schedule"
      ]
    },
    "projects": {
      "description": "Your project sites (links to GitHub, wikis or wherever your projects are hosted)",
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "radio_show": {
      "description": "A list of radio shows that your hackerspace might broadcast.",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": {
            "description": "The name of the radio show.",
            "type": "string"
          },
          "url": {
            "description": "The stream URL which must end in a filename or a semicolon such as <br><ul><li>http://signal.hackerspaces.org:8090/signal.mp3</li><li>http://85.214.64.213:8060/;</ul>",
            "type": "string"
          },
          "type": {
            "description": "The stream encoder.",
            "type": "string",
            "enum": [
              "mp3",
              "ogg"
            ]
          },
          "start": {
            "description": "Specify the start time by using the <a href=\"http://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\">ISO 8601</a> standard. This encodes the time as follows: <br><br><ul><li>Combined date and time in UTC: 2013-06-10T10:00Z</li><li>Combined date and time in localtime with the timezone offset: 2013-06-10T12:00+02:00</li><li>Combined date and time in localtime with the timezone offset: 2013-06-10T07:00-03:00</li></ul> The examples refer all to the same time.",
            "type": "string"
          },
          "end": {
            "description": "Specify the end time by using the <a href=\"http://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\">ISO 8601</a> standard. This encodes the time as follows: <br><br><ul><li>Combined date and time in UTC: 2013-06-10T10:00Z</li><li>Combined date and time in localtime with the timezone offset: 2013-06-10T12:00+02:00</li><li>Combined date and time in localtime with the timezone offset: 2013-06-10T07:00-03:00</li></ul> The examples refer all to the same time.",
            "type": "string"
          }
        },
        "required": [
          "name",
          "url",
          "type",
          "start",
          "end"
        ]
      }
    }
  },
  "required": [
    "api",
    "space",
    "logo",
    "url",
    "location",
    "state",
    "contact",
    "issue_report_channels"
  ]
}
`,
	"14": `{
  "$id": "https://schema.spaceapi.io/14-draft.json",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "description": "SpaceAPI 0.14",
  "type": "object",
  "properties": {
    "api": {
      "description": "The version of SpaceAPI your endpoint uses",
      "type": "string",
      "enum": [
        "0.14"
      ]
    },
    "space": {
      "description": "The name of your space",
      "type": "string"
    },
    "logo": {
      "description": "URL to your space logo",
      "type": "string"
    },
    "url": {
      "description": "URL to your space website",
      "type": "string"
    },
    "location": {
      "description": "Position data such as a postal address or geographic coordinates",
      "type": "object",
      "properties": {
        "address": {
          "description": "The postal address of your space (street, block, housenumber, zip code, city, whatever you usually need in your country, and the country itself).<br>Examples: <ul><li>Netzladen e.V., Breite Stra\u00dfe 74, 53111 Bonn, Germany</li></ul>",
          "type": "string"
        },
        "lat": {
          "description": "Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.",
          "type": "number"
        },
        "lon": {
          "description": "Longitude of your space location, in degree with decimal places. Use positive values for locations east of Greenwich, and negative values for locations west of Greenwich.",
          "type": "number"
        },
        "timezone": {
          "description": "The timezone the space is located in. It should be formatted according to the <a target=\"_blank\" href=\"https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\">TZ database location names</a>.",
          "type": "string"
        }
      },
      "required": [
        "lat",
        "lon"
      ]
    },
    "spacefed": {
      "description": "A flag indicating if the hackerspace uses SpaceFED, a federated login scheme so that visiting hackers can use the space WiFi with their home space credentials.",
      "type": "object",
      "properties": {
        "spacenet": {
          "description": "See the <a target=\"_blank\" href=\"https://spacefed.net/wiki/index.php/Category:Howto/Spacenet\">wiki</a>.",
          "type": "boolean"
        },
        "spacesaml": {
          "description": "See the <a target=\"_blank\" href=\"https://spacefed.net/wiki/index.php/Category:Howto/Spacesaml\">wiki</a>.",
          "type": "boolean"
        },
        "spacephone": {
          "description": "See the <a target=\"_blank\" href=\"https://spacefed.net/wiki/index.php/Category:Howto/Spacephone\">wiki</a>.",
          "type": "boolean"
        }
      },
      "required": [
        "spacenet",
        "spacesaml",
        "spacephone"
      ]
    },
    "cam": {
      "description": "URL(s) of webcams in your space",
      "type": "array",
      "items": {
        "type": "string"
      },
      "minItems": 1
    },
    "stream": {
      "description": "A mapping of stream types to stream URLs.If you use other stream types make a <a href=\"add-your-space\" target=\"_blank\">change request</a> or prefix yours with <samp>ext_</samp> .",
      "type": "object",
      "properties": {
        "mp4": {
          "description": "Your mpg stream URL. Example: <samp>{\"mp4\": \"http//example.org/stream.mpg\"}</samp>",
          "type": "string"
        },
        "mjpeg": {
          "description": "Your mjpeg stream URL. Example: <samp>{\"mjpeg\": \"http://example.org/stream.mjpeg\"}</samp>",
          "type": "string"
        },
        "ustream": {
          "description": "Your ustream stream URL. Example: <samp>{\"ustream\": \"http://www.ustream.tv/channel/hackspsps\"}</samp>",
          "type": "string"
        }
      }
    },
    "state": {
      "description": "A collection of status-related data: actual open/closed status, icons, last change timestamp etc.",
      "type": "object",
      "properties": {
        "open": {
          "description": "A flag which indicates whether the space is currently open or closed. The state 'undefined' can be achieved by omitting this field. A missing 'open' property carries the semantics of a temporary unavailability of the state, whereas the absence of the 'state' property itself means the state is generally not implemented for this space.",
          "type": "boolean"
        },
        "lastchange": {
          "description": "The Unix timestamp when the space status changed most recently",
          "type": "number"
        },
        "trigger_person": {
          "description": "The person who lastly changed the state e.g. opened or closed the space.",
          "type": "string"
        },
        "message": {
          "description": "An additional free-form string, could be something like <samp>'open for public'</samp>, <samp>'members only'</samp> or whatever you want it to be",
          "type": "string"
        },
        "icon": {
          "description": "Icons that show the status graphically",
          "type": "object",
          "properties": {
            "open": {
              "description": "The URL to your customized space logo showing an open space",
              "type": "string"
            },
            "closed": {
              "description": "The URL to your customized space logo showing a closed space",
              "type": "string"
            }
          },
          "required": [
            "open",
            "closed"
          ]
        }
      }
    },
    "events": {
      "description": "Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'",
      "type": "array",
      "items": {
        "required": [
          "name",
          "type",
          "timestamp"
        ],
        "type": "object",
        "properties": {
          "name": {
            "description": "Name or other identity of the subject (e.g. <samp>J. Random Hacker</samp>, <samp>fridge</samp>, <samp>3D printer</samp>, \u2026)",
            "type": "string"
          },
          "type": {
            "description": "Action (e.g. <samp>check-in</samp>, <samp>check-out</samp>, <samp>finish-print</samp>, \u2026). Define your own actions and use them consistently, canonical actions are not (yet) specified",
            "type": "string"
          },
          "timestamp": {
            "description": "Unix timestamp when the event occurred",
            "type": "number"
          },
          "extra": {
            "description": "A custom text field to give more information about the event",
            "type": "string"
          }
        }
      }
    },
    "contact": {
      "description": "Contact information about your space. You must define at least one which is in the list of allowed values of the issue_report_channels field.",
      "type": "object",
      "properties": {
        "phone": {
          "description": "Phone number, including country code with a leading plus sign. Example: <samp>+1 800 555 4567</samp>",
          "type": "string"
        },
        "sip": {
          "description": "URI for Voice-over-IP via SIP. Example: <samp>sip:yourspace@sip.example.org</samp>",
          "type": "string"
        },
        "keymasters": {
          "description": "Persons who carry a key and are able to open the space upon request. One of the fields irc_nick, phone, email or twitter must be specified.",
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "object",
            "properties": {
              "name": {
                "description": "Real name",
                "type": "string"
              },
              "irc_nick": {
                "description": "Contact the person with this nickname directly in irc if available. The irc channel to be used is defined in the contact/irc field.",
                "type": "string"
              },
              "phone": {
                "description": "Example: <samp>['+1 800 555 4567','+1 800 555 4544']</samp>",
                "type": "string"
              },
              "email": {
                "description": "Email address which can be base64 encoded.",
                "type": "string"
              },
              "twitter": {
                "description": "Twitter username with leading <samp>@</samp>.",
                "type": "string"
              },
              "xmpp": {
                  "description": "XMPP (Jabber) ID",
                  "type": "string"
              },
              "mastodon": {
                  "description": "Mastodon username. Example: @ordnung@chaos.social",
                  "type": "string"
              },
              "matrix": {
                "description": "Matrix username. Example: <samp>@user:example.org</samp>",
                "type": "string"
              }
            }
          }
        },
        "irc": {
          "description": "URL of the IRC channel, in the form <samp>irc://example.org/#channelname</samp>",
          "type": "string"
        },
        "twitter": {
          "description": "Twitter handle, with leading @",
          "type": "string"
        },
        "mastodon": {
          "description": "Mastodon username: Example: @ordnung@chaos.social.",
          "type": "string"
        },
        "facebook": {
          "description": "Facebook account URL.",
          "type": "string"
        },
        "identica": {
          "description": "Identi.ca or StatusNet account, in the form <samp>yourspace@example.org</samp>",
          "type": "string"
        },
        "foursquare": {
          "description": "Foursquare ID, in the form <samp>4d8a9114d85f3704eab301dc</samp>.",
          "type": "string"
        },
        "email": {
          "description": "E-mail address for contacting your space. If this is a mailing list consider to use the contact/ml field.",
          "type": "string"
        },
        "ml": {
          "description": "The e-mail address of your mailing list. If you use Google Groups then the e-mail looks like <samp>your-group@googlegroups.com</samp>.",
          "type": "string"
        },
        "xmpp": {
          "description": "A public Jabber/XMPP multi-user chatroom in the form <samp>chatroom@conference.example.net</samp>",
          "type": "string"
        },
        "issue_mail": {
          "description": "A separate email address for issue reports. This value can be Base64-encoded.",
          "type": "string"
        },
        "gopher": {
          "description": "A URL to find information about the Space in the Gopherspace. Example: gopher://gopher.binary-kitchen.de",
          "type": "string"
        },
        "matrix": {
          "description": "Matrix channel/community for the Hackerspace. Example: <samp>#spaceroom:example.org</samp> or <samp>+spacecommunity:example.org</samp>",
          "type": "string"
        }
      }
    },
    "sensors": {
      "description": "Data of various sensors in your space (e.g. temperature, humidity, amount of Club-Mate left, \u2026). The only canonical property is the <em>temp</em> property, additional sensor types may be defined by you. In this case, you are requested to share your definition for inclusion in this specification.",
      "type": "object",
      "properties": {
        "temperature": {
          "description": "Temperature sensor. To convert from one unit of temperature to another consider <a href=\"http://en.wikipedia.org/wiki/Temperature_conversion_formulas\" target=\"_blank\">Wikipedia</a>.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value.",
                "type": "string",
                "enum": [
                  "\u00b0C",
                  "\u00b0F",
                  "K",
                  "\u00b0De",
                  "\u00b0N",
                  "\u00b0R",
                  "\u00b0R\u00e9",
                  "\u00b0R\u00f8"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "door_locked": {
          "description": "Sensor type to indicate if a certain door is locked.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "boolean"
              },
              "location": {
                "description": "The location of your sensor such as <samp>front door</samp>, <samp>chill room</samp> or <samp>lab</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "location"
            ]
          }
        },
        "barometer": {
          "description": "Barometer sensor",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                "type": "string",
                "enum": [
                  "hPa"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "radiation": {
          "description": "Compound radiation sensor. Check this <a rel=\"nofollow\" href=\"https://sites.google.com/site/diygeigercounter/gm-tubes-supported\" target=\"_blank\">resource</a>.",
          "type": "object",
          "properties": {
            "alpha": {
              "description": "An alpha sensor",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            },
            "beta": {
              "description": "A beta sensor",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            },
            "gamma": {
              "description": "A gamma sensor",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            },
            "beta_gamma": {
              "description": "A sensor which cannot filter beta and gamma radiation separately.",
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "value": {
                    "description": "Observed counts per minute (ocpm) or actual radiation value. If the value are the observed counts then the dead_time and conversion_factor fields must be defined as well. CPM formula: <div>cpm = ocpm ( 1 + 1 / (1 - ocpm x dead_time) )</div> Conversion formula: <div>\u00b5Sv/h = cpm x conversion_factor</div>",
                    "type": "number"
                  },
                  "unit": {
                    "description": "Choose the appropriate unit for your radiation sensor instance.",
                    "type": "string",
                    "enum": [
                      "cpm",
                      "r/h",
                      "\u00b5Sv/h",
                      "mSv/a",
                      "\u00b5Sv/a"
                    ]
                  },
                  "dead_time": {
                    "description": "The dead time in \u00b5s. See the description of the value field to see how to use the dead time.",
                    "type": "number"
                  },
                  "conversion_factor": {
                    "description": "The conversion from the <em>cpm</em> unit to another unit hardly depends on your tube type. See the description of the value field to see how to use the conversion factor. <strong>Note:</strong> only trust your manufacturer if it comes to the actual factor value. The internet seems <a rel=\"nofollow\" href=\"http://sapporohibaku.wordpress.com/2011/10/15/conversion-factor/\" target=\"_blank\">full of wrong copy & pastes</a>, don't even trust your neighbour hackerspace. If in doubt ask the tube manufacturer.",
                    "type": "number"
                  },
                  "location": {
                    "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                    "type": "string"
                  },
                  "name": {
                    "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                    "type": "string"
                  },
                  "description": {
                    "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                    "type": "string"
                  }
                },
                "required": [
                  "value",
                  "unit"
                ]
              }
            }
          }
        },
        "humidity": {
          "description": "Humidity sensor",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                "type": "string",
                "enum": [
                  "%"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "beverage_supply": {
          "description": "How much Mate and beer is in your fridge?",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit, either <samp>btl</samp> for bottles or <samp>crt</samp> for crates.",
                "type": "string",
                "enum": [
                  "btl",
                  "crt"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Room 1</samp> or <samp>Room 2</samp> or <samp>Room 3</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit"
            ]
          }
        },
        "power_consumption": {
          "description": "The power consumption of a specific device or of your whole space.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The sensor value",
                "type": "number"
              },
              "unit": {
                "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                "type": "string",
                "enum": [
                  "mW",
                  "W",
                  "VA"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit",
              "location"
            ]
          }
        },
        "wind": {
          "description": "Your wind sensor.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "properties": {
                "description": "",
                "type": "object",
                "properties": {
                  "speed": {
                    "description": "",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "m/s",
                          "km/h",
                          "kn"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  },
                  "gust": {
                    "description": "",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "m/s",
                          "km/h",
                          "kn"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  },
                  "direction": {
                    "description": "The wind direction in degrees.",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "\u00b0"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  },
                  "elevation": {
                    "description": "Height above mean sea level.",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The sensor value",
                        "type": "number"
                      },
                      "unit": {
                        "description": "The unit of the sensor value. You should always define the unit though if the sensor is a flag of a boolean type then you can of course omit it.",
                        "type": "string",
                        "enum": [
                          "m"
                        ]
                      }
                    },
                    "required": [
                      "value",
                      "unit"
                    ]
                  }
                },
                "required": [
                  "speed",
                  "gust",
                  "direction",
                  "elevation"
                ]
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "properties",
              "location"
            ]
          }
        },
        "network_connections": {
          "description": "This sensor type is to specify the currently active  ethernet or wireless network devices. You can create different instances for each network type.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "type": {
                "description": "This field is optional but you can use it to the network type such as <samp>wifi</samp> or <samp>cable</samp>. You can even expose the number of <a href=\"https://spacefed.net/wiki/index.php/Spacenet\" target=\"_blank\">spacenet</a>-authenticated connections.",
                "type": "string",
                "enum": [
                  "wifi",
                  "cable",
                  "spacenet"
                ]
              },
              "value": {
                "description": "The amount of network connections.",
                "type": "number"
              },
              "machines": {
                "description": "The machines that are currently connected with the network.",
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "name": {
                      "description": "The machine name.",
                      "type": "string"
                    },
                    "mac": {
                      "description": "The machine's MAC address of the format <samp>D3:3A:DB:EE:FF:00</samp>.",
                      "type": "string"
                    }
                  },
                  "required": [
                    "mac"
                  ]
                }
              },
              "location": {
                "description": "The location of your sensor such as <samp>Outside</samp>, <samp>Inside</samp>, <samp>Ceiling</samp>, <samp>Roof</samp> or <samp>Room 1</samp>.",
                "type": "string"
              },
              "name": {
                "description": "This field is an additional field to give your sensor a name. This can be useful if you have multiple sensors in the same location.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value"
            ]
          }
        },
        "account_balance": {
          "description": "How rich is your hackerspace?",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "How much?",
                "type": "number"
              },
              "unit": {
                "description": "What's the currency? It should be formatted according to <a href=\"https://en.wikipedia.org/wiki/ISO_4217\" target=\"_blank\">ISO 4217</a> short-code format.",
                "type": "string"
              },
              "location": {
                "description": "If you have more than one account you can use this field to specify where it is.",
                "type": "string"
              },
              "name": {
                "description": "Give your sensor instance a name.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value",
              "unit"
            ]
          }
        },
        "total_member_count": {
          "description": "Specify the number of space members.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The amount of your space members.",
                "type": "number"
              },
              "location": {
                "description": "Specify the location if your hackerspace has different departments (for whatever reason). This field is for one department. Every department should have its own sensor instance.",
                "type": "string"
              },
              "name": {
                "description": "You can use this field to specify if this sensor instance counts active or inactive members.",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value"
            ]
          }
        },
        "people_now_present": {
          "description": "Specify the number of people that are currently in your space. Optionally you can define a list of names.",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "value": {
                "description": "The amount of present people.",
                "type": "number"
              },
              "location": {
                "description": "If you use multiple sensor instances for different rooms, use this field to indicate the location.",
                "type": "string"
              },
              "name": {
                "description": "Give this sensor a name if necessary at all. Use the location field for the rooms. This field is not intended to be used for names of hackerspace members. Use the field 'names' instead.",
                "type": "string"
              },
              "names": {
                "description": "List of hackerspace members that are currently occupying the space.",
                "type": "array",
                "items": {
                  "type": "string"
                },
                "minItems": 1
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance.",
                "type": "string"
              }
            },
            "required": [
              "value"
            ]
          }
        },
        "network_traffic": {
          "description": "The current network traffic, in bits/second or packets/second (or both)",
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "properties": {
                "type": "object",
                "properties": {
                  "bits_per_second": {
                    "description": "",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The measurement value, in bits/second",
                        "type": "number",
                        "minimum": 0
                      },
                      "maximum": {
                        "description": "The maximum available throughput in bits/second, e.g. as sold by your ISP",
                        "type": "number",
                        "minimum": 0
                      }
                    },
                    "required": [
                      "value"
                    ]
                  },
                  "packets_per_second": {
                    "description": "",
                    "type": "object",
                    "properties": {
                      "value": {
                        "description": "The measurement value, in packets/second",
                        "type": "number",
                        "minimum": 0
                      }
                    },
                    "required": [
                      "value"
                    ]
                  }
                }
              },
              "name": {
                "description": "Name of the measurement, e.g. to distinguish between upstream and downstream traffic",
                "type": "string"
              },
              "location": {
                "description": "Location the measurement relates to, e.g. <samp>WiFi</samp> or <samp>Uplink</samp>",
                "type": "string"
              },
              "description": {
                "description": "An extra field that you can use to attach some additional information to this sensor instance",
                "type": "string"
              }
            },
            "required": [
              "properties"
            ]
          },
          "minItems": 1
        }
      }
    },
    "feeds": {
      "description": "Feeds where users can get updates of your space",
      "type": "object",
      "properties": {
        "blog": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        },
        "wiki": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        },
        "calendar": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        },
        "flickr": {
          "description": "",
          "type": "object",
          "properties": {
            "type": {
              "description": "Type of the feed, for example <samp>rss</samp>, <samp>atom</atom>, <samp>ical</samp>",
              "type": "string"
            },
            "url": {
              "description": "Feed URL",
              "type": "string"
            }
          },
          "required": [
            "url"
          ]
        }
      }
    },
    "projects": {
      "description": "Your project sites (links to GitHub, wikis or wherever your projects are hosted)",
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "membership_plans": {
        "description": "A list of the different membership plans your hackerspace might have. Set the value according to your billing process. For example, if your membership fee is 10€ per month, but you bill it yearly (aka. the member pays the fee once per year), set the amount to 120 an the billing_interval to yearly.",
        "type": "array",
        "items": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "The name of the membership plan. For example Student Membership, Normal Membership etc.",
                    "type": "string"
                },
                "value": {
                    "description": "How much does this plan cost?",
                    "type": "number"
                },
                "currency": {
                    "description": "What's the currency? It should be formatted according to <a href=\"https://en.wikipedia.org/wiki/ISO_4217\" target=\"_blank\">ISO 4217</a> short-code format.",
                    "type": "string"
                },
                "billing_interval": {
                    "description": "How often is the membership billed? If you select other, please specify in the description what your billing interval is.",
                    "type": "string",
                    "enum": [
                        "yearly",
                        "monthly",
                        "weekly",
                        "daily",
                        "hourly",
                        "other"
                    ]
                },
                "description": {
                    "description": "A free form string.",
                    "type": "string"
                }
            },
            "required": [
                "name",
                "value",
                "currency",
                "billing_interval"
            ]
        }
    }
  },
  "required": [
    "api",
    "space",
    "logo",
    "url",
    "location",
    "contact"
  ]
}
`,
	"8": `{
  "$id": "https://schema.spaceapi.io/8.json",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "description": "SpaceAPI 0.8",
  "type": "object",
  "properties": {
    "api": {
      "description": "The version of SpaceAPI your endpoint uses",
      "type": "string",
      "enum": [
        "0.8"
      ]
    },
    "space": {
      "description": "The name of your space",
      "type": "string"
    },
    "logo": {
      "description": "The space logo",
      "type": "string"
    },
    "url": {
      "type": "string"
    },
    "address": {
      "description": "The main website",
      "type": "string"
    },
    "phone": {
      "description": "Phone number, including country code with a leading plus sign. Example: <samp>+1 800 555 4567</samp>",
      "type": "string"
    },
    "lat": {
      "description": "Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.",
      "type": "number"
    },
    "lon": {
      "description": "Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.",
      "type": "number"
    },
    "cam": {
      "description": "URL(s) of webcams in your space",
      "type": "array",
      "items": {
        "type": "string"
      },
      "minItems": 1
    },
    "stream": {
      "description": "A mapping of stream types to stream URLs. Example: <samp>{'mp4':'http//example.org/stream.mpg', 'mjpeg':'http://example.org/stream.mjpeg'}</samp>",
      "type": "object"
    },
    "open": {
      "description": "A boolean which indicates if the space is currently open",
      "type": "boolean"
    },
    "status": {
      "description": "An additional free-form string, could be something like <samp>'open for public'</samp>, <samp>'members only'</samp> or whatever you want it to be",
      "type": "string"
    },
    "lastchange": {
      "description": "The Unix timestamp when the space status changed most recently",
      "type": "number"
    },
    "events": {
      "description": "Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'",
      "type": "array",
      "items": {
        "required": [
          "name",
          "type",
          "t"
        ],
        "type": "object",
        "properties": {
          "name": {
            "description": "Name or other identity of the subject (e.g. <samp>J. Random Hacker</samp>, <samp>fridge</samp>, <samp>3D printer</samp>, \u2026)",
            "type": "string"
          },
          "type": {
            "description": "Action (e.g. <samp>check-in</samp>, <samp>check-out</samp>, <samp>finish-print</samp>, \u2026). Define your own actions and use them consistently, canonical actions are not (yet) specified",
            "type": "string"
          },
          "t": {
            "description": "Unix timestamp when the event occurred",
            "type": "number"
          },
          "extra": {
            "description": "A custom text field to give more information about the event",
            "type": "string"
          }
        }
      }
    }
  },
  "required": [
    "api",
    "space",
    "logo",
    "url",
    "open"
  ]
}
`,
	"9": `{
  "$id": "https://schema.spaceapi.io/9.json",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "description": "SpaceAPI 0.9",
  "type": "object",
  "properties": {
    "api": {
      "description": "The version of SpaceAPI your endpoint uses",
      "type": "string",
      "enum": [
        "0.9"
      ]
    },
    "space": {
      "description": "The name of your space",
      "type": "string"
    },
    "logo": {
      "description": "The space logo",
      "type": "string"
    },
    "url": {
      "description": "The main website",
      "type": "string"
    },
    "address": {
      "description": "The postal address of your space (street, block, housenumber, zip code, city, whatever you usually need in your country, and the country itself)",
      "type": "string"
    },
    "lat": {
      "description": "Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.",
      "type": "number"
    },
    "lon": {
      "description": "Longitude of your space location, in degree with decimal places. Use positive values for locations west of Greenwich, and negative values for locations east of Greenwich.",
      "type": "number"
    },
    "cam": {
      "description": "URL(s) of webcams in your space",
      "type": "array",
      "items": {
        "type": "string"
      },
      "minItems": 1
    },
    "stream": {
      "description": "A mapping of stream types to stream URLs. Example: <samp>{'mp4':'http//example.org/stream.mpg', 'mjpeg':'http://example.org/stream.mjpeg'}</samp>",
      "type": "object"
    },
    "open": {
      "description": "A boolean which indicates if the space is currently open",
      "type": "boolean"
    },
    "status": {
      "description": "An additional free-form string, could be something like <samp>'open for public'</samp>, <samp>'members only'</samp> or whatever you want it to be",
      "type": "string"
    },
    "lastchange": {
      "description": "The Unix timestamp when the space status changed most recently",
      "type": "number"
    },
    "events": {
      "description": "Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'",
      "type": "array",
      "items": {
        "required": [
          "name",
          "type",
          "t"
        ],
        "type": "object",
        "properties": {
          "name": {
            "description": "Name or other identity of the subject (e.g. <samp>J. Random Hacker</samp>, <samp>fridge</samp>, <samp>3D printer</samp>, \u2026)",
            "type": "string"
          },
          "type": {
            "description": "Action (e.g. <samp>check-in</samp>, <samp>check-out</samp>, <samp>finish-print</samp>, \u2026). Define your own actions and use them consistently, canonical actions are not (yet) specified",
            "type": "string"
          },
          "t": {
            "description": "Unix timestamp when the event occurred",
            "type": "number"
          },
          "extra": {
            "description": "A custom text field to give more information about the event",
            "type": "string"
          }
        }
      }
    },
    "contact": {
      "description": "Contact information about your space",
      "type": "object",
      "properties": {
        "phone": {
          "description": "Phone number, including country code with a leading plus sign. Example: <samp>+1 800 555 4567</samp>",
          "type": "string"
        },
        "sip": {
          "description": "URI for Voice-over-IP via SIP. Example: <samp>sip:yourspace@sip.example.org</samp>",
          "type": "string"
        },
        "keymaster": {
          "description": "Phone numbers of people who carry a key and are able to open the space upon request. Example: <samp>['+1 800 555 4567','+1 800 555 4544']</samp>",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "irc": {
          "description": "URL of the IRC channel, in the form <samp>irc://example.org/#channelname</samp>",
          "type": "string"
        },
        "twitter": {
          "description": "Twitter handle, with leading @",
          "type": "string"
        },
        "email": {
          "description": "E-mail address for contacting your space",
          "type": "string"
        },
        "ml": {
          "description": "The e-mail address of your mailing list. If you use Google Groups then the e-mail looks like <samp>your-group@googlegroups.com</samp>.",
          "type": "string"
        },
        "jabber": {
          "description": "A public Jabber/XMPP multi-user chatroom in the form <samp>chatroom@conference.example.net</samp>",
          "type": "string"
        }
      }
    }
  },
  "required": [
    "api",
    "space",
    "logo",
    "url",
    "open"
  ]
}
`,
}
