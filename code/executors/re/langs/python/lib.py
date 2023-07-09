import os
import asyncio
import json

DEFAULT_BINDINGS = None

class ActionContext:
    def __init__(self, rawdata, router):
        self.__id = ""
        self.__router = router
        self.__rawdata = rawdata
    def data(self):
        return json.loads(self.__rawdata)
    async def invoker_name(self):
        pass
    async def exec_method(self, name, data):
        pass
    async def user_context(self):
        pass
    async def user_info(self):
        pass
    async def message_user(self, msg):
        pass

class ActionRouter:
    def __init__(self):
        self.port = os.getenv("TEMPHIA_REMOTE_PORT") or "1234"
        self.tenant_id = os.getenv("TEMPHIA_TENANT_ID")
        self.plug_id = os.getenv("TEMPHIA_PLUG_ID")
        self.agent_id = os.getenv("TEMPHIA_AGENT_ID")
        self.token = os.getenv("TEMPHIA_TOKEN")
        self.bind_conns = {}

        print(self.port, self.tenant_id, self.plug_id, self.agent_id, self.token)

        self.actions = {}
    def register_action(self, name, action):
        self.actions[name] = action
    def run(self):
        asyncio.run(self.__run__())


    async def _run_client_conn(self, id):
        self.reader, self.writer = await asyncio.open_connection('localhost', int(self.port))
        auth_data = json.dumps({
            "type": "bind_auth",
            "token": self.token,
        })

        self.writer.write(auth_data.encode())
        await self.writer.drain()
        
        self.bind_conns[id] = (reader, writer)



    async def __run__(self):
        self.reader, self.writer = await asyncio.open_connection('localhost', int(self.port))

        auth_data = json.dumps({
            "type": "control_auth",
            "token": self.token,
        })

        self.writer.write(auth_data.encode())
        await self.writer.drain()


        tasks = []
        while True:
            print("@PY", "waiting")

            data = await self.reader.readline()
            if not data:
                break
            

            message = data.decode().strip()            
            
            print("@PY", message)

            self.__handle_message__(message)

            tasks.append(self.__handle_message__(message))

        await asyncio.gather(*tasks)  # Await all response tasks concurrently
        self.writer.close()

    async def __handle_message__(self, message):
        print("@handle_message", message)

        try:
            data = json.loads(message)
            name = data.get("name")
            if name in actions:
                response = await actions[name](data)
                response_message = json.dumps(response)
                self.writer.write(response_message.encode())
                await self.writer.drain()
            else:
                print("Unknown action type:", name)
        except json.JSONDecodeError as e:
            print("Failed to decode JSON message:", e)
        except Exception as e:
            print("Error processing message:", e)
            pass

async def hello_world(param):
    return "PONG"



if __name__ == "__main__":
    router = ActionRouter()
    router.register_action("ping", hello_world)
    router.run()

