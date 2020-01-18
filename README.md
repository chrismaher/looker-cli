# Looker CLI

This package provides a command-line interface to certain Looker 3.0 endpoints.

The interface is written with the [Cobra library](https://github.com/spf13/cobra).

## Installation

    make install

## Usage
After installing the library, the first step is to create a `$HOME/.looker/config.yaml` file with the information necessary to access Looker's APIs. The file should only contain `api_path`, `client_id`, and `client_secret`:

    api_path: https://<domain>.looker.com:19999
    client_id: <client_id>
    client_secret: <client_secret>
