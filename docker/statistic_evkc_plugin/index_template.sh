#!/bin/bash

# Define the index template JSON content

curl -X PUT "http://localhost:9200/_index_template/notifications" -H 'Content-Type: application/json' -d '{
  "index_patterns": ["notifications*"],
  "template": {
    "mappings": {
      "properties": {
        "headers": {
          "type": "object"
        },
        "message": {
          "properties": {
            "error": {
              "type": "text"
            },
            "message": {
              "type": "text"
            },
            "platform": {
              "type": "keyword"
            },
            "token": {
              "type": "keyword"
            },
            "type": {
              "type": "keyword"
            }
          }
        },
        "message_key": {
          "type": "keyword"
        },
        "offset": {
          "type": "long"
        },
        "partition": {
          "type": "integer"
        },
        "source_type": {
          "type": "keyword"
        },
        "timestamp": {
          "type": "date",
          "format": "strict_date_optional_time||epoch_millis"
        },
        "topic": {
          "type": "keyword"
        }
      }
    }
  }
}'
