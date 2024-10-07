Sim Card Management API

This project is a Sim Card Management API built using Golang, the Gin web framework, and GORM for database operations. The API provides functionality to activate, deactivate, and retrieve sim card details.

Features

Activate Sim Card: Create and activate a new sim card.
Deactivate Sim Card: Deactivate an existing sim card.
Retrieve Sim Card Details: Get the details of a sim card by its number.
Tech Stack

Golang: Programming language used to build the API.
Gin: Web framework used to handle routes and requests.
GORM: ORM (Object Relational Mapping) library for interacting with the database.
PostgreSQL: Database used for storing sim card records.








API Endpoints

**1. Activate a Sim Card**
**Endpoint: POST /activate**
Description: Activates a new sim card and stores it in the database.
Request Body:

{
  "simnumber": "1234567890",
  "phonno": "9876543210"
}


{
  "message": "Sim activated successfully",
  "sim": {
    "id": 1,
    "simnumber": "1234567890",
    "phonno": "9876543210",
    "status": "active",
    "created_at": "2024-08-25T10:20:30.000Z"
  }
}
**2. Deactivate a Sim Card
Endpoint: POST /deactivate?phoneno={phoneno}**

{
  "message": "Sim deactivated successfully",
  "sim": {
    "id": 1,
    "simnumber": "1234567890",
    "phonno": "9876543210",
    "status": "inactive",
    "created_at": "2024-08-25T10:20:30.000Z"
  }
}
**3. Get Sim Card Details
Endpoint: GET /simdetails/{simNumber}**

{
  "sim_details": {
    "id": 1,
    "simnumber": "1234567890",
    "phonno": "9876543210",
    "status": "active",
    "created_at": "2024-08-25T10:20:30.000Z"
  }
}

type SimcardSchema struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    SimNumber string    `json:"simnumber" gorm:"unique;not null"`
    Phoneno   string    `json:"phonno" gorm:"unique;not null"`
    Status    string    `json:"status" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
}






