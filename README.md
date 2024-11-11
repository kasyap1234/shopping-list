A RESTful API service for managing shopping lists built with Go and Encore.

Features
Create shopping list items
Get all items
Get single item by ID
Update item details
PostgreSQL database for persistence
Prerequisites
Install Encore
Go 1.21+

encore app create my-shopping-list
git clone <repository-url>
encore run


Access the local development dashboard at http://localhost:9400/

API Endpoints and Usage
Get all items:


curl http://localhost:4000/items


Get item by ID:

curl http://localhost:4000/items/1


Create new item:

curl -X POST http://localhost:4000/items \
  -H "Content-Type: application/json" \
  -d '{"name": "Milk", "price": 2.99, "quantity": 1, "bought": false}'



  Update item:

curl -X PUT http://localhost:4000/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Milk", "price": 3.99, "quantity": 2, "bought": false}'


Database Commands
Reset database:

encore db reset



Deployment
Deploy to Encore's development cloud:

git add .
git commit -m "Update message"
git push encore



Project Structure
.
├── shoppinglist/
│   ├── migrations/       # Database migrations
│   │   └── 1_create_tables.up.sql
│   └── main.go          # Service code and API endpoints




Built With
Go
Encore
GORM
PostgreSQL
License
MIT License

API Response Format
GET /items
{
    "items": [
        {
            "id": 1,
            "name": "Milk",
            "price": 2.99,
            "quantity": 1,
            "bought": false
        }
    ]
}



POST /items

{
    "success": true
}



The application uses Encore's built-in development environment. No additional configuration needed - just run encore run and start developing.

Contributing
Fork the repository
Create your feature branch
Commit your changes
Push to the branch
Create a new Pull Request