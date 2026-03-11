# Kubernetes ValidatingAdmissionPolicy Example
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicy
metadata:
  name: "force-token-expiration"
spec:
  failurePolicy: Fail
  matchConstraints:
    resourceRules:
    - apiGroups: ["*"]
      apiVersions: ["*"]
      operations: ["CREATE", "UPDATE"]
      resources: ["*"]
  validations:
    - expression: "request.auth.claims.exp > 0"
      message: "Security violation: tokens must have a valid expiration claim."