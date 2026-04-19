export class User {
  constructor({ email, role, token }) {
    this.email = email;
    this.role = role;
    this.token = token;
  }

  get isAuthenticated() {
    return !!this.token;
  }

  get isAdmin() {
    return this.role === 'admin';
  }
}
