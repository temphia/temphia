var __dirname = ''; var module = {}; module['exports']={};/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ 958:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

/* unused reexport */ __nccwpck_require__(370);

/***/ }),

/***/ 431:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);
var settle = __nccwpck_require__(60);
var buildFullPath = __nccwpck_require__(644);
var buildURL = __nccwpck_require__(358);
var http = __nccwpck_require__(685);
var https = __nccwpck_require__(687);
var httpFollow = (__nccwpck_require__(774).http);
var httpsFollow = (__nccwpck_require__(774).https);
var url = __nccwpck_require__(310);
var zlib = __nccwpck_require__(796);
var pkg = __nccwpck_require__(837);
var createError = __nccwpck_require__(287);
var enhanceError = __nccwpck_require__(234);

var isHttps = /https:?/;

/**
 *
 * @param {http.ClientRequestArgs} options
 * @param {AxiosProxyConfig} proxy
 * @param {string} location
 */
function setProxy(options, proxy, location) {
  options.hostname = proxy.host;
  options.host = proxy.host;
  options.port = proxy.port;
  options.path = location;

  // Basic proxy authorization
  if (proxy.auth) {
    var base64 = Buffer.from(proxy.auth.username + ':' + proxy.auth.password, 'utf8').toString('base64');
    options.headers['Proxy-Authorization'] = 'Basic ' + base64;
  }

  // If a proxy is used, any redirects must also pass through the proxy
  options.beforeRedirect = function beforeRedirect(redirection) {
    redirection.headers.host = redirection.host;
    setProxy(redirection, proxy, redirection.href);
  };
}

/*eslint consistent-return:0*/
module.exports = function httpAdapter(config) {
  return new Promise(function dispatchHttpRequest(resolvePromise, rejectPromise) {
    var resolve = function resolve(value) {
      resolvePromise(value);
    };
    var reject = function reject(value) {
      rejectPromise(value);
    };
    var data = config.data;
    var headers = config.headers;

    // Set User-Agent (required by some servers)
    // See https://github.com/axios/axios/issues/69
    if ('User-Agent' in headers || 'user-agent' in headers) {
      // User-Agent is specified; handle case where no UA header is desired
      if (!headers['User-Agent'] && !headers['user-agent']) {
        delete headers['User-Agent'];
        delete headers['user-agent'];
      }
      // Otherwise, use specified value
    } else {
      // Only set header if it hasn't been set in config
      headers['User-Agent'] = 'axios/' + pkg.version;
    }

    if (data && !utils.isStream(data)) {
      if (Buffer.isBuffer(data)) {
        // Nothing to do...
      } else if (utils.isArrayBuffer(data)) {
        data = Buffer.from(new Uint8Array(data));
      } else if (utils.isString(data)) {
        data = Buffer.from(data, 'utf-8');
      } else {
        return reject(createError(
          'Data after transformation must be a string, an ArrayBuffer, a Buffer, or a Stream',
          config
        ));
      }

      // Add Content-Length header if data exists
      headers['Content-Length'] = data.length;
    }

    // HTTP basic authentication
    var auth = undefined;
    if (config.auth) {
      var username = config.auth.username || '';
      var password = config.auth.password || '';
      auth = username + ':' + password;
    }

    // Parse url
    var fullPath = buildFullPath(config.baseURL, config.url);
    var parsed = url.parse(fullPath);
    var protocol = parsed.protocol || 'http:';

    if (!auth && parsed.auth) {
      var urlAuth = parsed.auth.split(':');
      var urlUsername = urlAuth[0] || '';
      var urlPassword = urlAuth[1] || '';
      auth = urlUsername + ':' + urlPassword;
    }

    if (auth) {
      delete headers.Authorization;
    }

    var isHttpsRequest = isHttps.test(protocol);
    var agent = isHttpsRequest ? config.httpsAgent : config.httpAgent;

    var options = {
      path: buildURL(parsed.path, config.params, config.paramsSerializer).replace(/^\?/, ''),
      method: config.method.toUpperCase(),
      headers: headers,
      agent: agent,
      agents: { http: config.httpAgent, https: config.httpsAgent },
      auth: auth
    };

    if (config.socketPath) {
      options.socketPath = config.socketPath;
    } else {
      options.hostname = parsed.hostname;
      options.port = parsed.port;
    }

    var proxy = config.proxy;
    if (!proxy && proxy !== false) {
      var proxyEnv = protocol.slice(0, -1) + '_proxy';
      var proxyUrl = process.env[proxyEnv] || process.env[proxyEnv.toUpperCase()];
      if (proxyUrl) {
        var parsedProxyUrl = url.parse(proxyUrl);
        var noProxyEnv = process.env.no_proxy || process.env.NO_PROXY;
        var shouldProxy = true;

        if (noProxyEnv) {
          var noProxy = noProxyEnv.split(',').map(function trim(s) {
            return s.trim();
          });

          shouldProxy = !noProxy.some(function proxyMatch(proxyElement) {
            if (!proxyElement) {
              return false;
            }
            if (proxyElement === '*') {
              return true;
            }
            if (proxyElement[0] === '.' &&
                parsed.hostname.substr(parsed.hostname.length - proxyElement.length) === proxyElement) {
              return true;
            }

            return parsed.hostname === proxyElement;
          });
        }

        if (shouldProxy) {
          proxy = {
            host: parsedProxyUrl.hostname,
            port: parsedProxyUrl.port,
            protocol: parsedProxyUrl.protocol
          };

          if (parsedProxyUrl.auth) {
            var proxyUrlAuth = parsedProxyUrl.auth.split(':');
            proxy.auth = {
              username: proxyUrlAuth[0],
              password: proxyUrlAuth[1]
            };
          }
        }
      }
    }

    if (proxy) {
      options.headers.host = parsed.hostname + (parsed.port ? ':' + parsed.port : '');
      setProxy(options, proxy, protocol + '//' + parsed.hostname + (parsed.port ? ':' + parsed.port : '') + options.path);
    }

    var transport;
    var isHttpsProxy = isHttpsRequest && (proxy ? isHttps.test(proxy.protocol) : true);
    if (config.transport) {
      transport = config.transport;
    } else if (config.maxRedirects === 0) {
      transport = isHttpsProxy ? https : http;
    } else {
      if (config.maxRedirects) {
        options.maxRedirects = config.maxRedirects;
      }
      transport = isHttpsProxy ? httpsFollow : httpFollow;
    }

    if (config.maxBodyLength > -1) {
      options.maxBodyLength = config.maxBodyLength;
    }

    // Create the request
    var req = transport.request(options, function handleResponse(res) {
      if (req.aborted) return;

      // uncompress the response body transparently if required
      var stream = res;

      // return the last request in case of redirects
      var lastRequest = res.req || req;


      // if no content, is HEAD request or decompress disabled we should not decompress
      if (res.statusCode !== 204 && lastRequest.method !== 'HEAD' && config.decompress !== false) {
        switch (res.headers['content-encoding']) {
        /*eslint default-case:0*/
        case 'gzip':
        case 'compress':
        case 'deflate':
        // add the unzipper to the body stream processing pipeline
          stream = stream.pipe(zlib.createUnzip());

          // remove the content-encoding in order to not confuse downstream operations
          delete res.headers['content-encoding'];
          break;
        }
      }

      var response = {
        status: res.statusCode,
        statusText: res.statusMessage,
        headers: res.headers,
        config: config,
        request: lastRequest
      };

      if (config.responseType === 'stream') {
        response.data = stream;
        settle(resolve, reject, response);
      } else {
        var responseBuffer = [];
        var totalResponseBytes = 0;
        stream.on('data', function handleStreamData(chunk) {
          responseBuffer.push(chunk);
          totalResponseBytes += chunk.length;

          // make sure the content length is not over the maxContentLength if specified
          if (config.maxContentLength > -1 && totalResponseBytes > config.maxContentLength) {
            stream.destroy();
            reject(createError('maxContentLength size of ' + config.maxContentLength + ' exceeded',
              config, null, lastRequest));
          }
        });

        stream.on('error', function handleStreamError(err) {
          if (req.aborted) return;
          reject(enhanceError(err, config, null, lastRequest));
        });

        stream.on('end', function handleStreamEnd() {
          var responseData = Buffer.concat(responseBuffer);
          if (config.responseType !== 'arraybuffer') {
            responseData = responseData.toString(config.responseEncoding);
            if (!config.responseEncoding || config.responseEncoding === 'utf8') {
              responseData = utils.stripBOM(responseData);
            }
          }

          response.data = responseData;
          settle(resolve, reject, response);
        });
      }
    });

    // Handle errors
    req.on('error', function handleRequestError(err) {
      if (req.aborted && err.code !== 'ERR_FR_TOO_MANY_REDIRECTS') return;
      reject(enhanceError(err, config, null, req));
    });

    // Handle request timeout
    if (config.timeout) {
      // This is forcing a int timeout to avoid problems if the `req` interface doesn't handle other types.
      var timeout = parseInt(config.timeout, 10);

      if (isNaN(timeout)) {
        reject(createError(
          'error trying to parse `config.timeout` to int',
          config,
          'ERR_PARSE_TIMEOUT',
          req
        ));

        return;
      }

      // Sometime, the response will be very slow, and does not respond, the connect event will be block by event loop system.
      // And timer callback will be fired, and abort() will be invoked before connection, then get "socket hang up" and code ECONNRESET.
      // At this time, if we have a large number of request, nodejs will hang up some socket on background. and the number will up and up.
      // And then these socket which be hang up will devoring CPU little by little.
      // ClientRequest.setTimeout will be fired on the specify milliseconds, and can make sure that abort() will be fired after connect.
      req.setTimeout(timeout, function handleRequestTimeout() {
        req.abort();
        reject(createError(
          'timeout of ' + timeout + 'ms exceeded',
          config,
          config.transitional && config.transitional.clarifyTimeoutError ? 'ETIMEDOUT' : 'ECONNABORTED',
          req
        ));
      });
    }

    if (config.cancelToken) {
      // Handle cancellation
      config.cancelToken.promise.then(function onCanceled(cancel) {
        if (req.aborted) return;

        req.abort();
        reject(cancel);
      });
    }

    // Send the request
    if (utils.isStream(data)) {
      data.on('error', function handleStreamError(err) {
        reject(enhanceError(err, config, null, req));
      }).pipe(req);
    } else {
      req.end(data);
    }
  });
};


/***/ }),

/***/ 697:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);
var settle = __nccwpck_require__(60);
var cookies = __nccwpck_require__(722);
var buildURL = __nccwpck_require__(358);
var buildFullPath = __nccwpck_require__(644);
var parseHeaders = __nccwpck_require__(695);
var isURLSameOrigin = __nccwpck_require__(913);
var createError = __nccwpck_require__(287);

module.exports = function xhrAdapter(config) {
  return new Promise(function dispatchXhrRequest(resolve, reject) {
    var requestData = config.data;
    var requestHeaders = config.headers;
    var responseType = config.responseType;

    if (utils.isFormData(requestData)) {
      delete requestHeaders['Content-Type']; // Let the browser set it
    }

    var request = new XMLHttpRequest();

    // HTTP basic authentication
    if (config.auth) {
      var username = config.auth.username || '';
      var password = config.auth.password ? unescape(encodeURIComponent(config.auth.password)) : '';
      requestHeaders.Authorization = 'Basic ' + btoa(username + ':' + password);
    }

    var fullPath = buildFullPath(config.baseURL, config.url);
    request.open(config.method.toUpperCase(), buildURL(fullPath, config.params, config.paramsSerializer), true);

    // Set the request timeout in MS
    request.timeout = config.timeout;

    function onloadend() {
      if (!request) {
        return;
      }
      // Prepare the response
      var responseHeaders = 'getAllResponseHeaders' in request ? parseHeaders(request.getAllResponseHeaders()) : null;
      var responseData = !responseType || responseType === 'text' ||  responseType === 'json' ?
        request.responseText : request.response;
      var response = {
        data: responseData,
        status: request.status,
        statusText: request.statusText,
        headers: responseHeaders,
        config: config,
        request: request
      };

      settle(resolve, reject, response);

      // Clean up request
      request = null;
    }

    if ('onloadend' in request) {
      // Use onloadend if available
      request.onloadend = onloadend;
    } else {
      // Listen for ready state to emulate onloadend
      request.onreadystatechange = function handleLoad() {
        if (!request || request.readyState !== 4) {
          return;
        }

        // The request errored out and we didn't get a response, this will be
        // handled by onerror instead
        // With one exception: request that using file: protocol, most browsers
        // will return status as 0 even though it's a successful request
        if (request.status === 0 && !(request.responseURL && request.responseURL.indexOf('file:') === 0)) {
          return;
        }
        // readystate handler is calling before onerror or ontimeout handlers,
        // so we should call onloadend on the next 'tick'
        setTimeout(onloadend);
      };
    }

    // Handle browser request cancellation (as opposed to a manual cancellation)
    request.onabort = function handleAbort() {
      if (!request) {
        return;
      }

      reject(createError('Request aborted', config, 'ECONNABORTED', request));

      // Clean up request
      request = null;
    };

    // Handle low level network errors
    request.onerror = function handleError() {
      // Real errors are hidden from us by the browser
      // onerror should only fire if it's a network error
      reject(createError('Network Error', config, null, request));

      // Clean up request
      request = null;
    };

    // Handle timeout
    request.ontimeout = function handleTimeout() {
      var timeoutErrorMessage = 'timeout of ' + config.timeout + 'ms exceeded';
      if (config.timeoutErrorMessage) {
        timeoutErrorMessage = config.timeoutErrorMessage;
      }
      reject(createError(
        timeoutErrorMessage,
        config,
        config.transitional && config.transitional.clarifyTimeoutError ? 'ETIMEDOUT' : 'ECONNABORTED',
        request));

      // Clean up request
      request = null;
    };

    // Add xsrf header
    // This is only done if running in a standard browser environment.
    // Specifically not if we're in a web worker, or react-native.
    if (utils.isStandardBrowserEnv()) {
      // Add xsrf header
      var xsrfValue = (config.withCredentials || isURLSameOrigin(fullPath)) && config.xsrfCookieName ?
        cookies.read(config.xsrfCookieName) :
        undefined;

      if (xsrfValue) {
        requestHeaders[config.xsrfHeaderName] = xsrfValue;
      }
    }

    // Add headers to the request
    if ('setRequestHeader' in request) {
      utils.forEach(requestHeaders, function setRequestHeader(val, key) {
        if (typeof requestData === 'undefined' && key.toLowerCase() === 'content-type') {
          // Remove Content-Type if data is undefined
          delete requestHeaders[key];
        } else {
          // Otherwise add header to the request
          request.setRequestHeader(key, val);
        }
      });
    }

    // Add withCredentials to request if needed
    if (!utils.isUndefined(config.withCredentials)) {
      request.withCredentials = !!config.withCredentials;
    }

    // Add responseType to request if needed
    if (responseType && responseType !== 'json') {
      request.responseType = config.responseType;
    }

    // Handle progress if needed
    if (typeof config.onDownloadProgress === 'function') {
      request.addEventListener('progress', config.onDownloadProgress);
    }

    // Not all browsers support upload events
    if (typeof config.onUploadProgress === 'function' && request.upload) {
      request.upload.addEventListener('progress', config.onUploadProgress);
    }

    if (config.cancelToken) {
      // Handle cancellation
      config.cancelToken.promise.then(function onCanceled(cancel) {
        if (!request) {
          return;
        }

        request.abort();
        reject(cancel);
        // Clean up request
        request = null;
      });
    }

    if (!requestData) {
      requestData = null;
    }

    // Send the request
    request.send(requestData);
  });
};


