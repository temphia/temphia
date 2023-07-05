from action_router import ActionRouter 


async def hello_world(param):
    return "HELLO_WORLD"

router = ActionRouter()

router.register_action("hello_world", hello_world)

router.run()