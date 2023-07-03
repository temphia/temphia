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
        self.port = os.getenv("TEMPHIA_REMOTE_PORT")
        self.tenant_id = os.getenv("TEMPHIA_TENANT_ID")
        self.plug_id = os.getenv("TEMPHIA_PLUG_ID")
        self.agent_id = os.getenv("TEMPHIA_AGENT_ID")
        self.token = os.getenv("TEMPHIA_TOKEN")
        print("___INIT___")

        self.actions = {}
    def register_action(name, action):
        self.actions[name] = action
    def run(self):
        asyncio.run(self.__run__())

    async def __run__(self):
        self.reader, self.writer = await asyncio.open_connection('localhost', int(self.port))

        auth_data = json.dumps({
            "type": "preform_auth",
            "token": self.token,
        })

        self.writer.write(auth_data.encode())
        self.writer.drain()


        tasks = []
        while True:
            data = await self.reader.readline()
            if not data:
                break
            message = data.decode().strip()            
            tasks.append(self.__handle_message__(message))
        await asyncio.gather(*tasks)  # Await all response tasks concurrently
        self.writer.close()

    async def __handle_message__(self):
        try:
            data = json.loads(message)
            action_type = data.get("action_type")
            if action_type in actions:
                response = await actions[action_type](data)
                response_message = json.dumps(response)
                self.writer.write(response_message.encode())
                await self.writer.drain()
            else:
                print("Unknown action type:", action_type)
        except json.JSONDecodeError as e:
            print("Failed to decode JSON message:", e)
        except Exception as e:
            print("Error processing message:", e)
            pass


router = ActionRouter()
router.run()

