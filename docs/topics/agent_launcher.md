# Agent Launcher

launcher is responsible for loading and running apps in different modes, each of which has its own trade-offs and considerations. The sub-origin mode provides a high level of isolation and security, but requires a way to manage DNS records or handle wildcard DNS requests. The iframe mode provides a moderate level of isolation and does not require any DNS changes, but may be limited in the actions it can perform. The native DOM mode provides the least isolation, but allows the app to have full access to the DOM and perform any actions it desires. It is important to choose the appropriate mode based on the needs of the app and the level of trust in the app's code.

## Sub-origin mode
Sub-origin mode, the launcher runs the app in a separate sub-origin, such as http://adminconsole.example.com, http://app1.example.com, or http://app2.example.com. This allows the app to be isolated from the main origin and prevents it from accessing sensitive data, such as cookies or local storage. However, this requires a way to dynamically modify DNS records or handle wildcard DNS requests to support the different sub-origins.

## Iframe mode
In the iframe mode, the launcher runs the app inside an <iframe> element, which provides a level of isolation from the main origin. This allows the app to operate within its own origin, without requiring any DNS changes. However, the app is still sandboxed within the iframe and may be limited in the actions it can perform.

## Native DOM mode
In the native DOM mode, the launcher runs the app directly in the native DOM without any sandboxing. This allows the app to have full access to the DOM and perform any actions it desires, but it also means that the app may potentially be able to perform malicious actions if it is not trusted. This mode is typically used when the app is running in a separate origin from the admin console or when it is being used as a headless CMS and needs to run dynamically in a whitelisted context.