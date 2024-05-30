# CHANGELOG

## v1.0.1

*May 30, 2024*

### FEATURES

- Add an alert for when 10 oracles are missing within a 30-block window.
- Fix the issue disconnecting the websocket.
    - Implement logic to automatically reconnect the WebSocket if it detects that it has been disconnected.
