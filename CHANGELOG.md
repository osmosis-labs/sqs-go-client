<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState
given same genesisState and txList.
Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## v0.0.12

- Update osmoutils-go to v0.0.16.

## v0.0.11

- Update api key header to x-api-key.

## v0.0.10

- Update osmoutils-go and add WithOutGivenInCustom option to GetQuote.

## v0.0.9

- Remove regex validation from RouterQuoteOptions tokenIn and tokenOut.

## v0.0.8

- Export SQSQuoteResponse, SQSPool, SQSRoute, SQSPriceInfo, Coin, Pool, Route, PriceInfo.

## v0.0.7

- Add WithAppendBaseFee option to GetQuote and extend SQSQuoteResponse with more data.

##  v0.0.5 - 0.0.6

- rename GetRoute to GetQuote

## v0.0.4

- Add SQSMock to support testing.

## v0.0.3

- Add WithCustomURL option to Initialize.

## v0.0.2

- Improve router/quote WithOutGivenIn and WithInGivenOut options.

## v0.0.1

-  Initial release.
