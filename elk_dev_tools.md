
GET general_mobin/_search
{

 "size": 10,
      "query": {
        "bool": {
          "should": [
            {
              "match": {
                
                "description": {"query": "Arabia persian","operator": "and","fuzziness": 2}
              }
            },
            
           
            
                            {
              "match": {
                
                "fa_description": {"query": "Arabia persian","operator": "and","fuzziness": 2}
              }
            }
            
            ,
            
                        
                            {
              "match": {
                
                "name": {"query": "Arabia persian","operator": "and","fuzziness": 2}
              }
            }
            
            ,
            
           
           
           
            
                            {
              "match": {
                
                "fa_name": {"query": "Arabia","operator": "and","fuzziness": 2}
              }
            }
            
            
            ]
        }
      }
}


GET /nvds/_search
{
  "size": 500,
  "query": {
    "bool": {
      "must": [
        {
          "terms": {
            "nvd.impact.baseMetricV3.cvssV3.attackVector.keyword": ["ADJACENT_NETWORK","LOCAL"]
          }
          
          
        }
        
      ]
    }
  }
}


GET /nvds/_search
{
  "size": 5,
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "nvd.impact.baseMetricV3.cvssV3.attackVector.keyword": "ADJACENT_NETWORK"
          }
        }
        
        ,
        
        
        
        
        
        {
          "bool": {
            "should": [
              {
                "match": {
                  "nvd.impact.baseMetricV3.cvssV3.attackVector.keyword": "ADJACENT_NETWORK"
                }
              }
            ]
          }
        }
      ]
    }
  }
}




















DELETE test16

PUT test16

GET test111/_search



GET /test16/_doc/1702123835777335



GET feeds/_doc/1704001470141345


GET feeds/_search
{
  "query": {
    "match_all": {}  
  },
  "sort": [
    {
      "@timestamp": {
        "order": "desc"  
      }
    }
  ],
  "size": 100
}


GET feeds/_search
{

  "sort": [
    {
      "_id": {
        "order": "asc"  
      }
    }
  ],
  "size": 100
}





GET /nvds/_search
{"size": 20,
  "sort": [
    {
      "nvd.impact.baseMetricV3.cvssV3.baseScore": {
        "order": "desc"
      }
    }
  ], 
  "query": {
    "terms": {
      "nvd.impact.baseMetricV3.cvssV3.attackVector.keyword": [
      "ADJACENT_NETWORK", "LOCAL"
      ]
    }  }
  
}

GET /feeds/_search
{"size": 10,

  "query": {
    "match": {
      "category.keyword": "tech"
    }
  },
  "aggs": {
    "sentiment_negative": {
    
        "terms": {
          "field": "nerlocation.name.keyword"
        }
    
     
      }
      
        , 
         "fff": {
           "avg": {
             "field": "docsentiment.positive"
           }
         }
       
      
    
  }
}









PUT mobin
{
  "mappings": {
    "dynamic": "false",
    "properties": {
      "description": {
        "type": "text"
      },
      "fa_description": {
        "type": "text"
      },
      
      "fa_name": {
        "type": "text"
      },
      
      "name": {
        "type": "text"
      },
      
      "status": {
        "type": "keyword"
      },
      
      "target_location": {
        "type": "keyword"
      },
      
      "tlp": {
        "type": "keyword"
      },
      
      
      "labels": {
        "type": "keyword"
      },
            
            
            
      "published_by": {
        "type": "keyword"
      },
        
      "published": {
        "type": "date"
      }
 
    }
  }
}






GET general_mobin/_search?_source
{"size": 0,
  "aggs": {
    "published_by": {
      "terms": {
        "field": "target_location.keyword"
       
      },
      
      "aggs": {
        "agg2": {
          "terms": {
            "field": "fa_target_location.keyword"
          }
        }
        
      }
      
    
      
      
    }
   
  }
}


GET feeds/_search?_source
{"size": 0,
  "aggs": {
    "published_by": {
      "terms": {
        "field": "ner_stats.name.keyword"
       
      }
    
      
      
    }
   
  }
}

POST test11/_delete_by_query
{
  "query": {
          "exists": {
            "field": "hash"
          }
  }
}

