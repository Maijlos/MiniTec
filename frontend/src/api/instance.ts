import axios from 'axios';

const instance = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});

instance.interceptors.request.use(config => {
  const token = import.meta.env.VITE_API_KEY;
  if (token) {
    config.headers['X-API-KEY'] = token;
  }
  return config;
}, error => {
  return Promise.reject(error);
});

export default instance;
