const axios = require('axios');
const API_URL = 'http://localhost:8080/api/';
function login(username, password) {
    return axios
      .post(API_URL + "signin", {
        username,
        password
      });
}

login('sdfhsde','sdfsdsdfsdfsd').then((result) => {
    console.log(result.data);
});
