// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "kongphop.c@kuranasaki.work"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/post": {
            "put": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "UpdateUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "UpdateUser",
                        "name": "UpdateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/postAPI.UpdatePostRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Update User Success",
                        "schema": {
                            "$ref": "#/definitions/postAPI.UpdatePostResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Update User Failed",
                        "schema": {
                            "$ref": "#/definitions/postAPI.UpdatePostResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "Create new post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "description": "UpdateUser",
                        "name": "UpdateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/postAPI.CreatePostRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Update User Success",
                        "schema": {
                            "$ref": "#/definitions/postAPI.UpdatePostResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Update User Failed",
                        "schema": {
                            "$ref": "#/definitions/postAPI.UpdatePostResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "DeleteUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "responses": {
                    "201": {
                        "description": "Delete User Success",
                        "schema": {
                            "$ref": "#/definitions/postAPI.UpdatePostResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Delete User Failed",
                        "schema": {
                            "$ref": "#/definitions/postAPI.UpdatePostResponse"
                        }
                    }
                }
            }
        },
        "/post/list": {
            "get": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "ListPost",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/postAPI.GetPostByIDResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/{PostID}": {
            "get": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "Get User's Info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "PostID",
                        "name": "PostID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/postAPI.GetPostByIDResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "Upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Upload"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update User Success",
                        "schema": {
                            "$ref": "#/definitions/uploadAPI.UploadSuccessResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Update User Failed",
                        "schema": {
                            "$ref": "#/definitions/uploadAPI.UploadErrorResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "Get User's Info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userAPI.GetUserByIDResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "UpdateUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "UpdateUser",
                        "name": "UpdateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userAPI.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Update User Success",
                        "schema": {
                            "$ref": "#/definitions/userAPI.UpdateUserResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Update User Failed",
                        "schema": {
                            "$ref": "#/definitions/userAPI.UpdateUserResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Firebase": []
                    }
                ],
                "description": "DeleteUser",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {
                    "201": {
                        "description": "Delete User Success",
                        "schema": {
                            "$ref": "#/definitions/userAPI.UpdateUserResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid Token",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "403": {
                        "description": "No Permissions",
                        "schema": {
                            "$ref": "#/definitions/authMiddleware.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Delete User Failed",
                        "schema": {
                            "$ref": "#/definitions/userAPI.UpdateUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authMiddleware.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.TAddress": {
            "type": "object",
            "properties": {
                "address_1": {
                    "type": "string"
                },
                "amphoe": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "tambon": {
                    "type": "string"
                }
            }
        },
        "model.TPackage": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                }
            }
        },
        "model.TPost": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "crafterID": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "packageList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TPackage"
                    }
                },
                "reviewList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TReview"
                    }
                },
                "thumbnail": {
                    "$ref": "#/definitions/model.TThumbnail"
                }
            }
        },
        "model.TReview": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "ratingStar": {
                    "type": "number"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "model.TThumbnail": {
            "type": "object",
            "properties": {
                "thumbnailType": {
                    "type": "string"
                },
                "thumbnailUrl": {
                    "type": "string"
                }
            }
        },
        "model.TUser": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/model.TAddress"
                },
                "phone_number": {
                    "type": "string"
                },
                "pictureUrl": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "postAPI.CreatePostRequest": {
            "type": "object",
            "properties": {
                "post": {
                    "$ref": "#/definitions/postAPI.Post"
                }
            }
        },
        "postAPI.GetPostByIDResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "post": {
                    "$ref": "#/definitions/model.TPost"
                }
            }
        },
        "postAPI.Package": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                }
            }
        },
        "postAPI.Post": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "crafterID": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "packageList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/postAPI.Package"
                    }
                },
                "reviewList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/postAPI.Review"
                    }
                },
                "thumbnail": {
                    "$ref": "#/definitions/postAPI.Thumbnail"
                }
            }
        },
        "postAPI.Review": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "ratingStar": {
                    "type": "number"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "postAPI.Thumbnail": {
            "type": "object",
            "properties": {
                "thumbnailType": {
                    "type": "string"
                },
                "thumbnailUrl": {
                    "type": "string"
                }
            }
        },
        "postAPI.UpdatePostRequest": {
            "type": "object",
            "properties": {
                "post": {
                    "$ref": "#/definitions/postAPI.Post"
                }
            }
        },
        "postAPI.UpdatePostResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "uploadAPI.UploadErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "uploadAPI.UploadSuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "userAPI.Address": {
            "type": "object",
            "properties": {
                "address_1": {
                    "type": "string"
                },
                "amphoe": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "tambon": {
                    "type": "string"
                }
            }
        },
        "userAPI.GetUserByIDResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.TUser"
                }
            }
        },
        "userAPI.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/userAPI.User"
                }
            }
        },
        "userAPI.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "userAPI.User": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/userAPI.Address"
                },
                "phone": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "Firebase": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://crafty.kuranasaki.work/api/v1/auth/oauth/token",
            "scopes": {
                "admin": "All Permissions Granted",
                "mod": "All Permissions Except adding MOD",
                "user": "Permissions granted upto project owner"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "CraftyUserAPI Documents",
	Description:      "This is an auto-generated API Docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
