// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/vouchers": {
            "post": {
                "description": "Create a new voucher with code, description, usage limit, and expiry date.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "Create a new voucher",
                "parameters": [
                    {
                        "description": "Create Voucher Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateVoucherRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/vouchers/redeem": {
            "patch": {
                "description": "Redeem a voucher code for a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "Redeem a voucher",
                "parameters": [
                    {
                        "description": "Redeem Voucher Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RedeemVoucherRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/vouchers/users/{user_id}/history": {
            "get": {
                "description": "Get a list of voucher redemption histories filtered by user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "List redeemed voucher histories by user ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
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
                                "$ref": "#/definitions/entity.VoucherRedemptionHistory"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/vouchers/{code}/history": {
            "get": {
                "description": "Get a list of voucher redemption histories filtered by voucher code.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vouchers"
                ],
                "summary": "List redeemed voucher histories by code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Voucher Code",
                        "name": "code",
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
                                "$ref": "#/definitions/entity.VoucherRedemptionHistory"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/serr.ErrorCode"
                },
                "message": {
                    "type": "string"
                },
                "trace_id": {
                    "type": "string"
                }
            }
        },
        "dto.CreateVoucherRequest": {
            "type": "object",
            "required": [
                "amount",
                "code",
                "description",
                "expiry_date",
                "usage_limit"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expiry_date": {
                    "type": "string"
                },
                "usage_limit": {
                    "type": "integer"
                }
            }
        },
        "dto.RedeemVoucherRequest": {
            "type": "object",
            "required": [
                "code",
                "user_id"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.VoucherRedemptionHistory": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "Amount of voucher",
                    "type": "integer"
                },
                "id": {
                    "description": "Unique identifier for the redemption record",
                    "type": "integer"
                },
                "redeemedAt": {
                    "description": "When the voucher was Redeemed",
                    "type": "string"
                },
                "userID": {
                    "description": "ID of the user who Redeemed the voucher",
                    "type": "string"
                },
                "voucherID": {
                    "description": "ID of the Redeemed voucher",
                    "type": "integer"
                }
            }
        },
        "serr.ErrorCode": {
            "type": "string",
            "enum": [
                "INTERNAL",
                "INVALID_VOUCHER",
                "REACH_LIMIT",
                "INVALID_USER",
                "INVALID_TIME",
                "INVALID_INPUT"
            ],
            "x-enum-varnames": [
                "ErrInternal",
                "ErrInvalidVoucher",
                "ErrReachLimit",
                "ErrInvalidUser",
                "ErrInvalidTime",
                "ErrInvalidInput"
            ]
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