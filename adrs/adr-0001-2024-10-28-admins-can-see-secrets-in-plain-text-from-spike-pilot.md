+++
#   \\ 
#  \\\\ SPIKE: Keep your secrets secret with SPIFFE.
# \\\\\\

title = "ADR-0001: Admins can see secrets in plain text from SPIKE Pilot CLI"
+++

- Status: draft
- Date: 2024-10-28
- Tags: Security, Operations, Convenience

## Context and Problem Statement

We might want to encrypt the secrets that we show through the SPIKE Pilot's 
admin interface. This will enable end-to-end encryption and allow things like
inadvertently displaying secrets in logs or history.

However, this causes more harm than good. Besides, if the intention of the
admin user is to decrypt the secret, they likely have access to the decryption
keys, and they can decrypt the encrypted secret anyway resulting in the very
same issues that we were trying to avoid while making admins mildly inconvenient
due to the additional decryption step.

If an admin can write and delete secrets, it make sense for them to view the
secrets too.

An alternative is to "never" allow secrets to be read from the admin interface,
but that will end up admins creating throwaway secret consumer apps to 
"just read" secrets, installing them on the system, and increasing the attack
surface even more.

Besides:

* SPIKE Pilot relies on mTLS to fetch secrets.
* SPIKE Pilot enforces short-lived sessions.
* SPIKE Nexus implements audit logging to track access.
* So, if an entity is authenticated and authorized, and has adequate policy to
  read a secret value; then they can get the secret value in plain.

## Decision Outcome

Display secrets in plain text in the SPIKE Pilot admin cli; but also have an
interface to just display keys, or metadata instead of secrets.