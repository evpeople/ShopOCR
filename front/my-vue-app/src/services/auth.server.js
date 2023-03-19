import axiosUser from '../utils/axios/axiosUser';


class AuthService {
  login(user) {
    return axiosUser
        .post('login', {username: user.username, password: user.password})
        .then(response => {
          if (response.data.accessToken) {
            localStorage.setItem('user', JSON.stringify(response.data));
          }

          return response.data;
        });
  }

  logout() {
    localStorage.removeItem('user');
  }

  register(user) {
    return axiosUser.post(
        'register',
        {username: user.username, email: user.email, password: user.password});
  }
}

export default new AuthService();