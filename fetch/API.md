openapi: 3.0.0
info:
  version: 1.0.0
  title: Fetch-App
  description: A simple collections of API for fetch-app 

servers:
  - url: http://localhost:3000/v1/fetch

components:
  securitySchemes:
    bearerAuth:            
      type: http
      scheme: bearer
      bearerFormat: JWT 

security:
  - bearerAuth: []  

paths:
  /resources:
    get:
      description: Get list of clean commodities data with additional USD currency for its price 
      response:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: array
                    description: List of all commodities data with additional USD currency for its price
                properties:
                    uuid:
                        type: string
                        description: Commodity data ID
                    komoditas:
                        type: string
                        description: Commodity name/type
                    area_provinsi:
                        type: string
                        description: Commodity sell area, province level
                    area_kota:
                        type: string
                        description: Commodity sell area, city level
                    size:
                        type: string
                        description: Commodity's size/amount
                    price:
                        type: string
                        description: Commodity's each item's price in rupiah
                    tgl_parsed: 
                        type: string
                        description: Commodity's parsed date time
                    timestamp:
                        type: string
                        description: Timestamp
                    price_usd:
                        type: string
                        description: Commodity's each item's price in USD   
        
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

  /resources/aggregate:
    get:
      description: Get list of aggregated and clean commodities data based on area_province and weekly data with max, min, avg, median of data 
      response:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: array
                    description: list of aggregated and clean commodities data based on area_province and weekly data with max, min, avg, median of data
                properties:
                    area_provinsi:
                        type: string
                        description: Commodity sell area, province level
                    minggu: 
                        type: int
                        description: Weekly 
                    tahun: 
                        type: int
                        description: Year 
                    min:
                        type: float
                        description: lowest weekly profit of the province 
                    max:
                        type: float
                        description: highest weekly profit of the province
                    avg:
                        type: float
                        description: average weekly profit of the province
                    median:
                        type: float
                        description: median of weekly profit of the province
                    totalPrice:
                        type: float
                        description: total price of the province
                    totalData:
                        type: float
                        description: count commodities of the province
                    prices:
                        type: []float
                        description: list commodities price of the province, weekly
        
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