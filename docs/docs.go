// Code generated by swaggo/swag. DO NOT EDIT
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
        "/order/:id": {
            "get": {
                "description": "Show order by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Show single order",
                "operationId": "order",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.OrderView"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "view.Delivery": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "view.Item": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "chrt_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nm_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "rid": {
                    "type": "string"
                },
                "sale": {
                    "type": "integer"
                },
                "size": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                }
            }
        },
        "view.OrderView": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "date_created": {
                    "type": "string"
                },
                "delivery": {
                    "$ref": "#/definitions/view.Delivery"
                },
                "delivery_service": {
                    "type": "string"
                },
                "entry": {
                    "type": "string"
                },
                "internal_signature": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.Item"
                    }
                },
                "locale": {
                    "type": "string"
                },
                "oof_shard": {
                    "type": "string"
                },
                "order_uid": {
                    "type": "string"
                },
                "payment": {
                    "$ref": "#/definitions/view.Payment"
                },
                "shardkey": {
                    "type": "string"
                },
                "sm_id": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                }
            }
        },
        "view.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "bank": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "custom_fee": {
                    "type": "integer"
                },
                "delivery_cost": {
                    "type": "integer"
                },
                "goods_total": {
                    "type": "integer"
                },
                "payment_dt": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "transaction": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Go Clean Template API",
	Description:      "Using a translation service as an example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
