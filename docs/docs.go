// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://tos.santoshk.dev",
        "contact": {
            "name": "M Nurilman Baehaqi",
            "url": "https://twitter.com/MOXSPOY",
            "email": "mnurilmanbaehaqi@gmail.com"
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
        "/app-info": {
            "get": {
                "description": "This endpoint does not require token (public)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App Config"
                ],
                "summary": "Show application info metadata as the startup info while client app is launched",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/database_model.AppConfigResponse"
                        }
                    }
                }
            }
        },
        "/auth/check-credential": {
            "post": {
                "description": "Usually this endpoint used before validate user's identity",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Check whether email or phone number exist on the database_model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email/Phone Number",
                        "name": "credential",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.CheckCredentialResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh-token": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Use this endpoint to refresh token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.JwtResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "This endpoint used to register user to the platform",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign up endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "app_build_version",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "app_name_version",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "credential",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "device_identifier",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "device_model",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "os_version",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "phone_number",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "platform",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/auth/validate-identity": {
            "post": {
                "description": "This endpoint used to check user in the platform (for login)",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Validate identity is like for login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "app_build_version",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "app_name_version",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "credential",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "device_identifier",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "device_model",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "os_version",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "platform",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "request_id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.JwtResponse"
                        }
                    }
                }
            }
        },
        "/otp/check": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used to authenticate user when doing some sensitive data changes",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Check OTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OTP from SMS/Whatsapp to be checked",
                        "name": "otp",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/otp/request": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used to authenticate user when doing some sensitive data changes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Request OTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/universities/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint used to fetch all universities data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all universities object",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/universities/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint used to create university when it is not available yet on the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create university object.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "acronym",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "State owned (1) or private (2)",
                        "name": "category",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "createdAt",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "indonesiaAreaID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "nameEn",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "Active (1) or Not Active (2)",
                        "name": "status",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "updatedAt",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "website",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "yearFounded",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/user-detail/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint used to fetch user's data but with more detail",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user detail object",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/database_model.User"
                        }
                    }
                }
            }
        },
        "/user-detail/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used to process kyc",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user's selfie image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "User image that will be saved to the database_model",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint used to fetch user's data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user object",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/database_model.User"
                        }
                    }
                }
            }
        },
        "/user/update-email": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used because user fill phone number first. Note that user need to request otp first",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user's email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email that will be saved to the database_model",
                        "name": "new_email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OTP for authentication (if pin already exist)",
                        "name": "otp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/user/update-phone-number": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used as part of onboarding. Note that it should contain country code like +6285911110000. Note that user need to request otp first.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user's phone number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone number that will be saved to the database_model",
                        "name": "phone_number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OTP for authentication (if pin already exist)",
                        "name": "otp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/user/update-pin": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used as part of onboarding.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user's pin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "6 digits that will be saved to the database_model",
                        "name": "pin",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OTP for authentication (if pin already exist)",
                        "name": "otp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/user/update-user-name": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used as part of onboarding.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user's name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name that will be saved to the database_model",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        },
        "/user/verify-phone-number": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Usually this endpoint used as part of onboarding. Note that user need to request otp first.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Verify user's phone number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone number that will be verified",
                        "name": "phone_number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "OTP for authentication (if pin already exist)",
                        "name": "otp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api_response_model.SuccessAPIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api_response_model.CheckCredentialResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_email_verified": {
                    "type": "boolean"
                },
                "is_phone_verified": {
                    "type": "boolean"
                },
                "is_pin_registered": {
                    "type": "boolean"
                },
                "is_registered": {
                    "type": "boolean"
                },
                "is_user_exist": {
                    "type": "boolean"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "api_response_model.JwtResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "expire": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api_response_model.SuccessAPIResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "database_model.AppConfigResponse": {
            "type": "object",
            "properties": {
                "customer_friend_phone_number": {
                    "type": "string"
                },
                "maintenance": {
                    "type": "string"
                },
                "minimum_version": {
                    "type": "string"
                },
                "operating_time_weekday": {
                    "type": "string"
                },
                "operating_time_weekend": {
                    "type": "string"
                }
            }
        },
        "database_model.IndonesiaArea": {
            "type": "object",
            "required": [
                "code",
                "name"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "database_model.University": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "acronym": {
                    "type": "string"
                },
                "category": {
                    "description": "State owned (1) or private (2)",
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "indonesiaArea": {
                    "$ref": "#/definitions/database_model.IndonesiaArea"
                },
                "indonesiaAreaID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nameEn": {
                    "type": "string"
                },
                "status": {
                    "description": "Active (1) or Not Active (2)",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                },
                "yearFounded": {
                    "type": "integer"
                }
            }
        },
        "database_model.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "email_verification_status": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "phone_verification_status": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "university": {
                    "$ref": "#/definitions/database_model.University"
                },
                "universityID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userDetail": {
                    "$ref": "#/definitions/database_model.UserDetail"
                },
                "userDetailID": {
                    "type": "integer"
                }
            }
        },
        "database_model.UserDetail": {
            "type": "object",
            "properties": {
                "address_line_1": {
                    "type": "string"
                },
                "address_line_2": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "document_number": {
                    "type": "string"
                },
                "document_type": {
                    "type": "string"
                },
                "expiry_date": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "identity_and_selfie_image_url": {
                    "type": "string"
                },
                "identity_image_url": {
                    "type": "string"
                },
                "indonesiaArea": {
                    "$ref": "#/definitions/database_model.IndonesiaArea"
                },
                "indonesiaAreaID": {
                    "type": "integer"
                },
                "issue_date": {
                    "type": "string"
                },
                "selfie_image_url": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "integer"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "sql.NullString": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if String is not NULL",
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8083",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Collagen Web Service",
	Description:      "Web service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
