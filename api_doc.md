## API Documentation For FrostLand
---
v1
```
/api/v1/ping - Ping

/api/v1/user/create - Create a user
Form data:
    uid - Unique identification
    premium - Premium user or not [0 or 1]

/api/v1/query/nickname/:username - Query a user
Params:
    username - Unique identification

/api/v1/user/import - Import a user with existed info
Form data:
    uid - Unique
    uuid - Existed UUID
    premium - ...
```
---