/***/ }),

/***/ 370:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);
var bind = __nccwpck_require__(117);
var Axios = __nccwpck_require__(304);
var mergeConfig = __nccwpck_require__(488);
var defaults = __nccwpck_require__(392);

/**
 * Create an instance of Axios
 *
 * @param {Object} defaultConfig The default config for the instance
 * @return {Axios} A new instance of Axios
 */
function createInstance(defaultConfig) {
  var context = new Axios(defaultConfig);
  var instance = bind(Axios.prototype.request, context);

  // Copy axios.prototype to instance
  utils.extend(instance, Axios.prototype, context);

  // Copy context to instance
  utils.extend(instance, context);

  return instance;
}

// Create the default instance to be exported
var axios = createInstance(defaults);

// Expose Axios class to allow class inheritance
axios.Axios = Axios;

// Factory for creating new instances
axios.create = function create(instanceConfig) {
  return createInstance(mergeConfig(axios.defaults, instanceConfig));
};

// Expose Cancel & CancelToken
axios.Cancel = __nccwpck_require__(896);
axios.CancelToken = __nccwpck_require__(897);
axios.isCancel = __nccwpck_require__(146);

// Expose all/spread
axios.all = function all(promises) {
  return Promise.all(promises);
};
axios.spread = __nccwpck_require__(666);

// Expose isAxiosError
axios.isAxiosError = __nccwpck_require__(446);

module.exports = axios;

// Allow use of default import syntax in TypeScript
module.exports["default"] = axios;


/***/ }),

/***/ 896:
/***/ ((module) => {

"use strict";


/**
 * A `Cancel` is an object that is thrown when an operation is canceled.
 *
 * @class
 * @param {string=} message The message.
 */
function Cancel(message) {
  this.message = message;
}

Cancel.prototype.toString = function toString() {
  return 'Cancel' + (this.message ? ': ' + this.message : '');
};

Cancel.prototype.__CANCEL__ = true;

module.exports = Cancel;


/***/ }),

/***/ 897:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var Cancel = __nccwpck_require__(896);

/**
 * A `CancelToken` is an object that can be used to request cancellation of an operation.
 *
 * @class
 * @param {Function} executor The executor function.
 */
function CancelToken(executor) {
  if (typeof executor !== 'function') {
    throw new TypeError('executor must be a function.');
  }

  var resolvePromise;
  this.promise = new Promise(function promiseExecutor(resolve) {
    resolvePromise = resolve;
  });

  var token = this;
  executor(function cancel(message) {
    if (token.reason) {
      // Cancellation has already been requested
      return;
    }

    token.reason = new Cancel(message);
    resolvePromise(token.reason);
  });
}

/**
 * Throws a `Cancel` if cancellation has been requested.
 */
CancelToken.prototype.throwIfRequested = function throwIfRequested() {
  if (this.reason) {
    throw this.reason;
  }
};

/**
 * Returns an object that contains a new `CancelToken` and a function that, when called,
 * cancels the `CancelToken`.
 */
CancelToken.source = function source() {
  var cancel;
  var token = new CancelToken(function executor(c) {
    cancel = c;
  });
  return {
    token: token,
    cancel: cancel
  };
};

module.exports = CancelToken;


/***/ }),

/***/ 146:
/***/ ((module) => {

"use strict";


module.exports = function isCancel(value) {
  return !!(value && value.__CANCEL__);
};


/***/ }),

/***/ 304:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);
var buildURL = __nccwpck_require__(358);
var InterceptorManager = __nccwpck_require__(314);
var dispatchRequest = __nccwpck_require__(938);
var mergeConfig = __nccwpck_require__(488);
var validator = __nccwpck_require__(359);

var validators = validator.validators;
/**
 * Create a new instance of Axios
 *
 * @param {Object} instanceConfig The default config for the instance
 */
function Axios(instanceConfig) {
  this.defaults = instanceConfig;
  this.interceptors = {
    request: new InterceptorManager(),
    response: new InterceptorManager()
  };
}

/**
 * Dispatch a request
 *
 * @param {Object} config The config specific for this request (merged with this.defaults)
 */
Axios.prototype.request = function request(config) {
  /*eslint no-param-reassign:0*/
  // Allow for axios('example/url'[, config]) a la fetch API
  if (typeof config === 'string') {
    config = arguments[1] || {};
    config.url = arguments[0];
  } else {
    config = config || {};
  }

  config = mergeConfig(this.defaults, config);

  // Set config.method
  if (config.method) {
    config.method = config.method.toLowerCase();
  } else if (this.defaults.method) {
    config.method = this.defaults.method.toLowerCase();
  } else {
    config.method = 'get';
  }

  var transitional = config.transitional;

  if (transitional !== undefined) {
    validator.assertOptions(transitional, {
      silentJSONParsing: validators.transitional(validators.boolean, '1.0.0'),
      forcedJSONParsing: validators.transitional(validators.boolean, '1.0.0'),
      clarifyTimeoutError: validators.transitional(validators.boolean, '1.0.0')
    }, false);
  }

  // filter out skipped interceptors
  var requestInterceptorChain = [];
  var synchronousRequestInterceptors = true;
  this.interceptors.request.forEach(function unshiftRequestInterceptors(interceptor) {
    if (typeof interceptor.runWhen === 'function' && interceptor.runWhen(config) === false) {
      return;
    }

    synchronousRequestInterceptors = synchronousRequestInterceptors && interceptor.synchronous;

    requestInterceptorChain.unshift(interceptor.fulfilled, interceptor.rejected);
  });

  var responseInterceptorChain = [];
  this.interceptors.response.forEach(function pushResponseInterceptors(interceptor) {
    responseInterceptorChain.push(interceptor.fulfilled, interceptor.rejected);
  });

  var promise;

  if (!synchronousRequestInterceptors) {
    var chain = [dispatchRequest, undefined];

    Array.prototype.unshift.apply(chain, requestInterceptorChain);
    chain = chain.concat(responseInterceptorChain);

    promise = Promise.resolve(config);
    while (chain.length) {
      promise = promise.then(chain.shift(), chain.shift());
    }

    return promise;
  }


  var newConfig = config;
  while (requestInterceptorChain.length) {
    var onFulfilled = requestInterceptorChain.shift();
    var onRejected = requestInterceptorChain.shift();
    try {
      newConfig = onFulfilled(newConfig);
    } catch (error) {
      onRejected(error);
      break;
    }
  }

  try {
    promise = dispatchRequest(newConfig);
  } catch (error) {
    return Promise.reject(error);
  }

  while (responseInterceptorChain.length) {
    promise = promise.then(responseInterceptorChain.shift(), responseInterceptorChain.shift());
  }

  return promise;
};

Axios.prototype.getUri = function getUri(config) {
  config = mergeConfig(this.defaults, config);
  return buildURL(config.url, config.params, config.paramsSerializer).replace(/^\?/, '');
};

// Provide aliases for supported request methods
utils.forEach(['delete', 'get', 'head', 'options'], function forEachMethodNoData(method) {
  /*eslint func-names:0*/
  Axios.prototype[method] = function(url, config) {
    return this.request(mergeConfig(config || {}, {
      method: method,
      url: url,
      data: (config || {}).data
    }));
  };
});

utils.forEach(['post', 'put', 'patch'], function forEachMethodWithData(method) {
  /*eslint func-names:0*/
  Axios.prototype[method] = function(url, data, config) {
    return this.request(mergeConfig(config || {}, {
      method: method,
      url: url,
      data: data
    }));
  };
});

module.exports = Axios;


/***/ }),

/***/ 314:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);

function InterceptorManager() {
  this.handlers = [];
}

/**
 * Add a new interceptor to the stack
 *
 * @param {Function} fulfilled The function to handle `then` for a `Promise`
 * @param {Function} rejected The function to handle `reject` for a `Promise`
 *
 * @return {Number} An ID used to remove interceptor later
 */
InterceptorManager.prototype.use = function use(fulfilled, rejected, options) {
  this.handlers.push({
    fulfilled: fulfilled,
    rejected: rejected,
    synchronous: options ? options.synchronous : false,
    runWhen: options ? options.runWhen : null
  });
  return this.handlers.length - 1;
};

/**
 * Remove an interceptor from the stack
 *
 * @param {Number} id The ID that was returned by `use`
 */
InterceptorManager.prototype.eject = function eject(id) {
  if (this.handlers[id]) {
    this.handlers[id] = null;
  }
};

/**
 * Iterate over all the registered interceptors
 *
 * This method is particularly useful for skipping over any
 * interceptors that may have become `null` calling `eject`.
 *
 * @param {Function} fn The function to call for each interceptor
 */
InterceptorManager.prototype.forEach = function forEach(fn) {
  utils.forEach(this.handlers, function forEachHandler(h) {
    if (h !== null) {
      fn(h);
    }
  });
};

module.exports = InterceptorManager;


/***/ }),

/***/ 644:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var isAbsoluteURL = __nccwpck_require__(187);
var combineURLs = __nccwpck_require__(4);

/**
 * Creates a new URL by combining the baseURL with the requestedURL,
 * only when the requestedURL is not already an absolute URL.
 * If the requestURL is absolute, this function returns the requestedURL untouched.
 *
 * @param {string} baseURL The base URL
 * @param {string} requestedURL Absolute or relative URL to combine
 * @returns {string} The combined full path
 */
module.exports = function buildFullPath(baseURL, requestedURL) {
  if (baseURL && !isAbsoluteURL(requestedURL)) {
    return combineURLs(baseURL, requestedURL);
  }
  return requestedURL;
};


/***/ }),

/***/ 287:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var enhanceError = __nccwpck_require__(234);

/**
 * Create an Error with the specified message, config, error code, request and response.
 *
 * @param {string} message The error message.
 * @param {Object} config The config.
 * @param {string} [code] The error code (for example, 'ECONNABORTED').
 * @param {Object} [request] The request.
 * @param {Object} [response] The response.
 * @returns {Error} The created error.
 */
module.exports = function createError(message, config, code, request, response) {
  var error = new Error(message);
  return enhanceError(error, config, code, request, response);
};


/***/ }),

/***/ 938:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);
var transformData = __nccwpck_require__(640);
var isCancel = __nccwpck_require__(146);
var defaults = __nccwpck_require__(392);

/**
 * Throws a `Cancel` if cancellation has been requested.
 */
function throwIfCancellationRequested(config) {
  if (config.cancelToken) {
    config.cancelToken.throwIfRequested();
  }
}

/**
 * Dispatch a request to the server using the configured adapter.
 *
 * @param {object} config The config that is to be used for the request
 * @returns {Promise} The Promise to be fulfilled
 */
module.exports = function dispatchRequest(config) {
  throwIfCancellationRequested(config);

  // Ensure headers exist
  config.headers = config.headers || {};

  // Transform request data
  config.data = transformData.call(
    config,
    config.data,
    config.headers,
    config.transformRequest
  );

  // Flatten headers
  config.headers = utils.merge(
    config.headers.common || {},
    config.headers[config.method] || {},
    config.headers
  );

  utils.forEach(
    ['delete', 'get', 'head', 'post', 'put', 'patch', 'common'],
    function cleanHeaderConfig(method) {
      delete config.headers[method];
    }
  );

  var adapter = config.adapter || defaults.adapter;

  return adapter(config).then(function onAdapterResolution(response) {
    throwIfCancellationRequested(config);

    // Transform response data
    response.data = transformData.call(
      config,
      response.data,
      response.headers,
      config.transformResponse
    );

    return response;
  }, function onAdapterRejection(reason) {
    if (!isCancel(reason)) {
      throwIfCancellationRequested(config);

      // Transform response data
      if (reason && reason.response) {
        reason.response.data = transformData.call(
          config,
          reason.response.data,
          reason.response.headers,
          config.transformResponse
        );
      }
    }

    return Promise.reject(reason);
  });
};


/***/ }),

/***/ 234:
/***/ ((module) => {

"use strict";


/**
 * Update an Error with the specified config, error code, and response.
 *
 * @param {Error} error The error to update.
 * @param {Object} config The config.
 * @param {string} [code] The error code (for example, 'ECONNABORTED').
 * @param {Object} [request] The request.
 * @param {Object} [response] The response.
 * @returns {Error} The error.
 */
module.exports = function enhanceError(error, config, code, request, response) {
  error.config = config;
  if (code) {
    error.code = code;
  }

  error.request = request;
  error.response = response;
  error.isAxiosError = true;

  error.toJSON = function toJSON() {
    return {
      // Standard
      message: this.message,
      name: this.name,
      // Microsoft
      description: this.description,
      number: this.number,
      // Mozilla
      fileName: this.fileName,
      lineNumber: this.lineNumber,
      columnNumber: this.columnNumber,
      stack: this.stack,
      // Axios
      config: this.config,
      code: this.code
    };
  };
  return error;
};


/***/ }),

/***/ 488:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);

/**
 * Config-specific merge-function which creates a new config-object
 * by merging two configuration objects together.
 *
 * @param {Object} config1
 * @param {Object} config2
 * @returns {Object} New object resulting from merging config2 to config1
 */
