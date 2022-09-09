---
layout: page
title: Overview
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Overview

![Moov ACH Logo](https://repository-images.githubusercontent.com/76497520/263dab00-c6d9-11ea-8bf0-8070d91f9135)

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

ACH implements a reader, writer, and validator for Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) files. ACH is the primary method of electronic money movement throughout the United States. The HTTP server is available in a [Docker image](/usage-docker.md) and the Go package `github.com/moov-io/ach` is available.

If you're looking for an event driven ACH engine for uploading/downloading files and operations we have built [moov-io/achgateway](https://github.com/moov-io/achgateway) and run it in production. Our article [How and When to use the Moov ACH Library](https://moov.io/blog/education/how-and-when-to-use-the-moov-ach-library/) will help to generate ACH files for upload to your ODFI.
