import api from '../services/api';
import { AuthRepository } from '@/domain/repositories/AuthRepository';
import { AuthDTO } from '../dto/AuthDTO';

export class AuthRepositoryImpl extends AuthRepository {
  async login({ email, password }) {
    const response = await api.post('/dashboard/v1/auth/login', { email, password });
    return AuthDTO.toEntity(response.data);
  }
}
