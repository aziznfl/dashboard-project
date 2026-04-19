import { AuthRepositoryImpl } from '@/data/repositories/AuthRepositoryImpl';

const authRepository = new AuthRepositoryImpl();

export const AuthService = {
  async login(email, password) {
    try {
      const user = await authRepository.login({ email, password });
      return user;
    } catch (error) {
      throw error.response?.data?.message || 'Login failed';
    }
  }
};
