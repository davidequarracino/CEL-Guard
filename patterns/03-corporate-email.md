# 🛡️ Pattern 03: Corporate Email Restriction
Ensures that only users with a specific corporate email domain can access the internal systems.
**CEL Expression:**
`request.auth.claims.email.endsWith('@google.com')`