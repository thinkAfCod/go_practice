import axios from 'axios'

axios.defaults.baseURL = 'http://192.168.50.109:3000'
axios.defaults.timeout = 5000
axios.defaults.headers['Content-Type'] = 'application/json;charset=UTF-8';
axios.defaults.headers['Access-Control-Allow-Origin'] = '*';

axios.interceptors.request.use((config) => {
  // config.setHeaders([
  //     // 在这里设置请求头与携带token信息
  // ]);
  return config
}, (error) => {
  return Promise.reject(error);
});

axios.interceptors.response.use(
  response => {
    return response.data;
  },
  error => {
    //return Promise.reject(error.response) // 返回接口返回的错误信息
    return {
      status: 400,
      message: error.toString(),
    }
  }
);

export async function request(method, url, params, opts = {}) {
  //opts = Object.assign({}, opts);
  //const headers = Object.assign({}, axios.defaults.headers, opts.headers);
  // const withCredentials = opts.withCredentials || WITH_CREDENTIALS;
  let config = {
    method: method,
    url: url,
    headers: axios.defaults.headers,
    //withCredentials: true,
    timeout: opts && opts.timeout ? opts.timeout : 20000,
    //onUploadProgress: opts.onUploadProgress // 上传进度
  };
  switch (method) {
    case 'GET': // get请求
      config = Object.assign(config, {params: params});
      break;
    case 'POST':
      config = Object.assign(config, {data: params});
      break;
    case 'PUT':
      config = Object.assign(config, {data: params});
      break;
    case 'DELETE':
      config = Object.assign(config, {data: params});
      break;
    case 'PATCH':
      config = Object.assign(config, {data: params});
      break;
    default:
      console.log('no_match');
  }
  try {
    const res = await axios(config);
    return res;
  } catch (e) {
    return {};
  }
}

/**
 * 发起一个 get 请求
 * @param {*} args 参数：url,param,opts
 */
export function get(...args) {
  return request('GET', ...args);
}

/**
 * 发起一个 post 请求
 * @param {*} args 参数：url,param,opts
 */
export function post(...args) {
  return request('POST', ...args);
}

/**
 * 发起一个 put 请求
 * @param {*} args 参数：url,param,opts
 */
export function put(...args) {
  return request('PUT', ...args);
}

/**
 * 发起一个 delete 请求
 * @param {*} args 参数：url,param,opts
 */
export function deletes(...args) {
  return request('DELETE', ...args);
}

/**
 * 发起一个 patch 请求
 * @param {*} args 参数：url,param,opts
 */
export function patch(...args) {
  return request('PATCH', ...args);
}