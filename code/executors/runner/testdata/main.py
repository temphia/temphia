import ActionRouter from temphia_python_executor

async def hello_world(param):
    return "HELLO_WORLD"

router = ActionRouter()

router.register_action(hello_world)

router.run()