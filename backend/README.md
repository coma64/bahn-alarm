## Migrations

```bash
# Create migration
CONFIGOR_ENV=dev go run . migrate create add-notification-departure sql

# Apply migrations
CONFIGOR_ENV=dev go run . migrate up
```