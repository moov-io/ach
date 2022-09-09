---
layout: page
title: Prometheus metrics
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Metrics

The Moov ACH HTTP server reports Prometheus metrics when running. These are available on the admin server at `:9090/metrics` in the Prometheus metric format.

## Files

There are two metrics for ACH files:

- `ach_files_created{destination="...", origin="..."}`
- `ach_files_deleted`

Example:

```
$ curl localhost:9090/metrics
# HELP ach_files_created The number of ACH files created
# TYPE ach_files_created counter
ach_files_created{destination="231380104",origin="121042882"} 5
# HELP ach_files_deleted The number of ACH files deleted
# TYPE ach_files_deleted counter
ach_files_deleted 1
```
