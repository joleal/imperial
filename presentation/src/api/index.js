
import { getInstance } from '@/auth';
import axios from 'axios';

// Register base URL
axios.defaults.baseURL = process.env.VUE_APP_BASE_URI;

// Add a request interceptor
axios.interceptors.request.use( async config => {
  const instance = getInstance(); //auth0
  const token = await instance.getTokenSilently();
  if(token){
    config.headers.Authorization = `Bearer ${token}`;
  }
  config.headers.Accept = 'application/json';
  config.headers['Content-Type'] = 'application/json;charset=UTF-8';

  return config;
});

const ImperialApi = {
    ActiveGames: async () => {
        const response = await axios.get('/games/active');
        return response.data;
    },
    OpenGames: async () => {
      const response = await axios.get('/games/open');
      return response.data;
    },
    CreateGame: async(options) => {
      await axios.post('/game',options);
    },
    Players: async(name) => {
      const response = await axios.get(`/users?name=${name}`);
      return response.data;
    }
}

export default ImperialApi;