module.exports = function mergeConfig(config1, config2) {
  // eslint-disable-next-line no-param-reassign
  config2 = config2 || {};
  var config = {};

  var valueFromConfig2Keys = ['url', 'method', 'data'];
  var mergeDeepPropertiesKeys = ['headers', 'auth', 'proxy', 'params'];
  var defaultToConfig2Keys = [
    'baseURL', 'transformRequest', 'transformResponse', 'paramsSerializer',
    'timeout', 'timeoutMessage', 'withCredentials', 'adapter', 'responseType', 'xsrfCookieName',
    'xsrfHeaderName', 'onUploadProgress', 'onDownloadProgress', 'decompress',
    'maxContentLength', 'maxBodyLength', 'maxRedirects', 'transport', 'httpAgent',
    'httpsAgent', 'cancelToken', 'socketPath', 'responseEncoding'
  ];
  var directMergeKeys = ['validateStatus'];

  function getMergedValue(target, source) {
    if (utils.isPlainObject(target) && utils.isPlainObject(source)) {
      return utils.merge(target, source);
    } else if (utils.isPlainObject(source)) {
      return utils.merge({}, source);
    } else if (utils.isArray(source)) {
      return source.slice();
    }
    return source;
  }

  function mergeDeepProperties(prop) {
    if (!utils.isUndefined(config2[prop])) {
      config[prop] = getMergedValue(config1[prop], config2[prop]);
    } else if (!utils.isUndefined(config1[prop])) {
      config[prop] = getMergedValue(undefined, config1[prop]);
    }
  }

  utils.forEach(valueFromConfig2Keys, function valueFromConfig2(prop) {
    if (!utils.isUndefined(config2[prop])) {
      config[prop] = getMergedValue(undefined, config2[prop]);
    }
  });

  utils.forEach(mergeDeepPropertiesKeys, mergeDeepProperties);

  utils.forEach(defaultToConfig2Keys, function defaultToConfig2(prop) {
    if (!utils.isUndefined(config2[prop])) {
      config[prop] = getMergedValue(undefined, config2[prop]);
    } else if (!utils.isUndefined(config1[prop])) {
      config[prop] = getMergedValue(undefined, config1[prop]);
    }
  });

  utils.forEach(directMergeKeys, function merge(prop) {
    if (prop in config2) {
      config[prop] = getMergedValue(config1[prop], config2[prop]);
    } else if (prop in config1) {
      config[prop] = getMergedValue(undefined, config1[prop]);
    }
  });

  var axiosKeys = valueFromConfig2Keys
    .concat(mergeDeepPropertiesKeys)
    .concat(defaultToConfig2Keys)
    .concat(directMergeKeys);

  var otherKeys = Object
    .keys(config1)
    .concat(Object.keys(config2))
    .filter(function filterAxiosKeys(key) {
      return axiosKeys.indexOf(key) === -1;
    });

  utils.forEach(otherKeys, mergeDeepProperties);

  return config;
};


/***/ }),

/***/ 60:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var createError = __nccwpck_require__(287);

/**
 * Resolve or reject a Promise based on response status.
 *
 * @param {Function} resolve A function that resolves the promise.
 * @param {Function} reject A function that rejects the promise.
 * @param {object} response The response.
 */
module.exports = function settle(resolve, reject, response) {
  var validateStatus = response.config.validateStatus;
  if (!response.status || !validateStatus || validateStatus(response.status)) {
    resolve(response);
  } else {
    reject(createError(
      'Request failed with status code ' + response.status,
      response.config,
      null,
      response.request,
      response
    ));
  }
};


/***/ }),

/***/ 640:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);
var defaults = __nccwpck_require__(392);

/**
 * Transform the data for a request or a response
 *
 * @param {Object|String} data The data to be transformed
 * @param {Array} headers The headers for the request or response
 * @param {Array|Function} fns A single function or Array of functions
 * @returns {*} The resulting transformed data
 */
module.exports = function transformData(data, headers, fns) {
  var context = this || defaults;
  /*eslint no-param-reassign:0*/
  utils.forEach(fns, function transform(fn) {
    data = fn.call(context, data, headers);
  });

  return data;
};


/***/ }),

/***/ 392:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);
var normalizeHeaderName = __nccwpck_require__(37);
var enhanceError = __nccwpck_require__(234);

var DEFAULT_CONTENT_TYPE = {
  'Content-Type': 'application/x-www-form-urlencoded'
};

function setContentTypeIfUnset(headers, value) {
  if (!utils.isUndefined(headers) && utils.isUndefined(headers['Content-Type'])) {
    headers['Content-Type'] = value;
  }
}

function getDefaultAdapter() {
  var adapter;
  if (typeof XMLHttpRequest !== 'undefined') {
    // For browsers use XHR adapter
    adapter = __nccwpck_require__(697);
  } else if (typeof process !== 'undefined' && Object.prototype.toString.call(process) === '[object process]') {
    // For node use HTTP adapter
    adapter = __nccwpck_require__(431);
  }
  return adapter;
}

function stringifySafely(rawValue, parser, encoder) {
  if (utils.isString(rawValue)) {
    try {
      (parser || JSON.parse)(rawValue);
      return utils.trim(rawValue);
    } catch (e) {
      if (e.name !== 'SyntaxError') {
        throw e;
      }
    }
  }

  return (encoder || JSON.stringify)(rawValue);
}

var defaults = {

  transitional: {
    silentJSONParsing: true,
    forcedJSONParsing: true,
    clarifyTimeoutError: false
  },

  adapter: getDefaultAdapter(),

  transformRequest: [function transformRequest(data, headers) {
    normalizeHeaderName(headers, 'Accept');
    normalizeHeaderName(headers, 'Content-Type');

    if (utils.isFormData(data) ||
      utils.isArrayBuffer(data) ||
      utils.isBuffer(data) ||
      utils.isStream(data) ||
      utils.isFile(data) ||
      utils.isBlob(data)
    ) {
      return data;
    }
    if (utils.isArrayBufferView(data)) {
      return data.buffer;
    }
    if (utils.isURLSearchParams(data)) {
      setContentTypeIfUnset(headers, 'application/x-www-form-urlencoded;charset=utf-8');
      return data.toString();
    }
    if (utils.isObject(data) || (headers && headers['Content-Type'] === 'application/json')) {
      setContentTypeIfUnset(headers, 'application/json');
      return stringifySafely(data);
    }
    return data;
  }],

  transformResponse: [function transformResponse(data) {
    var transitional = this.transitional;
    var silentJSONParsing = transitional && transitional.silentJSONParsing;
    var forcedJSONParsing = transitional && transitional.forcedJSONParsing;
    var strictJSONParsing = !silentJSONParsing && this.responseType === 'json';

    if (strictJSONParsing || (forcedJSONParsing && utils.isString(data) && data.length)) {
      try {
        return JSON.parse(data);
      } catch (e) {
        if (strictJSONParsing) {
          if (e.name === 'SyntaxError') {
            throw enhanceError(e, this, 'E_JSON_PARSE');
          }
          throw e;
        }
      }
    }

    return data;
  }],

  /**
   * A timeout in milliseconds to abort a request. If set to 0 (default) a
   * timeout is not created.
   */
  timeout: 0,

  xsrfCookieName: 'XSRF-TOKEN',
  xsrfHeaderName: 'X-XSRF-TOKEN',

  maxContentLength: -1,
  maxBodyLength: -1,

  validateStatus: function validateStatus(status) {
    return status >= 200 && status < 300;
  }
};

defaults.headers = {
  common: {
    'Accept': 'application/json, text/plain, */*'
  }
};

utils.forEach(['delete', 'get', 'head'], function forEachMethodNoData(method) {
  defaults.headers[method] = {};
});

utils.forEach(['post', 'put', 'patch'], function forEachMethodWithData(method) {
  defaults.headers[method] = utils.merge(DEFAULT_CONTENT_TYPE);
});

module.exports = defaults;


/***/ }),

/***/ 117:
/***/ ((module) => {

"use strict";


module.exports = function bind(fn, thisArg) {
  return function wrap() {
    var args = new Array(arguments.length);
    for (var i = 0; i < args.length; i++) {
      args[i] = arguments[i];
    }
    return fn.apply(thisArg, args);
  };
};


/***/ }),

/***/ 358:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);

function encode(val) {
  return encodeURIComponent(val).
    replace(/%3A/gi, ':').
    replace(/%24/g, '$').
    replace(/%2C/gi, ',').
    replace(/%20/g, '+').
    replace(/%5B/gi, '[').
    replace(/%5D/gi, ']');
}

/**
 * Build a URL by appending params to the end
 *
 * @param {string} url The base of the url (e.g., http://www.google.com)
 * @param {object} [params] The params to be appended
 * @returns {string} The formatted url
 */
module.exports = function buildURL(url, params, paramsSerializer) {
  /*eslint no-param-reassign:0*/
  if (!params) {
    return url;
  }

  var serializedParams;
  if (paramsSerializer) {
    serializedParams = paramsSerializer(params);
  } else if (utils.isURLSearchParams(params)) {
    serializedParams = params.toString();
  } else {
    var parts = [];

    utils.forEach(params, function serialize(val, key) {
      if (val === null || typeof val === 'undefined') {
        return;
      }

      if (utils.isArray(val)) {
        key = key + '[]';
      } else {
        val = [val];
      }

      utils.forEach(val, function parseValue(v) {
        if (utils.isDate(v)) {
          v = v.toISOString();
        } else if (utils.isObject(v)) {
          v = JSON.stringify(v);
        }
        parts.push(encode(key) + '=' + encode(v));
      });
    });

    serializedParams = parts.join('&');
  }

  if (serializedParams) {
    var hashmarkIndex = url.indexOf('#');
    if (hashmarkIndex !== -1) {
      url = url.slice(0, hashmarkIndex);
    }

    url += (url.indexOf('?') === -1 ? '?' : '&') + serializedParams;
  }

  return url;
};


/***/ }),

/***/ 4:
/***/ ((module) => {

"use strict";


/**
 * Creates a new URL by combining the specified URLs
 *
 * @param {string} baseURL The base URL
 * @param {string} relativeURL The relative URL
 * @returns {string} The combined URL
 */
module.exports = function combineURLs(baseURL, relativeURL) {
  return relativeURL
    ? baseURL.replace(/\/+$/, '') + '/' + relativeURL.replace(/^\/+/, '')
    : baseURL;
};


/***/ }),

/***/ 722:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);

module.exports = (
  utils.isStandardBrowserEnv() ?

  // Standard browser envs support document.cookie
    (function standardBrowserEnv() {
      return {
        write: function write(name, value, expires, path, domain, secure) {
          var cookie = [];
          cookie.push(name + '=' + encodeURIComponent(value));

          if (utils.isNumber(expires)) {
            cookie.push('expires=' + new Date(expires).toGMTString());
          }

          if (utils.isString(path)) {
            cookie.push('path=' + path);
          }

          if (utils.isString(domain)) {
            cookie.push('domain=' + domain);
          }

          if (secure === true) {
            cookie.push('secure');
          }

          document.cookie = cookie.join('; ');
        },

        read: function read(name) {
          var match = document.cookie.match(new RegExp('(^|;\\s*)(' + name + ')=([^;]*)'));
          return (match ? decodeURIComponent(match[3]) : null);
        },

        remove: function remove(name) {
          this.write(name, '', Date.now() - 86400000);
        }
      };
    })() :

  // Non standard browser env (web workers, react-native) lack needed support.
    (function nonStandardBrowserEnv() {
      return {
        write: function write() {},
        read: function read() { return null; },
        remove: function remove() {}
      };
    })()
);


/***/ }),

/***/ 187:
/***/ ((module) => {

"use strict";


/**
 * Determines whether the specified URL is absolute
 *
 * @param {string} url The URL to test
 * @returns {boolean} True if the specified URL is absolute, otherwise false
 */
module.exports = function isAbsoluteURL(url) {
  // A URL is considered absolute if it begins with "<scheme>://" or "//" (protocol-relative URL).
  // RFC 3986 defines scheme name as a sequence of characters beginning with a letter and followed
  // by any combination of letters, digits, plus, period, or hyphen.
  return /^([a-z][a-z\d\+\-\.]*:)?\/\//i.test(url);
};


/***/ }),

/***/ 446:
/***/ ((module) => {

"use strict";


/**
 * Determines whether the payload is an error thrown by Axios
 *
 * @param {*} payload The value to test
 * @returns {boolean} True if the payload is an error thrown by Axios, otherwise false
 */
module.exports = function isAxiosError(payload) {
  return (typeof payload === 'object') && (payload.isAxiosError === true);
};


/***/ }),

/***/ 913:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);

module.exports = (
  utils.isStandardBrowserEnv() ?

  // Standard browser envs have full support of the APIs needed to test
  // whether the request URL is of the same origin as current location.
    (function standardBrowserEnv() {
      var msie = /(msie|trident)/i.test(navigator.userAgent);
      var urlParsingNode = document.createElement('a');
      var originURL;

      /**
    * Parse a URL to discover it's components
    *
    * @param {String} url The URL to be parsed
    * @returns {Object}
    */
      function resolveURL(url) {
        var href = url;

        if (msie) {
        // IE needs attribute set twice to normalize properties
          urlParsingNode.setAttribute('href', href);
          href = urlParsingNode.href;
        }

        urlParsingNode.setAttribute('href', href);

        // urlParsingNode provides the UrlUtils interface - http://url.spec.whatwg.org/#urlutils
        return {
          href: urlParsingNode.href,
          protocol: urlParsingNode.protocol ? urlParsingNode.protocol.replace(/:$/, '') : '',
          host: urlParsingNode.host,
          search: urlParsingNode.search ? urlParsingNode.search.replace(/^\?/, '') : '',
          hash: urlParsingNode.hash ? urlParsingNode.hash.replace(/^#/, '') : '',
          hostname: urlParsingNode.hostname,
          port: urlParsingNode.port,
          pathname: (urlParsingNode.pathname.charAt(0) === '/') ?
            urlParsingNode.pathname :
            '/' + urlParsingNode.pathname
        };
      }

      originURL = resolveURL(window.location.href);

      /**
    * Determine if a URL shares the same origin as the current location
    *
    * @param {String} requestURL The URL to test
    * @returns {boolean} True if URL shares the same origin, otherwise false
    */
      return function isURLSameOrigin(requestURL) {
        var parsed = (utils.isString(requestURL)) ? resolveURL(requestURL) : requestURL;
        return (parsed.protocol === originURL.protocol &&
            parsed.host === originURL.host);
      };
    })() :

  // Non standard browser envs (web workers, react-native) lack needed support.
    (function nonStandardBrowserEnv() {
      return function isURLSameOrigin() {
        return true;
      };
    })()
);


/***/ }),

