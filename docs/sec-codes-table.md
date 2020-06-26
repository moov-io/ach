## Picking which Standard Entry Class Code code to use

The [NACHA Corporate Rules and Guidelines](https://github.com/moov-io/ach/blob/master/docs/2013-Corporate-Rules-and-Guidelines.pdf) offer a helpful table for choosing the correct Standard Entry Class (SEC) Code to use for a given enrollment and transaction. The table has been re-written below:

### Point of Sale Transactions

##### Physical Enrollment

- Customer uses their debit card at a POS terminal.

<p>Proper SEC Code: <strong>POS</strong></p>

##### Internet Enrollment

- Customer uses their debit card at a POS terminal.

<p>Proper SEC Code: <strong>POS</strong></p>

### Internet Transactions

##### Physical Enrollment

- Customer initiates debits via the bank’s web site or individual bill payments at a merchant web site.
- Customer uses their ACH based debit card to make a purchase at a web site.

<p>Proper SEC Code: <strong>PPD</strong></p>

##### Internet Enrollment

- Customer authorizes a transfer of funds into a savings account.
- Customer initiates bill payments at a merchant web site on their debit card.
- Customer uses their ACH based debit card to make a purchase at a web site.

<p>Proper SEC Code: <strong>WEB</strong></p>

## ATM Transactions

##### Physical Enrollment

- Customer uses the card at an ATM to withdraw cash.

<p>Proper SEC Code: <strong>MTE</strong></p>

##### Internet Enrollment

- Customer uses an ATM to withdraw cash:

<p>Proper SEC Code: <strong>MTE</strong></p>
