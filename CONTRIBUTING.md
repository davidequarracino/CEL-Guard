# Contributing to CEL-Guard

Thank you for your interest in improving CEL-Guard. As a project focused on security policy enforcement, we maintain high standards for code quality and policy accuracy.

## Development Workflow
1. **Fork the repository** and create your branch from `main`.
2. **Implement your CEL pattern** inside the `/patterns` directory using the Kubernetes `ValidatingAdmissionPolicy` YAML format.
3. **Verify your expression** using the local Go validator:
   ```bash
 go run validator/main.go "request.auth.claims.role == 'admin' && request.auth.claims.email.endsWith('@google.com')"