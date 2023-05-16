# Temphia

<div align="center">
    <img src="contrib/temphia.png" >
    
[Docs](./docs/readme.md) |
[FAQ](./docs/faq.md) |
[Architecture](./docs/arch.md) |
[Glossary](./docs/glossary.md) |
[Demo video](https://youtu.be/NGPT5krw1RI)
    

> **🚨🚨🚨Alpha software🚨🚨🚨**
</div>

Temphia is a platform for small apps, known as "plugs", which run inside a language virtual machine (VM) such as JavaScript (JS) or WebAssembly (Wasm), and communicate with each other through message passing. Apps have specific capabilities based on the resources they are assigned.

Each app also has its own private key-value storage for storing simple states, and can optionally assign a data table resource for more complex database needs. 
Additionally, app has an associated user interface (UI) written in JS, which is executed inside an iframe or sub-origin (e.g., xyzapp.mytemphia.com). [more](./docs/arch.md)


## Features
- Datatables (postgres, sqlite).
- Websocket rooms (called sockd rooms).
- Files and Folder (called cabinet).
- Pluggable language executor (js, wasm).
- Group bashed user management.
- Logging Abstraction and exploration.
- Key Value db for storing simple states.
- Repository/Store for importing packages, simple to implement own repository.
- Extendable in native golang (custom resource/modules, custom executor).
- Multi-Tenant
- Single static binary (ui embed)
