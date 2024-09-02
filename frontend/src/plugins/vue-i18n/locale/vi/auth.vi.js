const auth = {
  login: {
    title: 'LOGIN',
    placeholderEmail: 'Email address',
    placeholderPassword: 'Password',
    forgotPassword: 'Forgot password?',
    redirectPage: {
      title: 'Dont have an account yet?',
      labelButton: 'Register now',
    },
    notification: {
      success: {
        title: 'Success',
        message: 'Login successful.',
      },
      fail: {
        title: 'Failed',
      },
    },
  },
  register: {
    title: 'REGISTER',
    subtitle: 'Please enter registration information',
    placeholderLastName: 'Last Name',
    placeholderFirstName: 'First Name',
    placeholderPhone: 'Phone number',
    placeholderEmail: 'Email address',
    placeholderPassword: 'Password',
    placeholderRepeatPassword: 'Re-enter Password',
    redirectPage: {
      title: 'Already have an account?',
      labelButton: 'Login',
    },
    notification: {
      success: {
        title: 'Success',
        message: 'Registration successful.',
      },
      fail: {
        title: 'Failure',
      },
    },
  },
};

export default auth;
