// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-11 19:07:33.8007498 +0500 PKT m=+2.196850901

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "save microservices and deploy services on kubernetes cluster",
        "title": "Kubernetes Manifest Deployment Engine",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Cloudplex Support",
            "url": "http://www.cloudplex.io/support",
            "email": "haseeb@cloudplex.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/ksd/api/v1",
    "paths": {
        "/api/v1/registry": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ]
            }
        },
        "/api/v1/registry/{namespace}/{name}": {
            "get": {
                "description": "deploy services on kubernetes cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "deploy services on kubernetes cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            },
            "delete": {
                "description": "deploy services on kubernetes cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "deploy services on kubernetes cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/api/v1/statefulsets/{namespace}": {
            "get": {
                "description": "get status of all kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get status of  all kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of kubernetes cluster",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/api/v1/statefulsets/{name}/{namespace}": {
            "get": {
                "description": "get status of kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get status of kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/v1.StatefulSet"
                        }
                    }
                }
            },
            "delete": {
                "description": "get status of kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get status of kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/deployment/{namespace}": {
            "get": {
                "description": "get status of all kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deployment"
                ],
                "summary": "get status of  all kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of kubernetes cluster",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/deployment/{name}/{namespace}": {
            "get": {
                "description": "get status of kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deployment"
                ],
                "summary": "get status of kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/kubeservice/{namespace}": {
            "get": {
                "description": "get status of all kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get status of  all kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of kubernetes cluster",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/kubeservice/{name}/{namespace}": {
            "get": {
                "description": "get status of kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get status of kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            },
            "delete": {
                "description": "get status of kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get status of kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        },
        "/kubeservice/{name}/{namespace}/endpoint": {
            "get": {
                "description": "get status of kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get status of kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "project_id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"error\": \"\", \"external_ip\": \"\"}"
                    },
                    "404": {
                        "description": "{\"error\": \"\"}"
                    },
                    "500": {
                        "description": "{\"error\": \"\"}"
                    }
                }
            }
        },
        "/solution": {
            "get": {
                "description": "deploy services on kubernetes cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "solutions"
                ],
                "summary": "deploy services on kubernetes cluster",
                "parameters": [
                    {
                        "description": "body for services deployment",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/types.ServiceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"service\": map[string]interface{},\"project_id\":\"\"}"
                    },
                    "404": {
                        "description": "{\"error\": \"\"}"
                    },
                    "500": {
                        "description": "{\"error\": \"\"}"
                    }
                }
            },
            "put": {
                "description": "deploy services on kubernetes cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "solutions"
                ],
                "summary": "deploy services on kubernetes cluster",
                "parameters": [
                    {
                        "description": "body for services deployment",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/types.ServiceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"service\": map[string]interface{},\"project_id\":\"\"}"
                    },
                    "404": {
                        "description": "{\"error\": \"\"}"
                    },
                    "500": {
                        "description": "{\"error\": \"\"}"
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "solutions"
                ],
                "summary": "deploy services on kubernetes cluster",
                "parameters": [
                    {
                        "description": "body for services deployment",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/types.ServiceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"service\": map[string]interface{},\"project_id\":\"\"}"
                    },
                    "404": {
                        "description": "{\"error\": \"\"}"
                    },
                    "500": {
                        "description": "{\"error\": \"\"}"
                    }
                }
            },
            "delete": {
                "description": "deploy services on kubernetes cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "solutions"
                ],
                "summary": "deploy services on kubernetes cluster",
                "parameters": [
                    {
                        "description": "body for services deployment",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/types.ServiceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"service\": map[string]interface{},\"project_id\":\"\"}"
                    },
                    "404": {
                        "description": "{\"error\": \"\"}"
                    },
                    "500": {
                        "description": "{\"error\": \"\"}"
                    }
                }
            },
            "patch": {
                "description": "deploy services on kubernetes cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "solutions"
                ],
                "summary": "deploy services on kubernetes cluster",
                "parameters": [
                    {
                        "description": "body for services deployment",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/types.ServiceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"service\": map[string]interface{},\"project_id\":\"\"}"
                    },
                    "404": {
                        "description": "{\"error\": \"\"}"
                    },
                    "500": {
                        "description": "{\"error\": \"\"}"
                    }
                }
            }
        },
        "/statefulsets/{name}/{namespace}": {
            "delete": {
                "description": "get status of kubernetes services deployment on a Kubernetes Cluster. If you need all services status then pass namespace=\"\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deployment"
                ],
                "summary": "get status of kubernetes services deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "project id",
                        "name": "project_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the kubernetes service",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace of the kubernetes service",
                        "name": "namespace",
                        "in": "path",
                        "required": true
                    }
                ]
            }
        }
    },
    "definitions": {
        "types.ServiceRequest": {
            "type": "object",
            "required": [
                "project_id"
            ],
            "properties": {
                "project_id": {
                    "type": "string"
                },
                "service": {
                    "type": "object",
                    "required": [
                        "service"
                    ]
                }
            }
        },
        "v1.StatefulSet": {
            "type": "object",
            "properties": {
                "spec": {
                    "type": "StatefulSetSpec"
                },
                "status": {
                    "type": "StatefulSetStatus"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
