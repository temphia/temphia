import ActionRouter from action_router

async def hello_world(param):
    return "HELLO_WORLD"

router = ActionRouter()

router.register_action(hello_world)

router.run()