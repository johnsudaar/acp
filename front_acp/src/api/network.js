import axios from 'axios';

export default {
  getNetwork() {
    return axios.get('http://localhost:8888/static/network.json')
  }
}

