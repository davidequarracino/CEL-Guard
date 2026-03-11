# Kubernetes ValidatingAdmissionPolicy - Role Check
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicy
metadata:
  name: "require-admin-role"
spec:
  failurePolicy: Fail
  validations:
    - expression: "request.auth.claims.role == 'admin'"
      message: "Access denied: admin role required for this operation."