GET feeds/_mapping


GET /feeds/_search
{
  "query": {
    "range": {
      "docsentiment.positive": {
        "gte": 0.9,
        "lte": 1
      }
    }
  }
}



GET test/_search
		{
		  "query": {
		  "term": {
			"idd.keyword": "test5"
		  }
		}
		
		
		}
		
		
		
GET test/_search
{
 "query": {
    "term": {
          "idd.keyword": "test2"
    }
  }
}


PUT test7/
{
  "mappings": {
     "dynamic": "strict",
    "properties": {
      "description":    { "type": "text" },  
      "hash":  { "type": "keyword"  }, 
         "news_agancy":  { "type": "keyword"  }, 
      
      "link":  { "type": "keyword",
          "index": false  },      
        "create_time" : {
          "type" : "date"
        },
         "text":   { "type": "text"  } ,
      "title":   { "type": "text"  }     
    }
  }
}

POST test/_update_by_query
{
		"script": {
		  "source": "ctx._source.name = params.name;ctx._source.com = params.com;ctx._source.idd = params.idd",
		  "lang":   "painless",
		  "params": {
			"name": "qqqqqqqqq",
			"idd": "test2",
			
			"com": 123456
		  }
		},
		
		"query": {
		  "term": {
			"idd.keyword": "test2"
		  }
		}
}


POST feeds/_update_by_query
{
    "script": {
        "source": "ctx._source.docsentiment = params.docsentiment",
        "lang": "painless",
        "params": {
            "docsentiment": {
                "negative": 0.08,
                "neutral": 0.75,
                "positive": 0.17
            }
        }
    },
    "query": {
        "term": {
            "common_key.keyword": "b9e0c6de-9613-404d-9377-bdc43175d3b1"
        }
    }
}


PUT test14/_mapping
{
    "properties": {
        "authors" : {
          "type" : "keyword"
        }
    }
  
}


POST _bulk
{ "update": {"_index":"test","_id":"12345679"}}
{ "doc_as_upsert":true, "doc" :  {"Category":"Category"}}







POST feeds/_update_by_query
{"script": {
      "source":"ctx._source.summary_fa = params.summary_fa;ctx._source.doc_sentiment = params.doc_sentiment",
      "lang": "painless",
      "params": {"summary_fa":"not ffff","doc_sentiment":{"negative":0.26,"neutral":0.7,"positive":0.05}}
    },
    "query": {
      "term": {
            "common_key.keyword": "51c04a79-83a1-49d4-88eb-3cfc16cb0287"
      }
    }
    }






PUT test15/
{
    "mappings": {
        "dynamic": "false",
        "properties": {
                  "docsentiment" : {
          "properties" : {
            "negative" : {
              "type" : "float"
            },
            "neutral" : {
              "type" : "float"
            },
            "positive" : {
              "type" : "float"
            }
          }
        },
        
        
               "nerlocation" : {
          "properties" : {
            "coordinate" : {
              "properties" : {
                "lat" : {
                  "type" : "float"
                },
                "long" : {
                  "type" : "float"
                }
              }
            },
            "name" : {
              "type" : "keyword"
             }
          }
        },
      
        
        
            "tag": {
                "type": "keyword"
            },
        
        
               "category" : {
          "type" : "keyword"
        },
  
        
                "most_mentioned_ents" : {
              "properties" : {
                "count" : {
                  "type" : "byte"
                },
                "name" : {
                  "type" : "keyword"
                 }
              }
            },
  
                "ner_done" : {
          "type" : "boolean"
        },
        
        
            "summary": {
                "type": "text"
            },
            
                    
            "summary_fa": {
                "type": "text"
            },
            
           "title_fa": {
                "type": "text"
            },
            
            
            "common_key": {
                "type": "keyword"
            },
            "authors": {
                "type": "keyword"
            },
            "site": {
                "type": "keyword"
            },
            "url": {
                "type": "keyword",
                "index": false
            },
            "create_date": {
                "type": "date"
            },
            "main_content": {
                "type": "text"
            },
            "title": {
                "type": "text"
            }
        }
    }
}