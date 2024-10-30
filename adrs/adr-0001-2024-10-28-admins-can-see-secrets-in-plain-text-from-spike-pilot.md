```text
  \\ 
 \\\\ SPIKE: Keep your secrets secret with SPIFFE.
\\\\\\
```

# ADR-0001: Display Secrets in Plain Text in SPIKE Pilot Admin CLI

- Status: final
- Date: 2024-10-28
- Tags: Security, Operations, Convenience

## Context and Problem Statement

The SPIKE Pilot admin interface needs to provide access to secrets for 
administrative purposes. We need to determine the most secure and practical way 
to display these secrets while maintaining system security and operational 
efficiency.

## Decision Drivers

* Security of sensitive information
* Operational efficiency for administrators
* Prevention of workarounds that could increase attack surface
* Auditability of secret access
* User experience for administrators

## Considered Options

1. Display secrets in plain text
2. Display only encrypted secrets
3. Never display secrets through admin interface

## Decision

Display secrets in plain text in the SPIKE Pilot admin CLI, while providing 
additional interfaces to view only keys or metadata when full secret values 
aren't needed.

## Rationale

* Administrators who can write and delete secrets should logically be able to 
  view them.
* Encrypting displayed secrets provides false security since:
  * Admins with access likely have decryption keys anyway
  * It only adds inconvenience without meaningful security benefits
* Preventing secret viewing would likely lead to:
  * Creation of throwaway secret consumer apps
  * Increased attack surface through workarounds
* Existing security measures provide adequate protection:
  * mTLS for secret fetching
  * Short-lived sessions
  * Audit logging through SPIKE Nexus
  * Authentication and authorization checks

## Consequences

### Positive

* Simplified admin operations
* Reduced likelihood of workarounds
* Clear audit trail of secret access
* Consistent with principle of least surprise

### Negative

* Potential for secrets to appear in logs or command history
* Increased responsibility on admin access control

## Implementation Notes

* Implement separate interfaces for:
  * Full secret display
  * Keys-only view
  * Metadata-only view
* Ensure proper audit logging of all secret access
* Document proper terminal/session management for admins
