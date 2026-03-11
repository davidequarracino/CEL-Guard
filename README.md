# CEL-Guard

A repository of validated **Common Expression Language (CEL)** patterns for Kubernetes Admission Control and Cloud-Native security policies.

## Project Structure
- `/patterns`: Production-ready `ValidatingAdmissionPolicy` examples.
- `/validator`: A Go-based CLI utility to compile and verify CEL expressions against the `google/cel-go` engine.

## Getting Started
To validate a pattern locally using the Go engine:
```bash
go run validator/main.go "request.auth.claims.role == 'admin'"