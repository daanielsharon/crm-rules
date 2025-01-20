# crm-rules

## Prerequisites

Before getting started, ensure you have the following installed:
- [Docker](https://www.docker.com/get-started)
- [Make](https://www.gnu.org/software/make/)

## Project Setup

1. Initialize the project using Makefile:
```bash
make start
```

> **Note**: During the initialization, you'll be prompted for various configuration settings. 
> - You can press `Enter` to skip and use default values
> - Alternatively, provide custom configuration as needed
> - Default values are pre-configured for local development

## Troubleshooting

### Database Connection Issues

If you encounter errors related to data fetching or database connections:

- **First-time Setup**: The database comes pre-seeded with initial data
- **Common Causes**:
  - Postgres container might still be initializing

### Database Access
> - All database connection details are set during initial setup
> - Check your `.env` file for the exact configuration
> - Ensure Docker containers are running before attempting to connect

### Accessing the Application

- **Web Interface**: Open `http://localhost:80` in your browser
- **Restart**: If issues persist, try `make restart` to restart the application


[Back to README](../README.md)
