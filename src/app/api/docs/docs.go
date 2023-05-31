// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "DIT - IFAL",
            "email": "evs10@aluno.ifal.edu.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/artists": {
            "get": {
                "description": "Rota que permite que se busque todos os artistas do banco",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas do usuário"
                ],
                "summary": "Buscar todos os artistas do banco",
                "operationId": "GetArtists",
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.ArtistDTO"
                            }
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/artists/{artistID}": {
            "get": {
                "description": "Rota que permite que se busque as informações de um artista",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas do usuário"
                ],
                "summary": "Buscar os dados de um artista por id",
                "operationId": "GetArtistInformation",
                "parameters": [
                    {
                        "type": "string",
                        "default": "a6d6488b-6ed0-4d41-9c55-4e899fdd7e47",
                        "description": "ID do artista.",
                        "name": "artistID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.ArtistLowDTO"
                            }
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/artists/{artistID}/songs": {
            "get": {
                "description": "Rota que permite que se busque todas as músicas de um artista",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas do usuário"
                ],
                "summary": "Buscar todas as músicas de um artista específico",
                "operationId": "GetArtistSongs",
                "parameters": [
                    {
                        "type": "string",
                        "default": "cfd8b073-a303-4886-836a-f249e88be9bc",
                        "description": "ID do artista.",
                        "name": "artistID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Requisição realizada com sucesso.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.SongDTO"
                            }
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Rota que permite que um usuário se autentique no sistema utilizando seu endereço de e-mail e senha.\n| E-mail              | Senha     | Função                                                            |\n|---------------------|-----------|-------------------------------------------------------------------|\n| admin@ifal.edu.br   | Test1234! | Usuário comum do sistema\t\t\t\t\t\t\t\t\t\t  |",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rotas de autenticação"
                ],
                "summary": "Fazer login no sistema",
                "operationId": "Login",
                "responses": {
                    "400": {
                        "description": "Requisição mal formulada.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "401": {
                        "description": "Usuário não autorizado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "403": {
                        "description": "Acesso negado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "422": {
                        "description": "Algum dado informado não pôde ser processado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Ocorreu um erro inesperado.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    },
                    "503": {
                        "description": "A base de dados não está disponível.",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.ArtistDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "founded_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ArtistDTO"
                    }
                },
                "name": {
                    "type": "string"
                },
                "super_artist_id": {
                    "type": "string"
                },
                "terminated_at": {
                    "type": "string"
                }
            }
        },
        "response.ArtistLowDTO": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "founded_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ArtistDTO"
                    }
                },
                "name": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                },
                "record_company_id": {
                    "type": "string"
                },
                "spotify_url": {
                    "type": "string"
                },
                "super_artist_id": {
                    "type": "string"
                },
                "terminated_at": {
                    "type": "string"
                }
            }
        },
        "response.ErrorMessage": {
            "type": "object",
            "properties": {
                "duplicated_fields": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "invalid_fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.InvalidField"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "response.InvalidField": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "field_name": {
                    "type": "string"
                }
            }
        },
        "response.SongDTO": {
            "type": "object",
            "properties": {
                "album_id": {
                    "type": "string"
                },
                "artist_id": {
                    "type": "string"
                },
                "duration": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "SPOTIFY RECORDS API",
	Description:      "Aplicação de artistas do spotify",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