/***/ 37:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);

module.exports = function normalizeHeaderName(headers, normalizedName) {
  utils.forEach(headers, function processHeader(value, name) {
    if (name !== normalizedName && name.toUpperCase() === normalizedName.toUpperCase()) {
      headers[normalizedName] = value;
      delete headers[name];
    }
  });
};


/***/ }),

/***/ 695:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var utils = __nccwpck_require__(474);

// Headers whose duplicates are ignored by node
// c.f. https://nodejs.org/api/http.html#http_message_headers
var ignoreDuplicateOf = [
  'age', 'authorization', 'content-length', 'content-type', 'etag',
  'expires', 'from', 'host', 'if-modified-since', 'if-unmodified-since',
  'last-modified', 'location', 'max-forwards', 'proxy-authorization',
  'referer', 'retry-after', 'user-agent'
];

/**
 * Parse headers into an object
 *
 * ```
 * Date: Wed, 27 Aug 2014 08:58:49 GMT
 * Content-Type: application/json
 * Connection: keep-alive
 * Transfer-Encoding: chunked
 * ```
 *
 * @param {String} headers Headers needing to be parsed
 * @returns {Object} Headers parsed into an object
 */
module.exports = function parseHeaders(headers) {
  var parsed = {};
  var key;
  var val;
  var i;

  if (!headers) { return parsed; }

  utils.forEach(headers.split('\n'), function parser(line) {
    i = line.indexOf(':');
    key = utils.trim(line.substr(0, i)).toLowerCase();
    val = utils.trim(line.substr(i + 1));

    if (key) {
      if (parsed[key] && ignoreDuplicateOf.indexOf(key) >= 0) {
        return;
      }
      if (key === 'set-cookie') {
        parsed[key] = (parsed[key] ? parsed[key] : []).concat([val]);
      } else {
        parsed[key] = parsed[key] ? parsed[key] + ', ' + val : val;
      }
    }
  });

  return parsed;
};


/***/ }),

/***/ 666:
/***/ ((module) => {

"use strict";


/**
 * Syntactic sugar for invoking a function and expanding an array for arguments.
 *
 * Common use case would be to use `Function.prototype.apply`.
 *
 *  ```js
 *  function f(x, y, z) {}
 *  var args = [1, 2, 3];
 *  f.apply(null, args);
 *  ```
 *
 * With `spread` this example can be re-written.
 *
 *  ```js
 *  spread(function(x, y, z) {})([1, 2, 3]);
 *  ```
 *
 * @param {Function} callback
 * @returns {Function}
 */
module.exports = function spread(callback) {
  return function wrap(arr) {
    return callback.apply(null, arr);
  };
};


/***/ }),

/***/ 359:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var pkg = __nccwpck_require__(837);

var validators = {};

// eslint-disable-next-line func-names
['object', 'boolean', 'number', 'function', 'string', 'symbol'].forEach(function(type, i) {
  validators[type] = function validator(thing) {
    return typeof thing === type || 'a' + (i < 1 ? 'n ' : ' ') + type;
  };
});

var deprecatedWarnings = {};
var currentVerArr = pkg.version.split('.');

/**
 * Compare package versions
 * @param {string} version
 * @param {string?} thanVersion
 * @returns {boolean}
 */
function isOlderVersion(version, thanVersion) {
  var pkgVersionArr = thanVersion ? thanVersion.split('.') : currentVerArr;
  var destVer = version.split('.');
  for (var i = 0; i < 3; i++) {
    if (pkgVersionArr[i] > destVer[i]) {
      return true;
    } else if (pkgVersionArr[i] < destVer[i]) {
      return false;
    }
  }
  return false;
}

/**
 * Transitional option validator
 * @param {function|boolean?} validator
 * @param {string?} version
 * @param {string} message
 * @returns {function}
 */
validators.transitional = function transitional(validator, version, message) {
  var isDeprecated = version && isOlderVersion(version);

  function formatMessage(opt, desc) {
    return '[Axios v' + pkg.version + '] Transitional option \'' + opt + '\'' + desc + (message ? '. ' + message : '');
  }

  // eslint-disable-next-line func-names
  return function(value, opt, opts) {
    if (validator === false) {
      throw new Error(formatMessage(opt, ' has been removed in ' + version));
    }

    if (isDeprecated && !deprecatedWarnings[opt]) {
      deprecatedWarnings[opt] = true;
      // eslint-disable-next-line no-console
      console.warn(
        formatMessage(
          opt,
          ' has been deprecated since v' + version + ' and will be removed in the near future'
        )
      );
    }

    return validator ? validator(value, opt, opts) : true;
  };
};

/**
 * Assert object's properties type
 * @param {object} options
 * @param {object} schema
 * @param {boolean?} allowUnknown
 */

function assertOptions(options, schema, allowUnknown) {
  if (typeof options !== 'object') {
    throw new TypeError('options must be an object');
  }
  var keys = Object.keys(options);
  var i = keys.length;
  while (i-- > 0) {
    var opt = keys[i];
    var validator = schema[opt];
    if (validator) {
      var value = options[opt];
      var result = value === undefined || validator(value, opt, options);
      if (result !== true) {
        throw new TypeError('option ' + opt + ' must be ' + result);
      }
      continue;
    }
    if (allowUnknown !== true) {
      throw Error('Unknown option ' + opt);
    }
  }
}

module.exports = {
  isOlderVersion: isOlderVersion,
  assertOptions: assertOptions,
  validators: validators
};


/***/ }),

/***/ 474:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var bind = __nccwpck_require__(117);

// utils is a library of generic helper functions non-specific to axios

var toString = Object.prototype.toString;

/**
 * Determine if a value is an Array
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is an Array, otherwise false
 */
function isArray(val) {
  return toString.call(val) === '[object Array]';
}

/**
 * Determine if a value is undefined
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if the value is undefined, otherwise false
 */
function isUndefined(val) {
  return typeof val === 'undefined';
}

/**
 * Determine if a value is a Buffer
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a Buffer, otherwise false
 */
function isBuffer(val) {
  return val !== null && !isUndefined(val) && val.constructor !== null && !isUndefined(val.constructor)
    && typeof val.constructor.isBuffer === 'function' && val.constructor.isBuffer(val);
}

/**
 * Determine if a value is an ArrayBuffer
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is an ArrayBuffer, otherwise false
 */
function isArrayBuffer(val) {
  return toString.call(val) === '[object ArrayBuffer]';
}

/**
 * Determine if a value is a FormData
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is an FormData, otherwise false
 */
function isFormData(val) {
  return (typeof FormData !== 'undefined') && (val instanceof FormData);
}

/**
 * Determine if a value is a view on an ArrayBuffer
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a view on an ArrayBuffer, otherwise false
 */
function isArrayBufferView(val) {
  var result;
  if ((typeof ArrayBuffer !== 'undefined') && (ArrayBuffer.isView)) {
    result = ArrayBuffer.isView(val);
  } else {
    result = (val) && (val.buffer) && (val.buffer instanceof ArrayBuffer);
  }
  return result;
}

/**
 * Determine if a value is a String
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a String, otherwise false
 */
function isString(val) {
  return typeof val === 'string';
}

/**
 * Determine if a value is a Number
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a Number, otherwise false
 */
function isNumber(val) {
  return typeof val === 'number';
}

/**
 * Determine if a value is an Object
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is an Object, otherwise false
 */
function isObject(val) {
  return val !== null && typeof val === 'object';
}

/**
 * Determine if a value is a plain Object
 *
 * @param {Object} val The value to test
 * @return {boolean} True if value is a plain Object, otherwise false
 */
function isPlainObject(val) {
  if (toString.call(val) !== '[object Object]') {
    return false;
  }

  var prototype = Object.getPrototypeOf(val);
  return prototype === null || prototype === Object.prototype;
}

/**
 * Determine if a value is a Date
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a Date, otherwise false
 */
function isDate(val) {
  return toString.call(val) === '[object Date]';
}

/**
 * Determine if a value is a File
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a File, otherwise false
 */
function isFile(val) {
  return toString.call(val) === '[object File]';
}

/**
 * Determine if a value is a Blob
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a Blob, otherwise false
 */
function isBlob(val) {
  return toString.call(val) === '[object Blob]';
}

/**
 * Determine if a value is a Function
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a Function, otherwise false
 */
function isFunction(val) {
  return toString.call(val) === '[object Function]';
}

/**
 * Determine if a value is a Stream
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a Stream, otherwise false
 */
function isStream(val) {
  return isObject(val) && isFunction(val.pipe);
}

/**
 * Determine if a value is a URLSearchParams object
 *
 * @param {Object} val The value to test
 * @returns {boolean} True if value is a URLSearchParams object, otherwise false
 */
function isURLSearchParams(val) {
  return typeof URLSearchParams !== 'undefined' && val instanceof URLSearchParams;
}

/**
 * Trim excess whitespace off the beginning and end of a string
 *
 * @param {String} str The String to trim
 * @returns {String} The String freed of excess whitespace
 */
function trim(str) {
  return str.trim ? str.trim() : str.replace(/^\s+|\s+$/g, '');
}

/**
 * Determine if we're running in a standard browser environment
 *
 * This allows axios to run in a web worker, and react-native.
 * Both environments support XMLHttpRequest, but not fully standard globals.
 *
 * web workers:
 *  typeof window -> undefined
 *  typeof document -> undefined
 *
 * react-native:
 *  navigator.product -> 'ReactNative'
 * nativescript
 *  navigator.product -> 'NativeScript' or 'NS'
 */
function isStandardBrowserEnv() {
  if (typeof navigator !== 'undefined' && (navigator.product === 'ReactNative' ||
                                           navigator.product === 'NativeScript' ||
                                           navigator.product === 'NS')) {
    return false;
  }
  return (
    typeof window !== 'undefined' &&
    typeof document !== 'undefined'
  );
}

/**
 * Iterate over an Array or an Object invoking a function for each item.
 *
 * If `obj` is an Array callback will be called passing
 * the value, index, and complete array for each item.
 *
 * If 'obj' is an Object callback will be called passing
 * the value, key, and complete object for each property.
 *
 * @param {Object|Array} obj The object to iterate
 * @param {Function} fn The callback to invoke for each item
 */
function forEach(obj, fn) {
  // Don't bother if no value provided
  if (obj === null || typeof obj === 'undefined') {
    return;
  }

  // Force an array if not already something iterable
  if (typeof obj !== 'object') {
    /*eslint no-param-reassign:0*/
    obj = [obj];
  }

  if (isArray(obj)) {
    // Iterate over array values
    for (var i = 0, l = obj.length; i < l; i++) {
      fn.call(null, obj[i], i, obj);
    }
  } else {
    // Iterate over object keys
    for (var key in obj) {
      if (Object.prototype.hasOwnProperty.call(obj, key)) {
        fn.call(null, obj[key], key, obj);
      }
    }
  }
}

/**
 * Accepts varargs expecting each argument to be an object, then
 * immutably merges the properties of each object and returns result.
 *
 * When multiple objects contain the same key the later object in
 * the arguments list will take precedence.
 *
 * Example:
 *
 * ```js
 * var result = merge({foo: 123}, {foo: 456});
 * console.log(result.foo); // outputs 456
 * ```
 *
 * @param {Object} obj1 Object to merge
 * @returns {Object} Result of all merge properties
 */
function merge(/* obj1, obj2, obj3, ... */) {
  var result = {};
  function assignValue(val, key) {
    if (isPlainObject(result[key]) && isPlainObject(val)) {
      result[key] = merge(result[key], val);
    } else if (isPlainObject(val)) {
      result[key] = merge({}, val);
    } else if (isArray(val)) {
      result[key] = val.slice();
    } else {
      result[key] = val;
    }
  }

  for (var i = 0, l = arguments.length; i < l; i++) {
    forEach(arguments[i], assignValue);
  }
  return result;
}

/**
 * Extends object a by mutably adding to it the properties of object b.
 *
 * @param {Object} a The object to be extended
 * @param {Object} b The object to copy properties from
 * @param {Object} thisArg The object to bind function to
 * @return {Object} The resulting value of object a
 */
function extend(a, b, thisArg) {
  forEach(b, function assignValue(val, key) {
    if (thisArg && typeof val === 'function') {
      a[key] = bind(val, thisArg);
    } else {
      a[key] = val;
    }
  });
  return a;
}

/**
 * Remove byte order marker. This catches EF BB BF (the UTF-8 BOM)
 *
 * @param {string} content with BOM
 * @return {string} content value without BOM
 */
function stripBOM(content) {
  if (content.charCodeAt(0) === 0xFEFF) {
    content = content.slice(1);
  }
  return content;
}

module.exports = {
  isArray: isArray,
  isArrayBuffer: isArrayBuffer,
  isBuffer: isBuffer,
  isFormData: isFormData,
  isArrayBufferView: isArrayBufferView,
  isString: isString,
  isNumber: isNumber,
  isObject: isObject,
  isPlainObject: isPlainObject,
  isUndefined: isUndefined,
  isDate: isDate,
  isFile: isFile,
  isBlob: isBlob,
  isFunction: isFunction,
  isStream: isStream,
  isURLSearchParams: isURLSearchParams,
  isStandardBrowserEnv: isStandardBrowserEnv,
  forEach: forEach,
  merge: merge,
  extend: extend,
  trim: trim,
  stripBOM: stripBOM
};


/***/ }),

/***/ 114:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

var debug;

module.exports = function () {
  if (!debug) {
    try {
      /* eslint global-require: off */
      debug = __nccwpck_require__(267)("follow-redirects");
    }
    catch (error) { /* */ }
    if (typeof debug !== "function") {
      debug = function () { /* */ };
    }
  }
  debug.apply(null, arguments);
};


/***/ }),

/***/ 774:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

var url = __nccwpck_require__(310);
var URL = url.URL;
var http = __nccwpck_require__(685);
var https = __nccwpck_require__(687);
var Writable = (__nccwpck_require__(781).Writable);
var assert = __nccwpck_require__(491);
var debug = __nccwpck_require__(114);

