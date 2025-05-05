# API Endpoints To-Do List

## Authentication
- [ ] `POST /api/auth/login`
  - Purpose: User login
  - Body: `{ email: string, password: string }`
  - Response: `{ token: string, user: User }`

- [ ] `POST /api/auth/signup`
  - Purpose: User registration
  - Body: `{ email: string, password: string, name: string }`
  - Response: `{ token: string, user: User }`

## Route Search
- [ ] `GET /api/routes/search`
  - Purpose: Search for bike routes based on criteria
  - Query Parameters:
    - `address`: string (location)
    - `difficulty`: string (Easy/Medium/Hard)
    - `distance`: string (distance range)
  - Response: Array of RouteOption objects

- [ ] `GET /api/routes/{id}`
  - Purpose: Get detailed information about a specific route
  - Parameters:
    - `id`: number (route ID)
  - Response: Detailed route information including:
    - title
    - description
    - distance
    - duration
    - difficulty
    - imageUrl
    - highlights
    - mapUrl

## Location Services
- [ ] `GET /api/location/reverse-geocode`
  - Purpose: Convert coordinates to address
  - Query Parameters:
    - `lat`: number
    - `lng`: number
  - Response: `{ address: string }`

## User Profile
- [ ] `GET /api/user/profile`
  - Purpose: Get user profile information
  - Headers: `Authorization: Bearer {token}`
  - Response: User profile data

- [ ] `PUT /api/user/profile`
  - Purpose: Update user profile
  - Headers: `Authorization: Bearer {token}`
  - Body: User profile update data
  - Response: Updated user profile

## Route Favorites
- [ ] `POST /api/routes/{id}/favorite`
  - Purpose: Add a route to favorites
  - Headers: `Authorization: Bearer {token}`
  - Parameters:
    - `id`: number (route ID)
  - Response: Success status

- [ ] `DELETE /api/routes/{id}/favorite`
  - Purpose: Remove a route from favorites
  - Headers: `Authorization: Bearer {token}`
  - Parameters:
    - `id`: number (route ID)
  - Response: Success status

- [ ] `GET /api/routes/favorites`
  - Purpose: Get user's favorite routes
  - Headers: `Authorization: Bearer {token}`
  - Response: Array of favorite routes

## Route History
- [ ] `GET /api/routes/history`
  - Purpose: Get user's route history
  - Headers: `Authorization: Bearer {token}`
  - Response: Array of previously viewed routes

## Implementation Notes
For each endpoint:
- [ ] Include proper error handling
- [ ] Return appropriate HTTP status codes
- [ ] Include input validation
- [ ] Use proper authentication where required
- [ ] Follow RESTful conventions
- [ ] Include proper documentation 