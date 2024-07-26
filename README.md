# Golang web app starter template

This project is made to accelerate a web app creation by
- creating a correct structure
- handling the server
- html templates layout
- handling errors

## Usage

## Architecture

## error handling

<!-- middleware refers to a function that wraps an HTTP handler to add additional behavior before
or after the handler processes an HTTP request


The HandleError function will set the appropriate HTTP status code and render an error page with the provided message.
call HandleError directly within your HTTP handlers when you encounter an error.

The WithErrorHandling middleware uses HandleError within a defer statement to handle any panics that occur during the request processing. It logs the panic and sends an appropriate error response using HandleError. -->

## testing

## Attribution

This favicon was generated using the following graphics from Twitter Twemoji:

- Graphics Title: 2620.svg
- Graphics Author: Copyright 2020 Twitter, Inc and other contributors (https://github.com/twitter/twemoji)
- Graphics Source: https://github.com/twitter/twemoji/blob/master/assets/svg/2620.svg
- Graphics License: CC-BY 4.0 (https://creativecommons.org/licenses/by/4.0/)