// Create handlers that pass events from native requests
var events = ["abort", "aborted", "connect", "error", "socket", "timeout"];
var eventHandlers = Object.create(null);
events.forEach(function (event) {
  eventHandlers[event] = function (arg1, arg2, arg3) {
    this._redirectable.emit(event, arg1, arg2, arg3);
  };
});

// Error types with codes
var RedirectionError = createErrorType(
  "ERR_FR_REDIRECTION_FAILURE",
  "Redirected request failed"
);
var TooManyRedirectsError = createErrorType(
  "ERR_FR_TOO_MANY_REDIRECTS",
  "Maximum number of redirects exceeded"
);
var MaxBodyLengthExceededError = createErrorType(
  "ERR_FR_MAX_BODY_LENGTH_EXCEEDED",
  "Request body larger than maxBodyLength limit"
);
var WriteAfterEndError = createErrorType(
  "ERR_STREAM_WRITE_AFTER_END",
  "write after end"
);

// An HTTP(S) request that can be redirected
function RedirectableRequest(options, responseCallback) {
  // Initialize the request
  Writable.call(this);
  this._sanitizeOptions(options);
  this._options = options;
  this._ended = false;
  this._ending = false;
  this._redirectCount = 0;
  this._redirects = [];
  this._requestBodyLength = 0;
  this._requestBodyBuffers = [];

  // Attach a callback if passed
  if (responseCallback) {
    this.on("response", responseCallback);
  }

  // React to responses of native requests
  var self = this;
  this._onNativeResponse = function (response) {
    self._processResponse(response);
  };

  // Perform the first request
  this._performRequest();
}
RedirectableRequest.prototype = Object.create(Writable.prototype);

RedirectableRequest.prototype.abort = function () {
  abortRequest(this._currentRequest);
  this.emit("abort");
};

// Writes buffered data to the current native request
RedirectableRequest.prototype.write = function (data, encoding, callback) {
  // Writing is not allowed if end has been called
  if (this._ending) {
    throw new WriteAfterEndError();
  }

  // Validate input and shift parameters if necessary
  if (!(typeof data === "string" || typeof data === "object" && ("length" in data))) {
    throw new TypeError("data should be a string, Buffer or Uint8Array");
  }
  if (typeof encoding === "function") {
    callback = encoding;
    encoding = null;
  }

  // Ignore empty buffers, since writing them doesn't invoke the callback
  // https://github.com/nodejs/node/issues/22066
  if (data.length === 0) {
    if (callback) {
      callback();
    }
    return;
  }
  // Only write when we don't exceed the maximum body length
  if (this._requestBodyLength + data.length <= this._options.maxBodyLength) {
    this._requestBodyLength += data.length;
    this._requestBodyBuffers.push({ data: data, encoding: encoding });
    this._currentRequest.write(data, encoding, callback);
  }
  // Error when we exceed the maximum body length
  else {
    this.emit("error", new MaxBodyLengthExceededError());
    this.abort();
  }
};

// Ends the current native request
RedirectableRequest.prototype.end = function (data, encoding, callback) {
  // Shift parameters if necessary
  if (typeof data === "function") {
    callback = data;
    data = encoding = null;
  }
  else if (typeof encoding === "function") {
    callback = encoding;
    encoding = null;
  }

  // Write data if needed and end
  if (!data) {
    this._ended = this._ending = true;
    this._currentRequest.end(null, null, callback);
  }
  else {
    var self = this;
    var currentRequest = this._currentRequest;
    this.write(data, encoding, function () {
      self._ended = true;
      currentRequest.end(null, null, callback);
    });
    this._ending = true;
  }
};

// Sets a header value on the current native request
RedirectableRequest.prototype.setHeader = function (name, value) {
  this._options.headers[name] = value;
  this._currentRequest.setHeader(name, value);
};

// Clears a header value on the current native request
RedirectableRequest.prototype.removeHeader = function (name) {
  delete this._options.headers[name];
  this._currentRequest.removeHeader(name);
};

// Global timeout for all underlying requests
RedirectableRequest.prototype.setTimeout = function (msecs, callback) {
  var self = this;

  // Destroys the socket on timeout
  function destroyOnTimeout(socket) {
    socket.setTimeout(msecs);
    socket.removeListener("timeout", socket.destroy);
    socket.addListener("timeout", socket.destroy);
  }

  // Sets up a timer to trigger a timeout event
  function startTimer(socket) {
    if (self._timeout) {
      clearTimeout(self._timeout);
    }
    self._timeout = setTimeout(function () {
      self.emit("timeout");
      clearTimer();
    }, msecs);
    destroyOnTimeout(socket);
  }

  // Stops a timeout from triggering
  function clearTimer() {
    // Clear the timeout
    if (self._timeout) {
      clearTimeout(self._timeout);
      self._timeout = null;
    }

    // Clean up all attached listeners
    self.removeListener("abort", clearTimer);
    self.removeListener("error", clearTimer);
    self.removeListener("response", clearTimer);
    if (callback) {
      self.removeListener("timeout", callback);
    }
    if (!self.socket) {
      self._currentRequest.removeListener("socket", startTimer);
    }
  }

  // Attach callback if passed
  if (callback) {
    this.on("timeout", callback);
  }

  // Start the timer if or when the socket is opened
  if (this.socket) {
    startTimer(this.socket);
  }
  else {
    this._currentRequest.once("socket", startTimer);
  }

  // Clean up on events
  this.on("socket", destroyOnTimeout);
  this.on("abort", clearTimer);
  this.on("error", clearTimer);
  this.on("response", clearTimer);

  return this;
};

// Proxy all other public ClientRequest methods
[
  "flushHeaders", "getHeader",
  "setNoDelay", "setSocketKeepAlive",
].forEach(function (method) {
  RedirectableRequest.prototype[method] = function (a, b) {
    return this._currentRequest[method](a, b);
  };
});

// Proxy all public ClientRequest properties
["aborted", "connection", "socket"].forEach(function (property) {
  Object.defineProperty(RedirectableRequest.prototype, property, {
    get: function () { return this._currentRequest[property]; },
  });
});

RedirectableRequest.prototype._sanitizeOptions = function (options) {
  // Ensure headers are always present
  if (!options.headers) {
    options.headers = {};
  }

  // Since http.request treats host as an alias of hostname,
  // but the url module interprets host as hostname plus port,
  // eliminate the host property to avoid confusion.
  if (options.host) {
    // Use hostname if set, because it has precedence
    if (!options.hostname) {
      options.hostname = options.host;
    }
    delete options.host;
  }

  // Complete the URL object when necessary
  if (!options.pathname && options.path) {
    var searchPos = options.path.indexOf("?");
    if (searchPos < 0) {
      options.pathname = options.path;
    }
    else {
      options.pathname = options.path.substring(0, searchPos);
      options.search = options.path.substring(searchPos);
    }
  }
};


// Executes the next native request (initial or redirect)
RedirectableRequest.prototype._performRequest = function () {
  // Load the native protocol
  var protocol = this._options.protocol;
  var nativeProtocol = this._options.nativeProtocols[protocol];
  if (!nativeProtocol) {
    this.emit("error", new TypeError("Unsupported protocol " + protocol));
    return;
  }

  // If specified, use the agent corresponding to the protocol
  // (HTTP and HTTPS use different types of agents)
  if (this._options.agents) {
    var scheme = protocol.substr(0, protocol.length - 1);
    this._options.agent = this._options.agents[scheme];
  }

  // Create the native request
  var request = this._currentRequest =
        nativeProtocol.request(this._options, this._onNativeResponse);
  this._currentUrl = url.format(this._options);

  // Set up event handlers
  request._redirectable = this;
  for (var e = 0; e < events.length; e++) {
    request.on(events[e], eventHandlers[events[e]]);
  }

  // End a redirected request
  // (The first request must be ended explicitly with RedirectableRequest#end)
  if (this._isRedirect) {
    // Write the request entity and end.
    var i = 0;
    var self = this;
    var buffers = this._requestBodyBuffers;
    (function writeNext(error) {
      // Only write if this request has not been redirected yet
      /* istanbul ignore else */
      if (request === self._currentRequest) {
        // Report any write errors
        /* istanbul ignore if */
        if (error) {
          self.emit("error", error);
        }
        // Write the next buffer if there are still left
        else if (i < buffers.length) {
          var buffer = buffers[i++];
          /* istanbul ignore else */
          if (!request.finished) {
            request.write(buffer.data, buffer.encoding, writeNext);
          }
        }
        // End the request if `end` has been called on us
        else if (self._ended) {
          request.end();
        }
      }
    }());
  }
};

// Processes a response from the current native request
RedirectableRequest.prototype._processResponse = function (response) {
  // Store the redirected response
  var statusCode = response.statusCode;
  if (this._options.trackRedirects) {
    this._redirects.push({
      url: this._currentUrl,
      headers: response.headers,
      statusCode: statusCode,
    });
  }

  // RFC7231§6.4: The 3xx (Redirection) class of status code indicates
  // that further action needs to be taken by the user agent in order to
  // fulfill the request. If a Location header field is provided,
  // the user agent MAY automatically redirect its request to the URI
  // referenced by the Location field value,
  // even if the specific status code is not understood.

  // If the response is not a redirect; return it as-is
  var location = response.headers.location;
  if (!location || this._options.followRedirects === false ||
      statusCode < 300 || statusCode >= 400) {
    response.responseUrl = this._currentUrl;
    response.redirects = this._redirects;
    this.emit("response", response);

    // Clean up
    this._requestBodyBuffers = [];
    return;
  }

  // The response is a redirect, so abort the current request
  abortRequest(this._currentRequest);
  // Discard the remainder of the response to avoid waiting for data
  response.destroy();

  // RFC7231§6.4: A client SHOULD detect and intervene
  // in cyclical redirections (i.e., "infinite" redirection loops).
  if (++this._redirectCount > this._options.maxRedirects) {
    this.emit("error", new TooManyRedirectsError());
    return;
  }

  // RFC7231§6.4: Automatic redirection needs to done with
  // care for methods not known to be safe, […]
  // RFC7231§6.4.2–3: For historical reasons, a user agent MAY change
  // the request method from POST to GET for the subsequent request.
  if ((statusCode === 301 || statusCode === 302) && this._options.method === "POST" ||
      // RFC7231§6.4.4: The 303 (See Other) status code indicates that
      // the server is redirecting the user agent to a different resource […]
      // A user agent can perform a retrieval request targeting that URI
      // (a GET or HEAD request if using HTTP) […]
      (statusCode === 303) && !/^(?:GET|HEAD)$/.test(this._options.method)) {
    this._options.method = "GET";
    // Drop a possible entity and headers related to it
    this._requestBodyBuffers = [];
    removeMatchingHeaders(/^content-/i, this._options.headers);
  }

  // Drop the Host header, as the redirect might lead to a different host
  var currentHostHeader = removeMatchingHeaders(/^host$/i, this._options.headers);

  // If the redirect is relative, carry over the host of the last request
  var currentUrlParts = url.parse(this._currentUrl);
  var currentHost = currentHostHeader || currentUrlParts.host;
  var currentUrl = /^\w+:/.test(location) ? this._currentUrl :
    url.format(Object.assign(currentUrlParts, { host: currentHost }));

  // Determine the URL of the redirection
  var redirectUrl;
  try {
    redirectUrl = url.resolve(currentUrl, location);
  }
  catch (cause) {
    this.emit("error", new RedirectionError(cause));
    return;
  }

  // Create the redirected request
  debug("redirecting to", redirectUrl);
  this._isRedirect = true;
  var redirectUrlParts = url.parse(redirectUrl);
  Object.assign(this._options, redirectUrlParts);

  // Drop confidential headers when redirecting to a less secure protocol
  // or to a different domain that is not a superdomain
  if (redirectUrlParts.protocol !== currentUrlParts.protocol &&
     redirectUrlParts.protocol !== "https:" ||
     redirectUrlParts.host !== currentHost &&
     !isSubdomain(redirectUrlParts.host, currentHost)) {
    removeMatchingHeaders(/^(?:authorization|cookie)$/i, this._options.headers);
  }

  // Evaluate the beforeRedirect callback
  if (typeof this._options.beforeRedirect === "function") {
    var responseDetails = { headers: response.headers };
    try {
      this._options.beforeRedirect.call(null, this._options, responseDetails);
    }
    catch (err) {
      this.emit("error", err);
      return;
    }
    this._sanitizeOptions(this._options);
  }

  // Perform the redirected request
  try {
    this._performRequest();
  }
  catch (cause) {
    this.emit("error", new RedirectionError(cause));
  }
};

// Wraps the key/value object of protocols with redirect functionality
function wrap(protocols) {
  // Default settings
  var exports = {
    maxRedirects: 21,
    maxBodyLength: 10 * 1024 * 1024,
  };

  // Wrap each protocol
  var nativeProtocols = {};
  Object.keys(protocols).forEach(function (scheme) {
    var protocol = scheme + ":";
    var nativeProtocol = nativeProtocols[protocol] = protocols[scheme];
    var wrappedProtocol = exports[scheme] = Object.create(nativeProtocol);

    // Executes a request, following redirects
    function request(input, options, callback) {
      // Parse parameters
      if (typeof input === "string") {
        var urlStr = input;
        try {
          input = urlToOptions(new URL(urlStr));
        }
        catch (err) {
          /* istanbul ignore next */
          input = url.parse(urlStr);
        }
      }
      else if (URL && (input instanceof URL)) {
        input = urlToOptions(input);
      }
      else {
        callback = options;
        options = input;
        input = { protocol: protocol };
      }
      if (typeof options === "function") {
        callback = options;
        options = null;
      }

      // Set defaults
      options = Object.assign({
        maxRedirects: exports.maxRedirects,
        maxBodyLength: exports.maxBodyLength,
      }, input, options);
      options.nativeProtocols = nativeProtocols;

      assert.equal(options.protocol, protocol, "protocol mismatch");
      debug("options", options);
      return new RedirectableRequest(options, callback);
    }

    // Executes a GET request, following redirects
    function get(input, options, callback) {
      var wrappedRequest = wrappedProtocol.request(input, options, callback);
      wrappedRequest.end();
      return wrappedRequest;
    }

    // Expose the properties on the wrapped protocol
    Object.defineProperties(wrappedProtocol, {
      request: { value: request, configurable: true, enumerable: true, writable: true },
      get: { value: get, configurable: true, enumerable: true, writable: true },
    });
  });
  return exports;
}

/* istanbul ignore next */
function noop() { /* empty */ }

