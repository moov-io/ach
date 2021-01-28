---
layout: page
title: Overview
description: fill
hide_hero: true
show_sidebar: false
menubar: getting-started-menu
---

# Overview

![Moov ACH Logo](https://repository-images.githubusercontent.com/76497520/263dab00-c6d9-11ea-8bf0-8070d91f9135)

ACH implements a reader, writer, and validator for Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) files. ACH is the primary method of electronic money movement throughout the United States. The HTTP server is available in a [Docker image](/usage-docker.md) and the Go package `github.com/moov-io/ach` is available.

If you're looking for a complete implementation of ACH origination (file creation), OFAC checks, micro-deposits, SFTP uploading, and other features, the [moov-io/paygate](https://github.com/moov-io/paygate) project aims to be a full system for ACH transfers. Otherwise, check out our article on [How and When to use the Moov ACH Library](https://moov.io/blog/tutorials/how-and-when-to-use-the-moov-ach-library/).