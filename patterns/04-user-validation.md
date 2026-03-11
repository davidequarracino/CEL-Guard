# Kubernetes ValidatingAdmissionPolicy - User Identity Check
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingAdmissionPolicy
metadata:
  name: "require-non-empty-username"
spec:
  failurePolicy: Fail
  validations:
    # Verifica che la stringa 'username' non sia vuota
    - expression: "size(request.auth.claims.username) > 0"
      message: "Identity error: Username claim cannot be empty."