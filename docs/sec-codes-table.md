---
layout: page
title: SEC Codes Table
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

## Picking which Standard Entry Class code to use

The [NACHA Corporate Rules and Guidelines](https://github.com/moov-io/ach/blob/master/docs/2013-Corporate-Rules-and-Guidelines.pdf) offer a helpful table for choosing the correct Standard Entry Class (SEC) Code to use for a given enrollment and transaction. The table has been re-written below:

### Point of sale transactions

#### Physical enrollment

- Customer uses their debit card at a POS terminal.

<p>Proper SEC Code: <strong>POS</strong></p>

#### Internet enrollment

- Customer uses their debit card at a POS terminal.

<p>Proper SEC Code: <strong>POS</strong></p>

### Internet transactions

#### Physical enrollment

- Customer initiates debits via the bank’s web site or individual bill payments at a merchant web site.
- Customer uses their ACH-based debit card to make a purchase at a web site.

<p>Proper SEC Code: <strong>PPD</strong></p>

#### Internet enrollment

- Customer authorizes a transfer of funds into a savings account.
- Customer initiates bill payments at a merchant web site on their debit card.
- Customer uses their ACH-based debit card to make a purchase at a web site.

<p>Proper SEC Code: <strong>WEB</strong></p>

### ATM transactions

#### Physical enrollment

- Customer uses the card at an ATM to withdraw cash.

<p>Proper SEC Code: <strong>MTE</strong></p>

#### Internet enrollment

- Customer uses an ATM to withdraw cash:

<p>Proper SEC Code: <strong>MTE</strong></p>
