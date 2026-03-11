# 🛡️ Pattern 02: Admin Role Verification
Ensures that the user making the request has the 'admin' role assigned in their authentication token.
**CEL Expression:**
`request.auth.claims.role == 'admin'`