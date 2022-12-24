# Glossary

## Repo
A repository, or repo, is a storage location for blueprints. It can be an external service or a local service on the same machine. Repositories are similar to stores in other software applications, such as the Google Play Store or the Firefox Add-Ons store, in that they provide a centralized location for users to access and download blueprints.

## Blueprint
A blueprint, or bprint, is a package similar to a JAR file in Java or an APK file in Android, or an ADDON.ZIG file. However, blueprints are more generic because the type field define the specific nature of the blueprint. This might contain code files, data files, extension files, schema files for data tables etc, which can be instanced to plug/agent or data tables.

## Plug
Plug is a namespace for group of agents where agent represent the code that executes in specific context. All the agents in a plug have access to a key value database for quickly storing states and meta data.

## Agent
An agent is a piece of code that performs some action or task. It may have a user interface associated with it, typically implemented as a JavaScript entry file that is executed by an executor. The agent processes requests from the user interface by making method calls to the server-side code using a RPC like mechanism. Agents can also communicate with other agents that are connected through a link by sending similar RPC messages.

## Link
A link is a connection between two agents that allows them to communicate with another through RPCs like system using JSON messages. An agent can have multiple links, which can extend its functionality or allow it to send messages about events. For example, an e-commerce system might use links to connect to different payment processors dynamically. This allows the system to communicate with the payment processors and access their functionality as needed.

## Extension
Extension are way to extend the ui of the agent using js file. It allows to extend some agent instanced from one bprint to using js file using another bprint. This allows to modify (mod) a agent using third party (one which is not the original author of app/agent) similar to greasemonkey/userscript but without agent user fiddling with browser extension

## Resource
Resources are entities that can be associated with an agent, granting it specific capabilities. For example, a room resource may allow the agent to post messages to a specific websocket room, while a folder resource allows the agent to perform file operations within a particular folder.  

## Target App
A target app is an agent that is associated with a specific entity, such as a user, user group, or data table. It is typically executed within a browser iframe and may be used to perform a variety of tasks, such as displaying information or allowing users to interact with data. Target apps are typically associated with a particular agent and are executed based on the needs of the entity they are associated with.

## Target Hook
A target hook is an agent that is triggered in response to the occurrence of a specific event, such as a user modifying a database or logging in. It is executed on the server side and may be used to perform a variety of tasks, such as updating a database or sending a notification. Target hooks are associated with a particular agent and are executed based on the response to an event.

## Module
A module is a type of resource that grants a specific capability to an agent. It is implemented in native Go code and may be used to perform a variety of tasks, such as accessing external services or performing complex calculations. Modules are typically associated with a particular agent and are used to extend its functionality or capabilities.

## Invoker
An invoker is a context in which an agent is executed. It is responsible for invoking or executing the agent, and the agent can call certain methods on the invoker to perform context-specific tasks. For example, a target hook that is triggered when a data change event occurs may be invoked by a DataInvoker, and the agent can use that context to perform database operations. In this way, the invoker provides the agent with the necessary context and resources to perform its tasks.

## Domain Adapter
A domain adapter is a native Go method that handles HTTP requests for a specific domain name. It may be responsible for executing target hooks that are associated with a target domain.

## User Group
A user group is a collection of users that share a common set of access to data tables, authentication methods, and policies. User groups are often used to manage access to resources and to enforce security policies within an organization, and each user's membership in a particular user group determines the scope of their access and the actions they are allowed to perform. It's important to note that a user can only belong to one user group at a time, but a user can have multiple roles and permissions, which may grant them different levels of access to resources and allow them to perform different actions. Roles and permissions are often used in conjunction with user groups to provide more granular control over access and actions within a system.