var __dirname = ""; var module = {};
/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	// The require scope
/******/ 	var __nccwpck_require__ = {};
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__nccwpck_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__nccwpck_require__.o(definition, key) && !__nccwpck_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__nccwpck_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__nccwpck_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/compat */
/******/ 	
/******/ 	if (typeof __nccwpck_require__ !== 'undefined') __nccwpck_require__.ab = __dirname + "/";
/******/ 	
/************************************************************************/
var __webpack_exports__ = {};
// ESM COMPAT FLAG
__nccwpck_require__.r(__webpack_exports__);

// EXPORTS
__nccwpck_require__.d(__webpack_exports__, {
  "CabFolder": () => (/* reexport */ CabFolder),
  "Request": () => (/* reexport */ Request),
  "Response": () => (/* reexport */ Response),
  "SockdRoom": () => (/* reexport */ SockdRoom),
  "core": () => (/* reexport */ core),
  "invoker": () => (/* reexport */ invoker),
  "plugkv": () => (/* reexport */ plugkv),
  "self": () => (/* reexport */ self_self),
  "utils": () => (/* reexport */ utils)
});

;// CONCATENATED MODULE: ./lib/core.ts
var _buffer = [];
var core = {
    log: function (message) { return _log(message); },
    log_lazy: function (message) { return _buffer.push(message); },
    lazy_log_send: function () {
        _log_lazy(_buffer);
        _buffer.length = 0;
    },
    sleep: function (t) { return _sleep(t); },
    self_file: function (file) { return _get_self_file_as_str(file); }
};

;// CONCATENATED MODULE: ./lib/plugkv.ts
var PlugKV = /** @class */ (function () {
    function PlugKV(txid) {
        var _this = this;
        this.get_ticket = function (opts) {
            return _pkv_ticket(opts);
        };
        this.quick_get = function (key) {
            var _a = _pkv_get(_this.txid, key), val = _a[0], err = _a[1];
            if (err) {
                return [null, err];
            }
            return [val.value, null];
        };
        this.set = function (key, value, opts) {
            return _pkv_set(_this.txid, key, value, opts);
        };
        this.update = function (key, value, opts) {
            return _pkv_update(_this.txid, key, value, opts);
        };
        this.get = function (key) {
            var _a = _pkv_get(_this.txid, key), val = _a[0], err = _a[1];
            if (err) {
                return [null, err];
            }
            return [val, null];
        };
        this.query = function (opts) {
            return _pkv_query(_this.txid, opts);
        };
        this["delete"] = function (key) {
            return _pkv_del(_this.txid, key);
        };
        this.batch_delete = function (keys) {
            return _pkv_batch_del(_this.txid, keys);
        };
        this.new_txn = function () {
            var _a = _pkv_new_txn(), newtxid = _a[0], err = _a[1];
            if (err) {
                return [null, err];
            }
            return [new PlugKV(newtxid), null];
        };
        this.rollback = function () {
            return _pkv_rollback(_this.txid);
        };
        this.commit = function () {
            return _pkv_commit(_this.txid);
        };
        this.txid = txid;
    }
    return PlugKV;
}());

var plugkv = new PlugKV(0);

;// CONCATENATED MODULE: ./lib/http.ts
var Response = /** @class */ (function () {
    function Response(status, header, body) {
        var _this = this;
        this.ok = function () {
            return _this.status >= 200 && _this.status < 300;
        };
        this.get_header = function (key) {
            return _this.headers[key];
        };
        this.json_body = function () {
            return JSON.parse(_this.body);
        };
        this.status = status;
        this.headers = header;
        this.body = body;
    }
    return Response;
}());

var Request = /** @class */ (function () {
    function Request(url) {
        var _this = this;
        this.set_header = function (key, value) {
            _this._headers[key] = value;
        };
        this.set_body = function (body) {
            _this._body = body;
        };
        this.set_json_body = function (value) {
            _this._body = JSON.stringify(value);
        };
        this.get = function () {
            var _a = _http1("GET", _this._url, _this._headers, ""), status = _a[0], header = _a[1], body = _a[2];
            return new Response(status, header, body);
        };
        this.post = function () {
            var _a = _http1("POST", _this._url, _this._headers, _this._body), status = _a[0], header = _a[1], body = _a[2];
            return new Response(status, header, body);
        };
        this.put = function () {
            var _a = _http1("PUT", _this._url, _this._headers, _this._body), status = _a[0], header = _a[1], body = _a[2];
            return new Response(status, header, body);
        };
        this.patch = function () {
            var _a = _http1("PATCH", _this._url, _this._headers, _this._body), status = _a[0], header = _a[1], body = _a[2];
            return new Response(status, header, body);
        };
        this["delete"] = function () {
            var _a = _http1("DELETE", _this._url, _this._headers, _this._body), status = _a[0], header = _a[1], body = _a[2];
            return new Response(status, header, body);
        };
        this._url = url;
        this._headers = {};
    }
    return Request;
}());