// from https://github.com/nodejs/node/blob/master/lib/internal/url.js
function urlToOptions(urlObject) {
  var options = {
    protocol: urlObject.protocol,
    hostname: urlObject.hostname.startsWith("[") ?
      /* istanbul ignore next */
      urlObject.hostname.slice(1, -1) :
      urlObject.hostname,
    hash: urlObject.hash,
    search: urlObject.search,
    pathname: urlObject.pathname,
    path: urlObject.pathname + urlObject.search,
    href: urlObject.href,
  };
  if (urlObject.port !== "") {
    options.port = Number(urlObject.port);
  }
  return options;
}

function removeMatchingHeaders(regex, headers) {
  var lastValue;
  for (var header in headers) {
    if (regex.test(header)) {
      lastValue = headers[header];
      delete headers[header];
    }
  }
  return (lastValue === null || typeof lastValue === "undefined") ?
    undefined : String(lastValue).trim();
}

function createErrorType(code, defaultMessage) {
  function CustomError(cause) {
    Error.captureStackTrace(this, this.constructor);
    if (!cause) {
      this.message = defaultMessage;
    }
    else {
      this.message = defaultMessage + ": " + cause.message;
      this.cause = cause;
    }
  }
  CustomError.prototype = new Error();
  CustomError.prototype.constructor = CustomError;
  CustomError.prototype.name = "Error [" + code + "]";
  CustomError.prototype.code = code;
  return CustomError;
}

function abortRequest(request) {
  for (var e = 0; e < events.length; e++) {
    request.removeListener(events[e], eventHandlers[events[e]]);
  }
  request.on("error", noop);
  request.abort();
}

function isSubdomain(subdomain, domain) {
  const dot = subdomain.length - domain.length - 1;
  return dot > 0 && subdomain[dot] === "." && subdomain.endsWith(domain);
}

// Exports
module.exports = wrap({ http: http, https: https });
module.exports.wrap = wrap;


/***/ }),

/***/ 267:
/***/ ((module) => {

module.exports = eval("require")("debug");


/***/ }),

/***/ 491:
/***/ ((module) => {

"use strict";
module.exports = require("assert");

/***/ }),

/***/ 685:
/***/ ((module) => {

"use strict";
module.exports = require("http");

/***/ }),

/***/ 687:
/***/ ((module) => {

"use strict";
module.exports = require("https");

/***/ }),

/***/ 781:
/***/ ((module) => {

"use strict";
module.exports = require("stream");

/***/ }),

/***/ 310:
/***/ ((module) => {

"use strict";
module.exports = require("url");

/***/ }),

/***/ 796:
/***/ ((module) => {

"use strict";
module.exports = require("zlib");

/***/ }),

/***/ 837:
/***/ ((module) => {

"use strict";
module.exports = JSON.parse('{"name":"axios","version":"0.21.4","description":"Promise based HTTP client for the browser and node.js","main":"index.js","scripts":{"test":"grunt test","start":"node ./sandbox/server.js","build":"NODE_ENV=production grunt build","preversion":"npm test","version":"npm run build && grunt version && git add -A dist && git add CHANGELOG.md bower.json package.json","postversion":"git push && git push --tags","examples":"node ./examples/server.js","coveralls":"cat coverage/lcov.info | ./node_modules/coveralls/bin/coveralls.js","fix":"eslint --fix lib/**/*.js"},"repository":{"type":"git","url":"https://github.com/axios/axios.git"},"keywords":["xhr","http","ajax","promise","node"],"author":"Matt Zabriskie","license":"MIT","bugs":{"url":"https://github.com/axios/axios/issues"},"homepage":"https://axios-http.com","devDependencies":{"coveralls":"^3.0.0","es6-promise":"^4.2.4","grunt":"^1.3.0","grunt-banner":"^0.6.0","grunt-cli":"^1.2.0","grunt-contrib-clean":"^1.1.0","grunt-contrib-watch":"^1.0.0","grunt-eslint":"^23.0.0","grunt-karma":"^4.0.0","grunt-mocha-test":"^0.13.3","grunt-ts":"^6.0.0-beta.19","grunt-webpack":"^4.0.2","istanbul-instrumenter-loader":"^1.0.0","jasmine-core":"^2.4.1","karma":"^6.3.2","karma-chrome-launcher":"^3.1.0","karma-firefox-launcher":"^2.1.0","karma-jasmine":"^1.1.1","karma-jasmine-ajax":"^0.1.13","karma-safari-launcher":"^1.0.0","karma-sauce-launcher":"^4.3.6","karma-sinon":"^1.0.5","karma-sourcemap-loader":"^0.3.8","karma-webpack":"^4.0.2","load-grunt-tasks":"^3.5.2","minimist":"^1.2.0","mocha":"^8.2.1","sinon":"^4.5.0","terser-webpack-plugin":"^4.2.3","typescript":"^4.0.5","url-search-params":"^0.10.0","webpack":"^4.44.2","webpack-dev-server":"^3.11.0"},"browser":{"./lib/adapters/http.js":"./lib/adapters/xhr.js"},"jsdelivr":"dist/axios.min.js","unpkg":"dist/axios.min.js","typings":"./index.d.ts","dependencies":{"follow-redirects":"^1.14.0"},"bundlesize":[{"path":"./dist/axios.min.js","threshold":"5kB"}]}');

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __nccwpck_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		var threw = true;
/******/ 		try {
/******/ 			__webpack_modules__[moduleId](module, module.exports, __nccwpck_require__);
/******/ 			threw = false;
/******/ 		} finally {
/******/ 			if(threw) delete __webpack_module_cache__[moduleId];
/******/ 		}
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
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
// This entry need to be wrapped in an IIFE because it need to be in strict mode.
(() => {
"use strict";
// ESM COMPAT FLAG
__nccwpck_require__.r(__webpack_exports__);

;// CONCATENATED MODULE: ../service/engine/pipe.ts
class IFramePipe {
    constructor(secret) {
        this.set_handler = (fn) => {
            this._handlers.add(fn);
        };
        this.remove_handler = (fn) => {
            this._handlers.delete(fn);
        };
        this.send = (xid, action, data) => {
            const message = JSON.stringify({
                xid,
                data,
                action,
                parent_secret: this._secret,
            });
            window.parent.postMessage(message, '*');
        };
        this._secret = secret;
        this._handlers = new Set();
        window.addEventListener('message', (ev) => {
            const decoded = JSON.parse(ev.data);
            this._handlers.forEach((fn) => fn(decoded.xid, decoded.action, decoded.data));
        });
    }
}

;// CONCATENATED MODULE: ../core/engine/registry/index.ts
class Registry {
    constructor() {
        this.RegisterFactory = (type, name, factory) => {
            console.log(`START REGISTER FACTORY => type(${type}) name(${name})`);
            const key = [type, name].toString();
            this._factories.set(key, factory);
            const watchers = this._watchers.get(key);
            if (watchers) {
                console.log("Found watchers ", watchers);
                watchers.forEach((watcher) => watcher());
            }
            const typeWatchers = this._type_watchers.get(type);
            if (typeWatchers) {
                typeWatchers.forEach((f) => f(factory));
            }
            console.log(`END REGISTER FACTORY => type(${type}) name(${name})`);
        };
        this.WatchLoad = async (type, name, timeout) => {
            console.log("before Watching");
            const key = [type, name].toString();
            if (this._factories.has(key)) {
                console.log("found factories already");
                return Promise.resolve();
            }
            const p = new Promise((resolve, reject) => {
                console.log("making promise");
                let oldwatcher = this._watchers.get(key);
                if (!oldwatcher) {
                    oldwatcher = new Array(0);
                }
                oldwatcher.push(() => {
                    resolve();
                });
                this._watchers.set(key, oldwatcher);
                setTimeout(() => {
                    reject(`TimeOut loading type ${type} & name ${name}`);
                }, timeout);
            });
            return p;
        };
        this.OnTypeLoad = (typename, callback) => {
            let oldwatcher = this._type_watchers.get(typename);
            if (!oldwatcher) {
                oldwatcher = new Array(0);
            }
            oldwatcher.push(callback);
        };
        this.Get = (type, name) => {
            const key = [type, name].toString();
            return this._factories.get(key.toString());
        };
        this.GetAll = (type) => {
            const facts = Array(0);
            this._factories.forEach((fact, [_type, _]) => {
                if (type !== _type) {
                    return;
                }
                facts.push(fact);
            });
            return facts;
        };
        this.InstanceAll = (type, opts) => {
            this._factories.forEach((fact, key) => {
                const [_type, _] = key.split(',');
                if (type !== _type) {
                    return;
                }
                fact(opts);
            });
        };
        this.Instance = (type, name, opts) => {
            const key = [type, name].toString();
            this._factories.get(key)(opts);
        };
        this._factories = new Map();
        this._watchers = new Map();
        this._type_watchers = new Map();
    }
}
const initRegistry = () => {
    if (window["__registry__"]) {
        console.warn("Registry already loaded, skipping...");
        return;
    }
    const r = new Registry();
    r.RegisterFactory("loader.factory", "std.loader", async (opts) => {
        await opts.registry.WatchLoad("plug.factory", opts.entry, 200000);
        const factory = opts.registry.Get("plug.factory", opts.entry);
        if (!factory) {
            console.warn("could not load plug factory");
            return;
        }
        factory({
            plug: opts.plug,
            agent: opts.agent,
            entry: opts.entry,
            env: opts.env,
            target: opts.target,
            payload: opts.payload,
            registry: opts.registry
        });
    });
    console.log("GLOBAL_REGISTRY =>", r);
    window["__registry__"] = r;
    window["__register_factory__"] = r.RegisterFactory;
};
// it will find appoprate loader and call loader
// then its loader responsibility to start registered factories
// plugStart => loader => actual_plug_factory_start (using entry_name)
const plugStart = async (opts) => {
    console.log("let there be light", opts);
    const registry = window["__registry__"];
    if (!registry) {
        console.warn("registry not found");
        return;
    }
    if (!opts.exec_loader) {
        opts.exec_loader = "std.loader";
    }
    try {
        await registry.WatchLoad("loader.factory", opts.exec_loader, 100000);
    }
    catch (error) {
        console.warn("could not load, error occured:", error);
        return;
    }
    const loaderFactory = registry.Get("loader.factory", opts.exec_loader);
    if (!opts.target) {
        opts.target = document.body;
    }
    loaderFactory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: opts.env,
        registry: registry,
        target: opts.target,
        payload: opts.payload
    });
};

// EXTERNAL MODULE: ../../../node_modules/axios/index.js
var node_modules_axios = __nccwpck_require__(958);
;// CONCATENATED MODULE: ../core/api/base.ts

class base_ApiBase {
    constructor(opts) {
        this._user_token = opts.user_token;
        this._api_base_url = opts.url;
        this._service_options = opts.service_opts;
        this._session_token = "";
        this._http = null;
        this._service_path = opts.path;
        this.intercept_request = this.intercept_request.bind(this);
        this.intercept_request_err = this.intercept_request_err.bind(this);
        this._raw_http = axios.create({
            baseURL: opts.url,
        });
    }
    async init() {
        let resp = await this.refresh_token();
        this._service_resp_payload = resp.data["service_payload"] || null;
        this._session_token = resp.data.token;
        this._http = axios.create({
            headers: { Authorization: this._session_token },
            baseURL: this._api_base_url,
        });
        this._http.interceptors.request.use(this.intercept_request, this.intercept_request_err);
    }
    async refresh_token() {
        return this._raw_http.post(`/auth/refresh`, {
            user_token: this._user_token,
            options: this._service_options,
            path: this._service_path,
        });
    }
    intercept_request(config) {
        return config;
    }
    intercept_request_err(error) {
        // fixme => if error is 401, refresh the token
        return Promise.reject(error);
    }
    get(url, config) {
        return this._http.get(url, config);
    }
    post(url, data, config) {
        return this._http.post(url, data, config);
    }
    put(url, data, config) {
        return this._http.put(url, data, config);
    }
    patch(url, data, config) {
        return this._http.patch(url, data, config);
    }
    delete(url, config) {
        return this._http.delete(url, config);
    }
}

;// CONCATENATED MODULE: ../core/api/admin.ts

class TenantAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "tenant"],
        });
    }
    async update_tenant(data) {
        return this.post("/tenant", data);
    }
    async list_tenant_domain() {
        return this.get("/tenant/domain");
    }
    async add_tenant_domain(data) {
        return this.post("/tenant/domain/", data);
    }
    async get_tenant_domain(id) {
        return this.get(`/tenant/domain/${id}`);
    }
    async update_tenant_domain(id, data) {
        return this.post(`/tenant/domain/${id}`, data);
    }
    async remove_tenant_domain(id) {
        return this.delete(`/tenant/domain/${id}`);
    }
    // widget
    async list_domain_widget(did) {
        return this.get(`/tenant/domain/${did}/widget`);
    }
    async add_domain_widget(did, data) {
        return this.get(`/tenant/domain/${did}/widget`, data);
    }
    async get_domain_widget(did, wid) {
        return this.get(`/tenant/domain/${did}/widget${wid}`);
    }
    async update_domain_widget(did, wid, data) {
        return this.get(`/tenant/domain/${did}/widget${wid}`, data);
    }
    async remove_domain_widget(did, wid) {
        return this.delete(`/tenant/domain/${did}/widget${wid}`);
    }
}
class BprintAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "bprint"],
        });
    }
    async bprint_list() {
        return this.get("/bprint");
    }
    async bprint_create(data) {
        return this.post("/bprint", data);
    }
    async bprint_get(id) {
        return this.get(`/bprint/${id}`);
    }
    async bprint_update(id, data) {
        return this.post(`/bprint/${id}`, data);
    }
    async bprint_remove(id) {
        return this.delete(`/bprint/${id}`);
    }
    async bprint_install(id, opts) {
        return this.post(`/bprint/${id}/install`, opts);
    }
    async bprint_list_files(id) {
        return this.get(`/bprint/${id}/file`);
    }
    async bprint_get_file(id, file) {
        return this.get(`/bprint/${id}/file/${file}`);
    }
    async bprint_new_file(id, file, data) {
        return this.post(`/bprint/${id}/file/${file}`, data);
    }
    async bprint_update_file(id, file, data) {
        return this.patch(`/bprint/${id}/file/${file}`, data);
    }
    async bprint_del_file(id, file) {
        return this.delete(`/bprint/${id}/file/${file}`);
    }
    async bprint_import(data) {
        return this.post(`/import_bprint`, data);
    }
    async repo_sources() {
        return this.get(`/repo`);
    }
    async repo_list(source) {
        return this.get(`/repo/${source}`);
    }
    async repo_get(source, group, slug) {
        return this.get(`/repo/${source}/${group}/${slug}`);
    }
    async repo_get_file(source, slug, file) {
        return this.get(`/repo/${source}/${slug}/${file}`);
    }
}
class UserAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "user"],
        });
    }
    async list_users(group) {
        return this.get(`/user${group ? `?user_group=` + group : ""}`);
    }
    async add_user(data) {
        return this.post(`/user`, data);
    }
    async get_user_by_id(id) {
        return this.get(`/user/${id}`);
    }
    async update_user(id, data) {
        return this.post(`/user/${id}`, data);
    }
    async remove_user(id) {
        return this.delete(`/user/${id}`);
    }
    async list_user_group() {
        return this.get(`/user_group`);
    }
    async add_user_group(data) {
        return this.post(`/user_group`, data);
    }
    async get_user_group(gid) {
        return this.get(`/user_group/${gid}`);
    }
    async update_user_group(gid, data) {
        return this.post(`/user_group/${gid}`, data);
    }
    async remove_user_group(gid) {
        return this.delete(`/user_group/${gid}`);
    }
    // auth
    async user_group_list_auth(gid) {
        return this.get(`/user_auth/${gid}`);
    }
    async user_group_add_auth(gid, data) {
        return this.post(`/user_auth/${gid}`, data);
    }
    async user_group_get_auth(gid, id) {
        return this.get(`/user_auth/${gid}/${id}`);
    }
    async user_group_update_auth(gid, id, data) {
        return this.post(`/user_auth/${gid}/${id}`, data);
    }
    async user_group_remove_auth(gid, id) {
        return this.delete(`/user_auth/${gid}/${id}`);
    }
    // hook
    async user_group_list_hook(gid) {
        return this.get(`/user_hook/${gid}`);
    }
    async user_group_add_hook(gid, data) {
        return this.post(`/user_hook/${gid}`, data);
    }
    async user_group_get_hook(gid, id) {
        return this.get(`/user_hook/${gid}/${id}`);
    }
    async user_group_update_hook(gid, id, data) {
        return this.post(`/user_hook/${gid}/${id}`, data);
    }
    async user_group_remove_hook(gid, id) {
        return this.get(`/user_hook/${gid}/${id}`);
    }
    // plug
    async user_group_list_plug(gid) {
        return this.get(`/user_plug/${gid}`);
    }
    async user_group_add_plug(gid, data) {
        return this.post(`/user_plug/${gid}`, data);
    }
    async user_group_get_plug(gid, id) {
        return this.get(`/user_plug/${gid}/${id}`);
    }
    async user_group_update_plug(gid, id, data) {
        return this.post(`/user_plug/${gid}/${id}`, data);
    }
    async user_group_remove_plug(gid, id) {
        return this.get(`/user_plug/${gid}/${id}`);
    }
    // data
    async user_group_list_data(gid) {
        return this.get(`/user_data/${gid}`);
    }
    async user_group_add_data(gid, data) {
        return this.post(`/user_data/${gid}`, data);
    }
    async user_group_get_data(gid, id) {
        return this.get(`/user_data/${gid}/${id}`);
    }
    async user_group_update_data(gid, id, data) {
        return this.post(`/user_data/${gid}/${id}`, data);
    }
    async user_group_remove_data(gid, id) {
        return this.get(`/user_data/${gid}/${id}`);
    }
}
class PlugAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin", "plug"],
        });
    }
    async list_plug() {
        return this.get(`/plug`);
    }
    async new_plug(data) {
        return this.post(`/plug`, data);
    }
    async update_plug(id, data) {
        return this.post(`/plug/${id}`, data);
    }
    async get_plug(pid) {
        return this.get(`/plug/${pid}`);
    }
    async del_plug(pid) {
        return this.delete(`/plug/${pid}`);
    }
    async list_agent(pid) {
        return this.get(`/plug/${pid}/agent`);
    }
    async new_agent(pid, data) {
        return this.post(`/plug/${pid}/agent`, data);
    }
    async update_agent(pid, aid, data) {
        return this.post(`/plug/${pid}/agent/${aid}`, data);
    }
    async get_agent(pid, aid) {
        return this.get(`/plug/${pid}/agent/${aid}`);
    }
    async del_agent(pid, aid) {
        return this.delete(`/plug/${pid}/agent/${aid}`);
    }
    async launch_agent(plug, agent, data) {
        return this.post(`/engine/${plug}/${agent}/launcher/json`, data);
    }
}
class CabinetAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token, source) {
        super({
            url: url,
            user_token: user_token,
            path: ["cabinet", source],
        });
    }
    async list_root() {
        return this.get(`/cabinet`);
    }
    async list_folder(folder) {
        return this.get(`/cabinet/${folder}`);
    }
    async new_folder(folder) {
        return this.post(`/cabinet/${folder}`);
    }
    async get_file(folder, file) {
        return this.get(`/cabinet/${folder}/file/${file}`);
    }
    async upload_file(folder, file, data) {
        return this.post(`/cabinet/${folder}/file/${file}`, data);
    }
    async delete_file(folder, file) {
        return this.delete(`/cabinet/${folder}/file/${file}`);
    }
    async get_folder_ticket(folder) {
        return this.post(`/cabinet/${folder}/ticket`);
    }
}
class ResourceAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["resource"],
        });
    }
    async agent_resources_list(data) {
        return this.post("/agent_resources", data);
    }
    async resource_list() {
        return this.get("/resource");
    }
    async resource_create(data) {
        return this.post("/resource", data);
    }
    async resource_get(slug) {
        return this.get(`/resource/${slug}`);
    }
    async resource_update(slug, data) {
        return this.post(`/resource/${slug}`, data);
    }
    async resource_remove(slug) {
        return this.delete(`/resource/${slug}`);
    }
}

