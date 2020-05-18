PUT _template/examplelog/


{
  "index_patterns": [
    "examplelog*"
  ],
  "settings": {
    "index": {
      "number_of_shards": "1"
    }
  },
  "mappings": {
    "_doc": {
      "properties": {
        "logtext": {
          "type": "text"
        },
        "hostip": {
          "type": "text"
        },
        "image": {
          "type": "text"
        },
        "description": {
          "type": "text"
        },
        "status": {
          "type": "text"
        },
        "state": {
          "type": "text"
        },
        "time": {
          "type": "date",
          "format": "yyyy-MM-dd HH:mm:ss"
        },
        "source": {
          "type": "text"
        },
        "hostname": {
          "type": "text"
        }
      }
    }
  }
}