;// CONCATENATED MODULE: ./lib/cabinet.ts
var CabFolder = /** @class */ (function () {
    function CabFolder(folder) {
        var _this = this;
        this.add_file = function (file) {
            return _cab_add_file(_this._folder, file);
        };
        this.list_folder = function (file) {
            return _cab_list_folder(_this._folder, file);
        };
        this.get_file = function (file) {
            return _cab_get_file(_this._folder, file);
        };
        this.get_file_str = function (file) {
            return _cab_get_file_str(_this._folder, file);
        };
        this.del_file = function (file) {
            return _cab_del_file(_this._folder, file);
        };
        this.gen_ticket = function (opts) {
            return _cab_generate_ticket(_this._folder, opts);
        };
        this._folder = folder;
    }
    return CabFolder;
}());


;// CONCATENATED MODULE: ./lib/sockd.ts
var SockdRoom = /** @class */ (function () {
    function SockdRoom(room) {
        var _this = this;
        this.send_direct = function (connIds, value) {
            return _sd_send_direct(_this._room, connIds, value);
        };
        this.send_direct_batch = function (connIds, value) {
            return _sd_send_direct_batch(_this._room, connIds, value);
        };
        this.send_broadcast = function (value, ignores) {
            return _sd_send_broadcast(_this._room, value, ignores ? ignores : []);
        };
        this.send_tagged = function (tags, value, ignore) {
            return _sd_send_tagged(_this._room, tags, value, ignore ? ignore : []);
        };
        this.ticket = function (opts) {
            return _sd_ticket(_this._room, opts);
        };
        this._room = room;
    }
    return SockdRoom;
}());


;// CONCATENATED MODULE: ./lib/utils.ts
var no_found = "upper: no more rows in this result set";
var already_exists = "duplicate key value violates";
var utils = {
    is_db_not_found: function (err) {
        return err.indexOf(no_found) !== -1;
    },
    is_db_already_exists: function (err) {
        return err.indexOf(already_exists) !== -1;
    },
    ok_response: function (data) { return ({ payload: { ok: true, data: data } }); },
    err_response: function (message) { return ({
        payload: { ok: false, message: message }
    }); },
    ab2str: function (buf) {
        return String.fromCharCode.apply(null, new Uint16Array(buf));
    },
    str2ab: function (str) {
        var buf = new ArrayBuffer(str.length * 2);
        var bufView = new Uint16Array(buf);
        for (var i = 0, strLen = str.length; i < strLen; i++) {
            bufView[i] = str.charCodeAt(i);
        }
        return buf;
    },
    is_arraybuffer: function (value) {
        return (typeof ArrayBuffer === "function" &&
            (value instanceof ArrayBuffer ||
                toString.call(value) === "[object ArrayBuffer]"));
    },
    generate_str_id: function () { return Math.random().toString(36).slice(2); }
};

;// CONCATENATED MODULE: ./lib/self.ts
var self_self = {
    list_resource: function () { return _self_list_resource(); },
    get_resource: function (name) { return _self_get_resource(name); },
    inlinks: function () { return _self_inlinks(); },
    outlinks: function () { return _self_outlinks(); },
    new_module: function (name, data) { return _self_new_module(name, data); },
    module_execute: function (mid, method, data) {
        return _self_module_execute(mid, method, data);
    },
    link_execute: function (name, method, path, data, async, detached) { return _self_link_execute(name, method, path, data, async, detached); },
    fork_execute: function (method, data) { return _self_fork_execute(method, data); }
};

;// CONCATENATED MODULE: ./lib/invoker.ts
var invoker = {
    name: function () { return _invoker_name(); },
    exec_method: function (method, path, data) {
        return _invoker_exec_method(method, path, data);
    },
    context_user: function () { return _invoker_context_user(); },
    context_user_info: function () { return _invoker_context_user_info(); },
    context_user_message: function () { return _invoker_context_user_message(); }
};

;// CONCATENATED MODULE: ./lib/index.ts










module.exports = __webpack_exports__;
/******/ })()
;Object.assign(globalThis, module.exports);