;// CONCATENATED MODULE: ../core/api/auth.ts

class AuthAPI {
    constructor(api_url, site_token) {
        this.list_methods = async (ugroup) => {
            return this.http.get(`/auth?ugroup=${ugroup}`);
        };
        this.login_next = async (data) => {
            return this.http.post("/auth/login/next", data);
        };
        this.login_submit = async (data) => {
            return this.http.post("/auth/login/submit", data);
        };
        this.altauth_generate = async (id, data) => {
            return this.http.post(`/auth/alt/${id}/generate`, data);
        };
        this.altauth_next = async (id, stage, data) => {
            return this.http.post(`/auth/alt/${id}/next/${stage}`, data);
        };
        this.altauth_submit = async (id, data) => {
            return this.http.post(`/auth/alt/${id}/submit`, data);
        };
        this.finish = async (data) => {
            return this.http.post("/auth/finish", data);
        };
        this.signup_next = async (data) => {
            return this.http.post("/auth/signup/next", data);
        };
        this.signup_submit = async (data) => {
            return this.http.post("/auth/signup/submit", data);
        };
        this.reset_submit = async (data) => {
            return this.http.post("/reset/submit", data);
        };
        this.reset_finish = async (data) => {
            return this.http.post("/reset/finish", data);
        };
        this.site_token = site_token;
        this.http = axios.create({
            baseURL: api_url,
            headers: {
                Authorization: site_token,
            },
        });
    }
}

;// CONCATENATED MODULE: ../core/api/folder.ts
class FolderTktAPI {
    constructor(base_url, ticket) {
        this.ticket = ticket;
        this.base_url = base_url;
    }
    async list() {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}`);
        return resp.json();
    }
    async upload_file(file, data) {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}/${file}`, {
            method: "POST",
            body: data,
        });
        return resp.json();
    }
    get_file_link(file) {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/${file}`;
    }
    get_file_preview_link(file) {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/preview/${file}`;
    }
}

;// CONCATENATED MODULE: ../core/api/self.ts

class SelfAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin"],
        });
    }
    async list_cabinet_sources() {
        return this.get(`/cabinet_sources`);
    }
    async list_dgroup_sources() {
        return this.get(`/dgroup`);
    }
    async message_user(data) {
        return this.post("/self/message_user", data);
    }
    async get_user_info(userid) {
        return this.get(`/self/get_user_info/${userid}`);
    }
    async get_self_info() {
        return this.get("/self/get_self_info");
    }
    async update_self_info(data) {
        return this.post("/self/get_self_info", data);
    }
    async self_change_email(data) {
        return this.post("/self/change_email", data);
    }
    async self_change_auth(data) {
        return this.post("/self/change_auth", data);
    }
    async list_messages(data) {
        return this.post("/self/list_messages", data);
    }
    async modify_messages(data) {
        return this.post("/self/modify_messages", data);
    }
    async dtable_change(data) {
        return this.post("/self/dtable_change", data);
    }
    get_session_token() {
        return this._session_token;
    }
}

;// CONCATENATED MODULE: ../core/api/dyn.ts

class DtableAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token, source, group) {
        super({
            url: url,
            user_token: user_token,
            path: ["dtable", source, group],
        });
    }
    async load_group() {
        return this.get(`/dgroup_load`);
    }
    // dtable
    async list_tables() {
        return this.get(`/dtable`);
    }
    async add_table(data) {
        return this.post(`/dtable`, data);
    }
    async edit_table(tid, data) {
        return this.patch(`/dtable/${tid}`, data);
    }
    async get_table(tid) {
        return this.get(`/dtable/${tid}`);
    }
    async delete_table(tid) {
        return this.delete(`/dtable/${tid}`);
    }
    async list_columns(tid) {
        return this.get(`/dtable/${tid}/column`);
    }
    async add_column(tid, data) {
        return this.post(`/dtable/${tid}/column`, data);
    }
    async get_column(tid, cid) {
        return this.get(`/dtable/${tid}/column/${cid}`);
    }
    async edit_column(tid, cid, data) {
        return this.patch(`/dtable/${tid}/column/${cid}`, data);
    }
    async delete_column(tid, cid) {
        return this.delete(`/dtable/${tid}/column/${cid}`);
    }
    // view stuff
    async list_view(tid) {
        return this.get(`/dtable/${tid}/view`);
    }
    async new_view(tid, data) {
        return this.post(`/dtable/${tid}/view`, data);
    }
    async modify_view(tid, id, data) {
        return this.post(`/dtable/${tid}/view/${id}`, data);
    }
    async get_view(tid, id) {
        return this.get(`/dtable/${tid}/view/${id}`);
    }
    async del_view(tid, id) {
        return this.delete(`/dtable/${tid}/view/${id}`);
    }
    // hook stuff
    async list_hook(tid) {
        return this.get(`/dtable/${tid}/hook`);
    }
    async new_hook(tid, data) {
        return this.post(`/dtable/${tid}/hook`, data);
    }
    async modify_hook(tid, id, data) {
        return this.post(`/dtable/${tid}/hook/${id}`, data);
    }
    async get_hook(tid, id) {
        return this.get(`/dtable/${tid}/hook/${id}`);
    }
    async del_hook(tid, id) {
        return this.delete(`/dtable/${tid}/hook/${id}`);
    }
    // dtable ops
    async new_row(tid, data) {
        return this.post(`/dtable_ops/${tid}/row`, data);
    }
    async get_row(tid, rid) {
        return this.get(`/dtable_ops/${tid}/row/${rid}`);
    }
    async update_row(tid, rid, data) {
        return this.post(`/dtable_ops/${tid}/row/${rid}`, data);
    }
    async delete_row(tid, rid) {
        return this.delete(`/dtable_ops/${tid}/row/${rid}`);
    }
    async simple_query(tid, data) {
        if (!data) {
            data = {};
        }
        return this.post(`/dtable_ops/${tid}/simple_query`, data);
    }
    async fts_query(tid, str) {
        return this.post(`/dtable_ops/${tid}/fts_query`, {
            qstr: str,
        });
    }
    async ref_load(tid, data) {
        return this.post(`/dtable_ops/${tid}/ref_load`, data);
    }
    async ref_resolve(tid, data) {
        return this.post(`/dtable_ops/${tid}/ref_resolve`, data);
    }
    async rev_ref_load(tid, data) {
        return this.post(`/dtable_ops/${tid}/rev_ref_load`, data);
    }
    async list_activity(tid, rowid) {
        return this.get(`/dtable_ops/${tid}/activity/${rowid}`);
    }
    async comment_row(tid, rowid, msg) {
        return this.post(`/dtable_ops/${tid}/activity/${rowid}`, {
            message: msg,
        });
    }
}
class DynAPI extends (/* unused pure expression or super */ null && (ApiBase)) {
    constructor(url, user_token) {
        super({
            url: url,
            user_token: user_token,
            path: ["admin"],
        });
    }
    async list_group(source) {
        return this.get(`/dgroup/${source}`);
    }
    async get_group(source, group) {
        return this.get(`/dgroup/${source}/${group}`);
    }
    async new_group(source, data) {
        return this.post(`/dgroup/${source}`, data);
    }
    async edit_group(source, gid, data) {
        return this.patch(`/dgroup/${source}/${gid}`, data);
    }
    async delete_group(source, gid) {
        return this.delete(`/dgroup/${source}/${gid}`);
    }
}

;// CONCATENATED MODULE: ../core/api/index.ts







;// CONCATENATED MODULE: ../vendor/ws/backoff/linearbackoff.ts
/**
 * LinearBackoff increases the backoff-time by a constant number with
 * every step. An optional maximum can be provided as an upper bound
 * to the returned backoff.
 *
 * Example: for initial=0, increment=2000, maximum=8000 the Linear-
 * Backoff will produce the series [0, 2000, 4000, 6000, 8000].
 */
class LinearBackoff {
    constructor(initial, increment, maximum) {
        this.initial = initial;
        this.increment = increment;
        this.maximum = maximum;
        this.current = this.initial;
    }
    next() {
        const backoff = this.current;
        const next = this.current + this.increment;
        if (this.maximum === undefined)
            this.current = next;
        else if (next <= this.maximum)
            this.current = next;
        return backoff;
    }
    reset() {
        this.current = this.initial;
    }
}

;// CONCATENATED MODULE: ../vendor/ws/buffer/lrubuffer.ts
/**
 * LRUBuffer is a buffer that keeps the last n elements. When it is
 * full and written to, the oldest element in the buffer will be
 * replaced. When reading from the LRUBuffer, elements are returned
 * in FIFO-order (queue).
 *
 * LRUBuffer has linear space- and time-requirements. Internally
 * an array is used as a circular-buffer. All memory is allocated
 * on initialization.
 */
