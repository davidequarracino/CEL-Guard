# Kubernetes ValidatingAdmissionPolicy - Domain Check
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicy
metadata:
  name: "restrict-corporate-domain"
spec:
  failurePolicy: Fail
  validations:
    - expression: "request.auth.claims.email.endsWith('@google.com')"
      message: "Access denied: corporate email domain required."