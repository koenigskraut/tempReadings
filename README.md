# tempReadings

A simple but practical all-in-one-binary app that handles UDP temperature readings from three remote sensors (DS18B20), writes them into the database and renders a simple web page.

HTML and icons for the web page are embedded into a binary for portability with `//go:embed` directive.

All settings are handled as env vars:
* `DOMAIN`
* `CERT_FILE`/`KEY_FILE`
* `DB_USER`/`DB_PASS`, `DB_HOST`/`DB_PORT`, `DB_NAME`
* `UDP_PORT`