class LRUBuffer {
    constructor(len) {
        this.writePtr = 0;
        this.wrapped = false;
        this.buffer = Array(len);
    }
    len() {
        return this.wrapped ? this.buffer.length : this.writePtr;
    }
    cap() {
        return this.buffer.length;
    }
    read(es) {
        if (es === null || es === undefined || es.length === 0 || this.buffer.length === 0)
            return 0;
        if (this.writePtr === 0 && !this.wrapped)
            return 0;
        const first = this.wrapped ? this.writePtr : 0;
        const last = (first - 1) < 0 ?
            this.buffer.length - 1 :
            first - 1;
        for (let i = 0; i < es.length; i++) {
            let r = (first + i) % this.buffer.length;
            es[i] = this.buffer[r];
            if (r === last)
                return i + 1;
        }
        return es.length;
    }
    write(es) {
        if (es === null || es === undefined || es.length === 0 || this.buffer.length === 0)
            return 0;
        const start = es.length > this.buffer.length ? es.length - this.buffer.length : 0;
        for (let i = 0; i < es.length - start; i++) {
            this.buffer[this.writePtr] = es[start + i];
            this.writePtr = (this.writePtr + 1) % this.buffer.length;
            if (this.writePtr === 0)
                this.wrapped = true;
        }
        return es.length;
    }
    forEach(fn) {
        if (this.writePtr === 0 && !this.wrapped)
            return 0;
        let cur = this.wrapped ? this.writePtr : 0;
        const last = this.wrapped ? (cur - 1) < 0 ? this.buffer.length - 1 : cur - 1 : this.writePtr - 1;
        const len = this.len();
        while (true) {
            fn(this.buffer[cur]);
            if (cur === last)
                break;
            cur = (cur + 1) % this.buffer.length;
        }
        return len;
    }
    clear() {
        this.writePtr = 0;
        this.wrapped = false;
    }
}

;// CONCATENATED MODULE: ../vendor/ws/websocket.ts
var WebsocketEvents;
(function (WebsocketEvents) {
    WebsocketEvents["open"] = "open";
    WebsocketEvents["close"] = "close";
    WebsocketEvents["error"] = "error";
    WebsocketEvents["message"] = "message";
    WebsocketEvents["retry"] = "retry"; // A try to re-connect is made
})(WebsocketEvents || (WebsocketEvents = {}));
class Websocket {
    constructor(url, protocols, buffer, backoff) {
        this.eventListeners = { open: [], close: [], error: [], message: [], retry: [] };
        this.closedByUser = false;
        this.retries = 0;
        this.handleOpenEvent = (ev) => this.handleEvent(WebsocketEvents.open, ev);
        this.handleCloseEvent = (ev) => this.handleEvent(WebsocketEvents.close, ev);
        this.handleErrorEvent = (ev) => this.handleEvent(WebsocketEvents.error, ev);
        this.handleMessageEvent = (ev) => this.handleEvent(WebsocketEvents.message, ev);
        this.url = url;
        this.protocols = protocols;
        this.buffer = buffer;
        this.backoff = backoff;
        this.tryConnect();
    }
    getUnderlyingWebsocket() {
        return this.websocket;
    }
    send(data) {
        if (this.closedByUser)
            return;
        if (this.websocket === undefined || this.websocket.readyState !== this.websocket.OPEN)
            this.buffer?.write([data]);
        else
            this.websocket.send(data);
    }
    close(code, reason) {
        this.closedByUser = true;
        this.websocket?.close(code, reason);
    }
    addEventListener(type, listener, options) {
        const eventListener = { listener, options };
        const eventListeners = this.eventListeners[type];
        eventListeners.push(eventListener);
    }
    removeEventListener(type, listener, options) {
        this.eventListeners[type] =
            this.eventListeners[type]
                .filter(l => {
                return l.listener !== listener && (l.options === undefined || l.options !== options);
            });
    }
    dispatchEvent(type, ev) {
        const listeners = this.eventListeners[type];
        const onceListeners = [];
        listeners.forEach(l => {
            l.listener(this, ev); // call listener
            if (l.options !== undefined && l.options.once)
                onceListeners.push(l);
        });
        onceListeners.forEach(l => this.removeEventListener(type, l.listener, l.options)); // remove 'once'-listeners
    }
    tryConnect() {
        if (this.websocket !== undefined) { // remove all event-listeners from broken socket
            this.websocket.removeEventListener(WebsocketEvents.open, this.handleOpenEvent);
            this.websocket.removeEventListener(WebsocketEvents.close, this.handleCloseEvent);
            this.websocket.removeEventListener(WebsocketEvents.error, this.handleErrorEvent);
            this.websocket.removeEventListener(WebsocketEvents.message, this.handleMessageEvent);
            this.websocket.close();
        }
        this.websocket = new WebSocket(this.url, this.protocols); // create new socket and attach handlers
        this.websocket.addEventListener(WebsocketEvents.open, this.handleOpenEvent);
        this.websocket.addEventListener(WebsocketEvents.close, this.handleCloseEvent);
        this.websocket.addEventListener(WebsocketEvents.error, this.handleErrorEvent);
        this.websocket.addEventListener(WebsocketEvents.message, this.handleMessageEvent);
    }
    handleEvent(type, ev) {
        switch (type) {
            case WebsocketEvents.close:
                if (!this.closedByUser) // failed to connect or connection lost, try to reconnect
                    this.reconnect();
                break;
            case WebsocketEvents.open:
                this.retries = 0;
                this.backoff?.reset(); // reset backoff
                this.buffer?.forEach(this.send.bind(this)); // send all buffered messages
                this.buffer?.clear();
                break;
        }
        this.dispatchEvent(type, ev); // forward to all listeners
    }
    reconnect() {
        if (this.backoff === undefined) // no backoff, we're done
            return;
        const backoff = this.backoff.next();
        setTimeout(() => {
            this.dispatchEvent(WebsocketEvents.retry, new CustomEvent(WebsocketEvents.retry, {
                detail: {
                    retries: ++this.retries,
                    backoff: backoff
                }
            }));
            this.tryConnect();
        }, backoff);
    }
}

;// CONCATENATED MODULE: ../vendor/ws/websocketBuilder.ts

/**
 * Used to build Websocket-instances.
 */
class WebsocketBuilder {
    constructor(url) {
        this.ws = null;
        this.onOpenListeners = [];
        this.onCloseListeners = [];
        this.onErrorListeners = [];
        this.onMessageListeners = [];
        this.onRetryListeners = [];
        this.url = url.replace("http://", "ws://").replace("https://", "wss://");
    }
    withProtocols(p) {
        this.protocols = p;
        return this;
    }
    withBackoff(backoff) {
        this.backoff = backoff;
        return this;
    }
    withBuffer(buffer) {
        this.buffer = buffer;
        return this;
    }
    onOpen(listener, options) {
        this.onOpenListeners.push({ listener, options });
        return this;
    }
    onClose(listener, options) {
        this.onCloseListeners.push({ listener, options });
        return this;
    }
    onError(listener, options) {
        this.onErrorListeners.push({ listener, options });
        return this;
    }
    onMessage(listener, options) {
        this.onMessageListeners.push({ listener, options });
        return this;
    }
    onRetry(listener, options) {
        this.onRetryListeners.push({ listener, options });
        return this;
    }
    /**
     * Multiple calls to build() will always return the same websocket-instance.
     */
    build() {
        if (this.ws !== null)
            return this.ws;
        this.ws = new Websocket(this.url, this.protocols, this.buffer, this.backoff);
        this.onOpenListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.open, h.listener, h.options));
        this.onCloseListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.close, h.listener, h.options));
        this.onErrorListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.error, h.listener, h.options));
        this.onMessageListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.message, h.listener, h.options));
        this.onRetryListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.retry, h.listener, h.options));
        return this.ws;
    }
}

;// CONCATENATED MODULE: ../vendor/ws/index.ts








;// CONCATENATED MODULE: ../core/sockd/sockd.ts

class Sockd {
    constructor(url) {
        this.init = async () => {
            this._ws = this._builder.build();
        };
        this.handleIncoming = (_, ev) => {
            // fixme => handle system messages
            const data = JSON.parse(ev.data);
            this._handler(data);
        };
        this.OnSockdMessage = (h) => {
            this._handler = h;
        };
        this.SendSockd = (message) => {
            this._ws.send(JSON.stringify(message));
        };
        console.log("CONNECTING WS @ ", url);
        this._builder = new WebsocketBuilder(url);
        this._builder.onMessage(this.handleIncoming);
        this._builder.withBackoff(new LinearBackoff(1, 3));
        this._builder.withBuffer(new LRUBuffer(20));
    }
}

;// CONCATENATED MODULE: ../core/sockd/stypes.ts
const MESSAGE_SERVER_DIRECT = "server_direct";
const MESSAGE_SERVER_BROADCAST = "server_broadcast";
const MESSAGE_SERVER_PUBLISH = "server_publish";
const MESSAGE_PEER_DIRECT = "peer_direct";
const MESSAGE_PEER_BROADCAST = "peer_broadcast";
const MESSAGE_PEER_PUBLISH = "peer_publish";


;// CONCATENATED MODULE: ../core/sockd/room.ts

class SockdRoom {
    constructor(socket, room) {
        this.SendDirect = (data) => {
            this._socket.SendSockd({
                payload: data,
                type: MESSAGE_PEER_DIRECT,
                xid: "",
                from_id: "",
                room: this._room,
            });
        };
        this.SendBroadcast = (data) => {
            this._socket.SendSockd({
                payload: data,
                type: MESSAGE_PEER_BROADCAST,
                xid: "",
                from_id: "",
                room: this._room,
            });
        };
        this.SendTagged = (data, ticket, targets) => {
            this._socket.SendSockd({
                payload: data,
                type: MESSAGE_PEER_PUBLISH,
                xid: "",
                from_id: "",
                room: this._room,
                targets: targets,
                ticket: ticket,
            });
        };
        this.onMessage = (handler) => {
            this._onMessage = handler;
        };
        this.onPeer = (handler) => {
            this._onPeer = handler;
        };
        this.onServer = (handler) => {
            this._onServer = handler;
        };
        this.ProcessMessage = (message) => {
            if (this._onMessage) {
                this._onMessage(message);
            }
            switch (message.type) {
                case MESSAGE_SERVER_DIRECT:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case MESSAGE_SERVER_BROADCAST:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case MESSAGE_SERVER_PUBLISH:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case MESSAGE_PEER_DIRECT:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                case MESSAGE_PEER_BROADCAST:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                case MESSAGE_PEER_PUBLISH:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                default:
                    break;
            }
        };
        this.IsConnected = async () => {
            return false;
        };
        this.LeaveRoom = () => {
            // fixme => impl
        };
        this._socket = socket;
        this._room = room;
    }
}

;// CONCATENATED MODULE: ../core/engine/env/fetch.ts
const actionFetch = (actionUrl, token) => async (name, data) => {
    const response = await fetch(`${actionUrl}/${name}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: token,
        },
        redirect: "follow",
        referrerPolicy: "strict-origin-when-cross-origin",
        body: data,
    });
    return response;
};

;// CONCATENATED MODULE: ../core/engine/env/env.ts




class Env {
    constructor(opts) {
        this.init = async () => {
            await this._sockd.init();
        };
        this.PreformAction = async (name, data) => {
            const encoded = JSON.stringify(data);
            try {
                const resp = await this._fetch(name, encoded);
                const ctype = resp.headers.get("Content-Type");
                if (resp.status !== 200) {
                    const txt = await resp.text();
                    return {
                        status_ok: false,
                        content_type: ctype,
                        body: txt,
                    };
                }
                const respData = await resp.json();
                return {
                    body: respData,
                    content_type: ctype,
                    status_ok: true,
                };
            }
            catch (error) {
                return {
                    status_ok: false,
                    body: error,
                };
            }
        };
        this.startup_payload = () => {
            return this._startup_payload;
        };
        this.PreformParentAction = async (name, data) => {
            const key = "fixme => generate";
            const p = new Promise((resolve, reject) => {
            });
            this._pending_pipe_msg.set(key, null);
            this._pipe.send("aaa", name, data);
            // fixme => implement
        };
        this.FolderTktAPI = (ticket) => {
            return new FolderTktAPI(this._opts.base_url, ticket);
        };
        this.SockdAPI = (room) => {
            let rs = this._sockd_rooms.get(room);
            if (!rs) {
                rs = new SockdRoom(this._sockd, room);
                this._sockd_rooms.set(room, rs);
            }
            return rs;
        };
        window["debug_env"] = this; // only for debug remove this 
        this._opts = opts;
        this._sockd_rooms = new Map();
        this._pending_pipe_msg = new Map();
        this._pipe = opts.pipe;
        this._startup_payload = opts.startup_payload;
        this._fetch = actionFetch(`${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_con`, opts.token);
        const sockdUrl = `${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_ws`;
        this._sockd = new Sockd(sockdUrl);
        this._sockd.OnSockdMessage((msg) => {
            if (!msg.room) {
                console.log("no room message", msg);
                return;
            }
            if (msg.room === "plugs_dev") {
                console.log("PLUG DEBUG =>", msg.payload);
                return;
            }
            const room = this._sockd_rooms.get(msg.room);
            if (!room) {
                console.log("room without handler =>");
                return;
            }
            room.ProcessMessage(msg);
        });
    }
}

;// CONCATENATED MODULE: ../core/engine/env/index.ts


;// CONCATENATED MODULE: ../altentry/execiframe/index.ts



console.log("init registry");
initRegistry();
window.addEventListener("load", async () => {
    const opts = window["__loader_options__"];
    if (!opts) {
        console.log("Loader Options not found");
        return;
    }
    console.log("iframe portal opts @=>", opts);
    const pipe = new IFramePipe(opts.parent_secret);
    const env = new Env({
        agent: opts.agent,
        plug: opts.plug,
        token: opts.token,
        base_url: opts.base_url,
        parent_secret: opts.parent_secret,
        pipe,
    });
    await env.init();
    pipe.send("", "env_loaded", {});
    plugStart({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: env,
        target: document.getElementById("plugroot"),
        exec_loader: opts.exec_loader,
        payload: null,
    });
}, false);

})();

module.exports = __webpack_exports__;
/******/ })()
;