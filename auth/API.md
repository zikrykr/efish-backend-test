openapi: 3.0.0
info:
  version: 1.0.0
  title: Auth-App
  description: A simple collections of API for auth-app 

servers:
  - url: http://localhost:3000/v1/auth

components:
  securitySchemes:
    bearerAuth:            
      type: http
      scheme: bearer
      bearerFormat: JWT 

security:
  - bearerAuth: []  

paths:
  /register:
    post:
      description: Register new user
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                phone:
                  type: string
                  description: User's phone number
                name:
                  type: string
                  description: User's username
                role:
                  type: string
                  description: User's role
      response:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: object
                properties:
                    phone:
                        type: string
                        description: User's phone number
                    name:
                        type: string
                        description: User's username
                    role:
                        type: string
                        description: User's role
                    createdAt:
                        type: string
                        description: User's registration time
                    password:
                        type: string
                        description: User's generated password
        
        '500'
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message

        default:
          description: Unregistered error/response

  /login:
    get:
      description: Users Login to get access token 
      response:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: object
                properties:
                    phone:
                        type: string
                        description: User's phone number
                    accessToken:
                        type: string
                        description: User's access token
        
        '500'
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message
        default:
          description: Unregistered error/response

/verify-token:
    get:
      description: Verify JWT Token and return its private claim:
        - bearerAuth: []
      response:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: object
                properties:
                    phone:
                        type: string
                        description: User's phone number
                    name:
                        type: string
                        description: User's username
                    role:
                        type: string
                        description: User's role
                    createdAt:
                        type: string
                        description: String formatted timestamp 
                    exp:
                        type: string
                        description: String formatted timestamp 
                    iat:
                        type: string
                        description: String formatted timestamp 
        
        '500'
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message