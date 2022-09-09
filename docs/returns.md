---
layout: page
title: Return files
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Return files

ACH files and their EntryDetail records may be returned for a variety of reasons defined by "return codes". Either the originating (ODFI) or receiving (RDFI) financial institution can issue a return. Some of the most common reasons include: insufficient funds (`R01`) and account information not found (`R03`).

Returns are identified by an [Addenda99](https://pkg.go.dev/github.com/moov-io/ach?tab=doc#Addenda99) record on the EntryDetail with a [ReturnCode](https://pkg.go.dev/github.com/moov-io/ach?tab=doc#ReturnCode) that can be processed.

### Processing

An Originator must have a process set up with their ODFI to handle returns. This may involve manual verification of actions to take, retrying entries (only in the case of an `R01` or `R09`), or revocation of future transfers for the customer (`R07` - Authorization revoked).

Fees may be applied by the ODFI/RDFI as a result of returned files. Written authorization may be requested as a result of returned files as well. Such authorization requests may be prompted if the description, amount, or effective date is different than what the customer had authorized.

The `R05` return code is used when an unauthorized corporate debit (standard entry class code of `CCD` or `CTX`) was used

### Creation

When creating a return entry add an [Addenda99](https://pkg.go.dev/github.com/moov-io/ach?tab=doc#Addenda99) record onto the EntryDetail with the appropriate return code. You'll need to copy the original `TraceNumber` and populate a new value when returning the file.

```go
addenda99 := NewAddenda99()
addenda99.ReturnCode = "R07"
addenda99.OriginalTrace = "99912340000015"
addenda99.AddendaInformation = "Authorization Revoked"
addenda99.OriginalDFI = "9101298"

// entry.Addenda99 = addenda99
```

### Return codes

| Code | Reason | Description |
|----|-----|------|
| `R01` | Insufficient Funds | Available balance is not sufficient to cover the dollar value of the debit entry |
| `R02` | Account Closed | Previously active account has been closed by customer or RDFI |
| `R03` | No Account/Unable to Locate Account | Account number structure is valid and passes editing process, but does not correspond to individual or is not an open account |
| `R04` | Invalid Account Number | Account number structure not valid; entry may fail check digit validation or may contain an incorrect number of digits. |
| `R05` | Improper Debit to Consumer Account | A CCD, CTX, or CBR debit entry was transmitted to a Consumer Account of the Receiver and was not authorized by the Receiver |
| `R06` | Returned per ODFI's Request | ODFI has requested RDFI to return the ACH entry (optional to RDFI - ODFI indemnifies RDFI) |
| `R07` | Authorization Revoked by Customer | Consumer, who previously authorized ACH payment, has revoked authorization from Originator (must be returned no later than 60 days from settlement date and customer must sign affidavit) |
| `R08` | Payment Stopped | Receiver of a recurring debit transaction has stopped payment to a specific ACH debit. RDFI should verify the Receiver's intent when a request for stop payment is made to insure this is not intended to be a revocation of authorization |
| `R09` | Uncollected Funds | Sufficient book or ledger balance exists to satisfy dollar value of the transaction, but the dollar value of transaction is in process of collection (i.e., uncollected checks) or cash reserve balance below dollar value of the debit entry. |
| `R10` | Customer Advises Originator is Not Known to Receiver and/or Originator is Not Authorized by Receiver to Debit Receiver’s Account | The receiver does not know the Originator’s identity and/or has not authorized the Originator to debit. Alternatively, for ARC, BOC, and POP entries, the signature is not authentic or authorized. |
| `R11` | Customer Advises Entry Not in Accordance with the Terms of the Authorization | The Originator and Receiver have a relationship, and an authorization to debit exists, but there is an error or defect in the payment such that the entry does not conform to the terms of the authorization. The Originator may correct the error and submit a new entry within 60 days of the return entry settlement date without the need for re-authorization by the Receiver. |
| `R12` | Branch Sold to Another DFI | Financial institution receives entry destined for an account at a branch that has been sold to another financial institution. |
| `R13` | RDFI not qualified to participate | Financial institution does not receive commercial ACH entries |
| `R14` | Representative payee deceased or unable to continue in that capacity | The representative payee authorized to accept entries on behalf of a beneficiary is either deceased or unable to continue in that capacity |
| `R15` | Beneficiary or bank account holder | (Other than representative payee) deceased* - (1) the beneficiary entitled to payments is deceased or (2) the bank account holder other than a representative payee is deceased |
| `R16` | Bank account frozen | Funds in bank account are unavailable due to action by RDFI or legal order |
| `R17` | File record edit criteria | Fields rejected by RDFI processing (identified in return addenda) |
| `R18` | Improper effective entry date | Entries have been presented prior to the first available processing window for the effective date. |
| `R19` | Amount field error | Improper formatting of the amount field |
| `R20` | Non-payment bank account | Entry destined for non-payment bank account defined by reg. |
| `R21` | Invalid company ID number | The company ID information not valid (normally CIE entries) |
| `R22` | Invalid individual ID number | Individual id used by receiver is incorrect (CIE entries) |
| `R23` | Credit entry refused by receiver | Receiver returned entry because minimum or exact amount not remitted, bank account is subject to litigation, or payment represents an overpayment, originator is not known to receiver or receiver has not authorized this credit entry to this bank account |
| `R24` | Duplicate entry | RDFI has received a duplicate entry |
| `R25` | Addenda error | Improper formatting of the addenda record information |
| `R26` | Mandatory field error | Improper information in one of the mandatory fields |
| `R27` | Trace number error | Original entry trace number is not valid for return entry; or addenda trace numbers do not correspond with entry detail record |
| `R28` | Transit routing number check digit error | Check digit for the transit routing number is incorrect |
| `R29` | Corporate customer advises not authorized | RDFI has bee notified by corporate receiver that debit entry of originator is not authorized |
| `R30` | RDFI not participant in check truncation program | Financial institution not participating in automated check safekeeping application |
| `R31` | Permissible return entry (CCD and CTX only) | RDFI has been notified by the ODFI that it agrees to accept a CCD or CTX return entry |
| `R32` | RDFI non-settlement | RDFI is not able to settle the entry |
| `R33` | Return of XCK entry | RDFI determines at its sole discretion to return an XCK entry; an XCK return entry may be initiated by midnight of the sixtieth day following the settlement date if the XCK entry |
| `R34` | Limited participation RDFI | RDFI participation has been limited by a federal or state supervisor |
| `R35` | Return of improper debit entry | ACH debit not permitted for use with the CIE standard entry class code (except for reversals) |
| `R37` | Source Document Presented for Payment (Adjustment Entry) | The source document to which an ARC, BOC or POP entry relates has been presented for payment. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date |
| `R38` | Stop Payment on Source Document (Adjustment Entry) | A stop payment has been placed on the source document to which the ARC or BOC entry relates. RDFI must return no later than 60 days following Settlement Date. No Written Statement is required as the original stop payment form covers the return |
| `R39` | Improper Source Document | The RDFI has determined the source document used for the ARC, BOC or POP entry to its Receiver's account is improper. |

#### Used for ENR entries and are initiated by a Federal Government Agency

| Code | Reason | Description |
|----|-----|------|
| `R40` | Return of ENR Entry by Federal Government Agency (ENR Only) | This return reason code may only be used to return ENR entries and is at the federal Government Agency's Sole discretion |
| `R41` | Invalid Transaction Code (ENR only) | Either the Transaction Code included in Field 3 of the Addenda Record does not conform to the ACH Record Format Specifications contained in Appendix Three (ACH Record Format Specifications) or it is not appropriate with regard to an Automated Enrollment Entry. |
| `R42` | Routing Number/Check Digit Error (ENR Only) | The Routing Number and the Check Digit included in Field 3 of the Addenda Record is either not a valid number or it does not conform to the Modulus 10 formula. |
| `R43` | Invalid DFI Account Number (ENR Only) | The Receiver's account number included in Field 3 of the Addenda Record must include at least one alphameric character. |
| `R44` | Invalid Individual ID Number/Identification Number (ENR only) | The Individual ID Number/Identification Number provided in Field 3 of the Addenda Record does not match a corresponding ID number in the Federal Government Agency's records. |
| `R45` | Invalid Individual Name/Company Name (ENR only) | The name of the consumer or company provided in Field 3 of the Addenda Record either does not match a corresponding name in the Federal Government Agency's records or fails to include at least one alphameric character. |
| `R46` | Invalid Representative Payee Indicator (ENR Only) | The Representative Payee Indicator Code included in Field 3 of the Addenda Record has been omitted or it is not consistent with the Federal Government Agency's records. |
| `R47` | Duplicate Enrollment (ENR Only) | The Entry is a duplicate of an Automated Enrollment Entry previously initiated by a DFI. |

#### Used for RCK entries only and are initiated by an RDFI

| Code | Reason | Description |
|----|-----|------|
| `R50` | State Law Affecting RCK Acceptance | RDFI is located in a state that has not adopted Revised Article 4 of the UCC or the RDFI is located in a state that requires all canceled checks to be returned within the periodic statement |
| `R51` | Item Related to RCK Entry is Ineligible or RCK Entry is Improper | The item to which the RCK entry relates was not eligible, Originator did not provide notice of the RCK policy, signature on the item was not genuine, the item has been altered or amount of the entry was not accurately obtained from the item. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date |
| `R52` | Stop Payment on Item (Adjustment Entry) | A stop payment has been placed on the item to which the RCK entry relates. RDFI must return no later than 60 days following Settlement Date. No Written Statement is required as the original stop payment form covers the return. |
| `R53` | Item and RCK Entry Presented for Payment (Adjustment Entry) | Both the RCK entry and check have been presented for payment. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date |

#### Used by the ODFI for dishonored return entries

| Code | Reason | Description |
|----|-----|------|
| `R61` | Misrouted Return | The financial institution preparing the Return Entry (the RDFI of the original Entry) has placed the incorrect Routing Number in the Receiving DFI Identification field. |
| `R67` | Duplicate Return | The ODFI has received more than one Return for the same Entry. |
| `R68` | Untimely Return | The Return Entry has not been sent within the time frame established by these Rules. |
| `R69` | Field Error(s) | One or more of the field requirements are incorrect. |
| `R70` | Permissible Return Entry Not Accepted/Return Not Requested by ODFI | The ODFI has received a Return Entry identified by the RDFI as being returned with the permission of, or at the request of, the ODFI, but the ODFI has not agreed to accept the Entry or has not requested the return of the Entry. |

#### Used by the RDFI for contested dishonored return entries

| Code | Reason | Description |
|----|-----|------|
| `R71` | Misrouted Dishonored Return | The financial institution preparing the dishonored Return Entry (the ODFI of the original Entry) has placed the incorrect Routing Number in the Receiving DFI Identification field. |
| `R72` | Untimely Dishonored Return | The dishonored Return Entry has not been sent within the designated time frame. |
| `R73` | Timely Original Return | The RDFI is certifying that the original Return Entry was sent within the time frame designated in these Rules. |
| `R74` | Corrected Return | The RDFI is correcting a previous Return Entry that was dishonored using Return Reason Code R69 (Field Error(s)) because it contained incomplete or incorrect information. |
| `R75` | Return Not a Duplicate | The Return Entry was not a duplicate of an Entry previously returned by the RDFI. |
| `R76` | No Errors Found | The original Return Entry did not contain the errors indicated by the ODFI in the dishonored Return Entry. |

#### Used by Gateways for the return of international payments

| Code | Reason | Description |
|----|-----|------|
| `R80` | IAT Entry Coding Error | The IAT Entry is being returned due to one or more of the following conditions: Invalid DFI/Bank Branch Country Code, invalid DFI/Bank Identification Number Qualifier, invalid Foreign Exchange Indicator, invalid ISO Originating Currency Code, invalid ISO Destination Currency Code, invalid ISO Destination Country Code, invalid Transaction Type Code |
| `R81` | Non-Participant in IAT Program | The IAT Entry is being returned because the Gateway does not have an agreement with either the ODFI or the Gateway's customer to transmit Outbound IAT Entries. |
| `R82` | Invalid Foreign Receiving DFI Identification | The reference used to identify the Foreign Receiving DFI of an Outbound IAT Entry is invalid. |
| `R83` | Foreign Receiving DFI Unable to Settle | The IAT Entry is being returned due to settlement problems in the foreign payment system. |
| `R84` | Entry Not Processed by Gateway | For Outbound IAT Entries, the Entry has not been processed and is being returned at the Gateway's discretion because either (1) the processing of such Entry may expose the Gateway to excessive risk, or (2) the foreign payment system does not support the functions needed to process the transaction. |
| `R85` | Incorrectly Coded Outbound International Payment | The RDFI/Gateway has identified the Entry as an Outbound international payment and is returning the Entry because it bears an SEC Code that lacks information required by the Gateway for OFAC compliance. |
