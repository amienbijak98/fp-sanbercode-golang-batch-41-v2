# Simple Hospital Reservation System

This is a simple hospital reservation system built using the Gin framework and PostgreSQL. It allows users to register, login, and make appointments with doctors. It also allows admin users to manage patients, doctors, and departements. The system is built in Go language, and uses CRUD operations for database management. It also has a simple API documentation to make it easy for developers to use the system.

## Features

- JWT Auth. Login and logout.
- User admin registration
- CRUD Patients
- CRUD Doctors
- CRUD Departement
- CRUD Appointment

## Database ERD

![ERD Image](https://github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/blob/master/erd.png?raw=true)

## Endpoints

### Public Endpoints

- `POST /api/v1/register`: Allows user to register as admin
- `POST /api/v1/login`: Allows user to login as admin

### Private Endpoints

- `POST /api/v1/logout`: Allows user to logout as admin
- `GET /api/v1/admins`: Retrieve all admin data

#### Patients Endpoints

- `GET /api/v1/patients`: Retrieve all patient data
- `GET /api/v1/patients/:id`: Retrieve patient data by ID
- `POST /api/v1/patients`: Insert patient data
- `PUT /api/v1/patients/:id`: Update patient data by ID
- `DELETE /api/v1/patients/:id`: Delete patient data by ID

#### Doctors Endpoints

- `GET /api/v1/doctors`: Retrieve all doctor data
- `GET /api/v1/doctors/:id`: Retrieve doctor data by ID
- `POST /api/v1/doctors`: Insert doctor data
- `PUT /api/v1/doctors/:id`: Update doctor data by ID
- `DELETE /api/v1/doctors/:id`: Delete doctor data by ID

#### Departements Endpoints

- `GET /api/v1/departements`: Retrieve all department data
- `GET /api/v1/departements/:id`: Retrieve department data by ID
- `POST /api/v1/departements`: Insert department data
- `PUT /api/v1/departements/:id`: Update department data by ID
- `DELETE /api/v1/departements/:id`: Delete department data by ID

#### Appointments Endpoints

- `GET /api/v1/appointments`: Retrieve all appointment data
- `GET /api/v1/appointments/:id`: Retrieve appointment data by ID
- `POST /api/v1/appointments`: Insert appointment data
- `PUT /api/v1/appointments/:id`: Update appointment data by ID
- `DELETE /api/v1/appointments/:id`: Delete appointment data by ID

## Miscellaneous Links
`https://docs.google.com/presentation/d/1Id3tfO_fND_jRXOV4QCXPRpqODhUWeI-9gYt3I15v04/edit?usp=sharing` : Google Slide Presentation
`https://fp-sanbercode-golang-batch-41-v2-production.up.railway.app/` : Railway Deployment Link
