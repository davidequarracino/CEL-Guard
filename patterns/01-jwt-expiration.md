# 🛡️ Pattern 01: JWT Expiration Validation
Prevents the use of infinite or malformed authentication tokens.
**CEL Expression:**
`request.auth.claims.exp > 0`