# vault-demo-server

This is the source for the server that serves the interactive tutorial for [Vault](https://vaultproject.io).

The interactive tutorial uses WebSockets to communicate to this service.
When the WebSocket is opened, this service starts a _real, fully featured_
Vault instance. The API to this service is directly the CLI commands to 
execute, and the demo server actually invokes the CLI in-memory and forwards
back the response and exit code. 
