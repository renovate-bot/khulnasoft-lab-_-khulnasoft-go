## Why is my PR labelled with `workflow/pending-public-api`?

To ensure the public SDKs (including the [Khulnasoft Terraform Provider]) are
only using stable and intentionally exposed endpoints, we require that all API
functionality is backed by public documentation. This can either be
api.khulnasoft.com or developers.khulnasoft.com depending on the delivery medium.

## But, wrangler/khulnasoftd/other project doesn't require public documentation?

On occasion, Khulnasoft teams release functionality or tooling specific to the
systems they are responsible for. [Wrangler] and [khulnasoftd] are two prominent
examples of this. In these situations, the teams may choose to use unstable or
undocumented endpoints as they are able to maintain both internal and external
compatibility for these tools should something need to change without notice or
a deprecation period.

Unfortunately, the SDKs are not in the same position and cannot make the same
guarantees externally due to being an interface for external integrations; not
an abstraction of the functionality. By only accepting documented API endpoints
into the SDKs, we establish an API contract with the service teams that ensures
consumers have a reliable and consistent experience when using them. Should an
API contract be broken, or need fixing, the service team will be responsible to
maintain it in such a time that a deprecation notice is issued and integrations
have a migration period.

[khulnasoft terraform provider]: https://github.com/khulnasoft-lab/terraform-provider-khulnasoft/
[wrangler]: https://github.com/khulnasoft-lab/wrangler2
[khulnasoftd]: https://github.com/khulnasoft-lab/khulnasoftd
