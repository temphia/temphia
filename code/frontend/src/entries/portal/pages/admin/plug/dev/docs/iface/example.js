export default {
    "name": "todo.interface",
    "methods": {
        "list_items": {
            "info": "This method lists a todo items",
            "arg": {
                "type": "null"
            },
            "return_type": {
                "type": "array",
                "values": [
                    {
                        "type": "any"
                    }
                ]
            },

            "error_types": {
                "could_not_connect_db": "Cannot connect to database"
            }
        },

        "item_add": {
            "info": "Add a new todo item",
            "arg": {
                "ref": "todo_item"
            },
            "return_type": {
                "type": "null"
            }
        }
    },
    "events": {
        "on_item_add": {
            "type": "object",
            "async": false
        },
        "on_item_update": {
            "type": "object",
            "async": false
        }
    },
    "schemas": {
        "todo_item": {
            "type": "object",
            "values": [
                {
                    "property": "value",
                    "type": "string"
                },
                {
                    "property": "done",
                    "type": "boolean"
                }
            ]
        }
    },
    "definations": {
        "init_cms_engine": {
            "sub_routes": {},
            "global_routes": {},
            "widget_hooks": {},
        }

    }
}
