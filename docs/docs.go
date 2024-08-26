// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.0",
  "info": {
    "title": "Like Movies",
    "version": "1.0.0"
  },
  "servers": [
    {
      "url": "http://localhost:5000"
    }
  ],
  "tags": [
    {
      "name": "Users"
    },
    {
      "name": "Movies"
    }
  ],
  "paths": {
    "/api/v1/user/create": {
      "post": {
        "tags": [
          "Users"
        ],
        "summary": "register user",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "example": {
                  "name": "User Test",
                  "password": "Teste@123",
                  "username": "user_test"
                }
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "Created",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": true,
                  "data": {
                    "message": "Usuário cadastrado com sucesso.",
                    "response": {
                      "id": 1,
                      "name": "Bruno Lopes",
                      "username": "brunolls",
                      "created_at": "2024-08-22 22:32:51",
                      "updated_at": "2024-08-22 22:32:51"
                    }
                  }
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": false,
                  "data": {
                    "message": "Erro ao cadastrar usuário.",
                    "response": {
                      "errors": {
                        "name": "Nome é um campo obrigatório."
                      }
                    }
                  }
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": false,
                  "data": {
                    "message": "Erro ao cadastrar usuário.",
                    "response": {
                      "log_id": "70fcd8c7-da8f-4337-bb42-ddbac29d97cc",
                      "event_date": "2024-08-24 22:50:21"
                    }
                  }
                }
              }
            }
          }
        },
        "security": [
          {
            "basicAuth": []
          }
        ]
      }
    },
    "/api/v1/movies/favorite": {
      "post": {
        "tags": [
          "Movies"
        ],
        "summary": "Favorite movie",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "example": {
                  "id": 1,
                  "rating": 4
                }
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "Created",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": true,
                  "data": {
                    "message": "Filme adicionado à lista com sucesso.",
                    "response": {
                      "id": 970689,
                      "comment": "",
                      "rating": 4,
                      "user_id": 1,
                      "created_at": "2024-08-22 22:55:40",
                      "updated_at": "2024-08-22 22:55:40"
                    }
                  }
                }
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": false,
                  "data": {
                    "message": "Erro ao adicionar filme na lista.",
                    "response": {
                      "errors": {
                        "id": "Id é um campo obrigatório"
                      }
                    }
                  }
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": false,
                  "data": {
                    "message": "Erro ao adicionar filme na lista.",
                    "response": {
                      "log_id": "70fcd8c7-da8f-4337-bb42-ddbac29d97cc",
                      "event_date": "2024-08-22 22:50:32"
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/api/v1/movies/list": {
      "get": {
        "tags": [
          "Movies"
        ],
        "summary": "List favorites",
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": true,
                  "data": {
                    "message": "Busca da lista realizada com sucesso.",
                    "response": [
                      {
                        "id": 970689,
                        "comment": "",
                        "rating": 4,
                        "user_id": 1,
                        "craeted_at": "2024-08-22 22:55:40",
                        "updated_at": "2024-08-22 22:55:40"
                      },
                      {
                        "id": 238,
                        "comment": "",
                        "rating": 2,
                        "user_id": 1,
                        "craeted_at": "2024-08-22 22:55:40",
                        "updated_at": "2024-08-22 22:55:40"
                      }
                    ]
                  }
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": false,
                  "data": {
                    "message": "Erro ao buscar lista de favoritos.",
                    "response": {}
                  }
                }
              }
            }
          }
        },
        "security": [
          {
            "basicAuth": []
          }
        ]
      }
    },
    "/api/v1/movies/{id}/remove": {
      "delete": {
        "tags": [
          "Movies"
        ],
        "summary": "Remove favorite",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "style": "simple",
            "explode": false,
            "schema": {
              "type": "integer"
            },
            "example": "970689"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "examples": {
                  "example-0": {
                    "summary": "success",
                    "value": {
                      "status": true,
                      "data": {
                        "message": "Filme removido da lista com sucesso."
                      }
                    }
                  },
                  "example-1": {
                    "summary": "generic error",
                    "value": {
                      "status": false,
                      "data": {
                        "message": "Erro ao remover filme da lista.",
                        "response": {
                          "log_id": "70fcd8c7-da8f-4337-bb42-ddbac29d97cc",
                          "event_date": "2024-08-22 22:50:32"
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/api/v1/movies/{id}/update": {
      "put": {
        "tags": [
          "Movies"
        ],
        "summary": "Update favorite",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "style": "simple",
            "explode": false,
            "schema": {
              "type": "integer"
            },
            "example": "970689"
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "example": {
                  "comment": "Comentário teste 2",
                  "rating": 2
                }
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": true,
                  "data": {
                    "message": "Filme atualizado com sucesso.",
                    "response": {
                      "id": 970689,
                      "comment": "",
                      "rating": 4,
                      "user_id": 1,
                      "created_at": "2024-08-22 22:55:40",
                      "updated_at": "2024-08-22 22:55:40"
                    }
                  }
                }
              }
            }
          },
          "500": {
            "description": "Internal Server Error",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "example": {
                  "status": false,
                  "data": {
                    "message": "Erro ao atualizar filme da lista.",
                    "response": {
                      "log_id": "70fcd8c7-da8f-4337-bb42-ddbac29d97cc",
                      "event_date": "2024-08-22 22:50:32"
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/api/v1/movies/search": {
      "get": {
        "tags": [
          "Movies"
        ],
        "summary": "Search movies",
        "parameters": [
          {
            "name": "name",
            "in": "query",
            "required": false,
            "style": "form",
            "explode": true,
            "schema": {
              "type": "string"
            },
            "example": "the godfa"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "examples": {
                  "example-0": {
                    "summary": "success",
                    "value": {
                      "status": true,
                      "data": {
                        "message": "Busca realizada com sucesso.",
                        "response": [
                          {
                            "id": 238,
                            "title": "O Poderoso Chefão",
                            "description": "Em 1945, Don Corleone é o chefe de uma mafiosa família italiana de Nova York. Ele costuma apadrinhar várias pessoas, realizando importantes favores para elas, em troca de favores futuros. Com a chegada das drogas, as famílias começam uma disputa pelo promissor mercado. Quando Corleone se recusa a facilitar a entrada dos narcóticos na cidade, não oferecendo ajuda política e policial, sua família começa a sofrer atentados para que mudem de posição. É nessa complicada época que Michael, um herói de guerra nunca envolvido nos negócios da família, vê a necessidade de proteger o seu pai e tudo o que ele construiu ao longo dos anos.",
                            "release_date": "1972-03-14"
                          },
                          {
                            "id": 970689,
                            "title": "The Godfather",
                            "description": "Quando Corleone, um respeitado chefe da máfia nova-iorquina,",
                            "release_date": "1981-03-24"
                          },
                          {
                            "id": 242,
                            "title": "O Poderoso Chefão: Parte III",
                            "description": "Don Michael Corleone está envelhecendo e, com a ajuda do sobrinho Vicente Mancini, busca a legitimação dos interesses da família, em Nova York e na Itália, antes de sua morte. Mas seu protegido não está só interessado em parte do império da família. Ele também deseja a filha de Michael, Mary.",
                            "release_date": "1990-12-25"
                          }
                        ]
                      }
                    }
                  },
                  "example-1": {
                    "summary": "generic error",
                    "value": {
                      "status": false,
                      "data": {
                        "message": "Erro ao realizar busca.",
                        "response": {
                          "log_id": "70fcd8c7-da8f-4337-bb42-ddbac29d97cc",
                          "event_date": "2024-08-22 22:50:32"
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/api/v1/movies/{id}/details": {
      "get": {
        "tags": [
          "Movies"
        ],
        "summary": "Details movie",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "style": "simple",
            "explode": false,
            "schema": {
              "type": "integer"
            },
            "example": "238"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "headers": {
              "Content-Type": {
                "style": "simple",
                "explode": false,
                "schema": {
                  "type": "string",
                  "example": "application/json"
                }
              }
            },
            "content": {
              "application/json": {
                "schema": {
                  "type": "object"
                },
                "examples": {
                  "example-0": {
                    "summary": "success",
                    "value": {
                      "status": true,
                      "data": {
                        "message": "Busca realizada com sucesso.",
                        "response": {
                          "Adult": false,
                          "backdrop_path": "/tmU7GeKVybMWFButWEGl2M4GeiP.jpg",
                          "belongs_to_collection": {
                            "ID": 230,
                            "Name": "The Godfather Collection",
                            "poster_path": "/zqV8MGXfpLZiFVObLxpAI7wWonJ.jpg",
                            "backdrop_path": "/mDMCET9Ens5ANvZAWRpluBsMAtS.jpg"
                          },
                          "Budget": 6000000,
                          "Genres": [
                            {
                              "ID": 18,
                              "Name": "Drama"
                            },
                            {
                              "ID": 80,
                              "Name": "Crime"
                            }
                          ],
                          "Homepage": "http://www.thegodfather.com/",
                          "ID": 238,
                          "imdb_id": "tt0068646",
                          "original_language": "en",
                          "original_title": "The Godfather",
                          "Overview": "Spanning the years 1945 to 1955, a chronicle of the fictional Italian-American Corleone crime family. When organized crime family patriarch, Vito Corleone barely survives an attempt on his life, his youngest son, Michael steps in to take care of the would-be killers, launching a campaign of bloody revenge.",
                          "Popularity": 290.926,
                          "poster_path": "/3bhkrj58Vtu7enYsRolD1fZdja1.jpg",
                          "production_companies": [
                            {
                              "ID": 4,
                              "Name": "Paramount Pictures",
                              "logo_path": "/gz66EfNoYPqHTYI4q9UEN4CbHRc.png",
                              "origin_country": "US"
                            },
                            {
                              "ID": 10211,
                              "Name": "Alfran Productions",
                              "logo_path": "",
                              "origin_country": "US"
                            },
                            {
                              "ID": 70,
                              "Name": "American Zoetrope",
                              "logo_path": "/ueaENQkPcy8rlr5fGZVBXKOhlBh.png",
                              "origin_country": "US"
                            }
                          ],
                          "production_countries": [
                            {
                              "iso_3166_1": "US",
                              "Name": "United States of America"
                            }
                          ],
                          "release_date": "1972-03-14",
                          "Revenue": 245066411,
                          "Runtime": 175,
                          "spoken_languages": [
                            {
                              "iso_639_1": "en",
                              "Name": "English"
                            },
                            {
                              "iso_639_1": "it",
                              "Name": "Italiano"
                            },
                            {
                              "iso_639_1": "la",
                              "Name": "Latin"
                            }
                          ],
                          "Status": "Released",
                          "Tagline": "An offer you can't refuse.",
                          "Title": "The Godfather",
                          "Video": false,
                          "vote_average": 8.69,
                          "vote_count": 20249
                        }
                      }
                    }
                  },
                  "example-1": {
                    "summary": "generic error",
                    "value": {
                      "status": false,
                      "data": {
                        "message": "Erro ao buscar detalhes do filme.",
                        "response": {
                          "log_id": "70fcd8c7-da8f-4337-bb42-ddbac29d97cc",
                          "event_date": "2024-08-22 22:50:32"
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {},
    "securitySchemes": {
      "basicAuth": {
        "type": "http",
        "scheme": "basic"
      }
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
