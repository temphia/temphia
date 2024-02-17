# Temphia


**THIS PROJECT AGAIN GOING THROUGH REWRITE JOIN DISCORD IF YOU WANT TO BE INVOLVED :D**

[Discord Link](https://discord.gg/HUpaNA35XR)



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
- 💾 Datatables (postgres, sqlite).
- 🧦 Websocket rooms (called sockd rooms).
- 📂 Files and Folder (called cabinet).
- 🧩 Pluggable language executor (js, wasm).
- 👥 Group bashed user management.
- 🔍 Logging Abstraction and exploration.
- 🗄️ Key Value db for storing simple states.
- 📦 Repository/Store for importing packages, simple to implement own repository.
- 🐹 Extendable in native golang (custom resource/modules, custom executor).
- 🏢 Multi-Tenant
- 🗃️ Single static binary (ui embed)

## Why
- Batteries included means it's much faster to develop apps, as well as developed apps will be leaner because you are developing on top of primitives provided by the runtime.
- If everyone makes apps targeting a common runtime, then system management burden becomes much less. System operations like data backup and restore are standardized across all apps.
- Apps developed this way are much more composable, apps can emit message for special events and other apps can react to that messages.
- As an app developer, you're relieved from the responsibility of hosting your app and dealing with associated problems. Instead, users can import and install apps from a repository of their choice on hardware they own or choose vendor they like.
